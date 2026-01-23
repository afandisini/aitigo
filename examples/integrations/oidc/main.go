package main

import (
	"context"
	"log"

	"aitigo/pkg/integrations/auth"
)

func main() {
	provider, err := auth.Discover(context.Background(), "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := auth.BuildOAuth2Config(provider, auth.OAuth2Config{
		ClientID:    "client-id",
		RedirectURL: "http://localhost:8080/callback",
		Scopes:      []string{"profile", "email"},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("auth url:", cfg.AuthCodeURL("state"))
}
