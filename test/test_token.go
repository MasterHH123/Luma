package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	client *http.Client 
	tellerURL string
	token string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("There was an error loading the .env file: %v\n", err)
	}
	
	cert, err := tls.LoadX509KeyPair("../certificate.pem", "../private_key.pem")
	if err != nil{
		fmt.Printf("Coudln't load cert and key pair: %v\n", err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}	
	
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
		},
	}
	tellerURL = os.Getenv("TELLER_URL")
	token = os.Getenv("wells_fargo_token")

	if tellerURL == "" || token == "" {
		fmt.Printf("Teller URL or token are empty\n")
	} 

}

func callTellerAPI(c *gin.Context) {
	req, err := http.NewRequest("GET", tellerURL, nil)
	if err != nil {
		fmt.Printf("Couldn't prepare request to teller API: %v\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	req.SetBasicAuth(token, "")

	dumpReq, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Printf("Error dumping the Teller API request %v\n", err)
		return
	}
	fmt.Printf("Request being sent to Teller API looks like: ")
	fmt.Println(string(dumpReq))
	fmt.Println("----------------------")


	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Couldn't make request to teller API: %v\n", err)
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.IndentedJSON(resp.StatusCode, gin.H{"error": resp.Body})
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body response: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}


	fmt.Printf("Successfully retrieved data from Teller API\n")
	c.Data(resp.StatusCode, "application/json", body)
}

func main(){
	r := gin.Default()

	r.GET("/", callTellerAPI)

	fmt.Printf("Test Teller API Call starting...\n")
	r.Run(":8080")
}
