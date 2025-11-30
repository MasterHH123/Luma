package config

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ClientService struct {
	Client *http.Client
	BaseURL string
	Token string
}

func InitTellerService() (*ClientService) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get current working directory", err)
	}
	fmt.Printf("Making the call from this directory: %s\n", cwd)

	err = godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file: ", err)	
	}

	cert, err := tls.LoadX509KeyPair("certificate.pem", "private_key.pem")
	if err != nil{
		log.Fatal(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}	
	
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
		},
	}

	tellerURL := os.Getenv("TELLER_URL")
	token := os.Getenv("wells_fargo_token")


	return &ClientService{
		Client: client,
		BaseURL: tellerURL,
		Token: token,
	}
}

