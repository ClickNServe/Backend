package oauth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOAuthConfig() *oauth2.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Println("clientID not found in .env")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Println("clientSecret not found in .env")
	}

	redirectURL := os.Getenv("REDIRECT_URL")
	if redirectURL == "" {
		log.Println("redirectURL not found in .env")
	}

	oauthConf := &oauth2.Config {
		ClientID: 		clientID,
		ClientSecret: 	clientSecret,
		RedirectURL: 	redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	
	return oauthConf
}