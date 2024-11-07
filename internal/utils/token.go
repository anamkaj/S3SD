package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AuthToken struct {
	AccessToken string
	DirectTable string
	URL         string
}

func GetToken() (AuthToken, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		return AuthToken{}, err
	}

	access_token := os.Getenv("ACCESS_TOKEN")
	client_table := os.Getenv("DIRECT_TABLE")
	url := os.Getenv("URL")

	token := AuthToken{
		AccessToken: access_token,
		DirectTable: client_table,
		URL:         url,
	}

	return token, nil
}
