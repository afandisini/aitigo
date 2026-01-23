package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

func Discover(ctx context.Context, issuer string) (*oidc.Provider, error) {
	if issuer == "" {
		return nil, fmt.Errorf("issuer is required")
	}
	return oidc.NewProvider(ctx, issuer)
}

func NewVerifier(provider *oidc.Provider, clientID string) (*oidc.IDTokenVerifier, error) {
	if provider == nil {
		return nil, fmt.Errorf("provider is nil")
	}
	if clientID == "" {
		return nil, fmt.Errorf("clientID is required")
	}
	return provider.Verifier(&oidc.Config{ClientID: clientID}), nil
}

func VerifyJWT(ctx context.Context, verifier *oidc.IDTokenVerifier, rawToken string) (*oidc.IDToken, error) {
	if verifier == nil {
		return nil, fmt.Errorf("verifier is nil")
	}
	if rawToken == "" {
		return nil, fmt.Errorf("token is required")
	}
	return verifier.Verify(ctx, rawToken)
}

func BuildOAuth2Config(provider *oidc.Provider, cfg OAuth2Config) (*oauth2.Config, error) {
	if provider == nil {
		return nil, fmt.Errorf("provider is nil")
	}
	if cfg.ClientID == "" {
		return nil, fmt.Errorf("clientID is required")
	}
	if cfg.RedirectURL == "" {
		return nil, fmt.Errorf("redirectURL is required")
	}
	scopes := ensureOpenIDScope(cfg.Scopes)
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  cfg.RedirectURL,
		Scopes:       scopes,
	}, nil
}

func ensureOpenIDScope(scopes []string) []string {
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			return scopes
		}
	}
	return append([]string{oidc.ScopeOpenID}, scopes...)
}
