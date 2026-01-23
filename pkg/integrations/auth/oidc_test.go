package auth

import (
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
)

func TestEnsureOpenIDScope(t *testing.T) {
	scopes := ensureOpenIDScope([]string{"profile"})
	if len(scopes) == 0 || scopes[0] != oidc.ScopeOpenID {
		t.Fatalf("expected openid scope prefix")
	}
}
