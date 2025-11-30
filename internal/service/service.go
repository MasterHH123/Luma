package service

import (
	"Luma/config"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TellerSevice struct {
	clientService *config.ClientService
}

func NewTellerService() (*TellerSevice) {
	clientService := config.InitTellerService()

	return &TellerSevice{
		clientService: clientService,
	}
}

func (s *TellerSevice) GetEnrollmentDetails(c *gin.Context) ([]byte, error) {
	/*	
	if err := c.ShouldBind(&enrollmentToken); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	*/

	req, err := http.NewRequest("GET", s.clientService.BaseURL, nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return nil, fmt.Errorf("failed to create request", err)
	}

	req.SetBasicAuth(s.clientService.Token, "")


	resp, err := s.clientService.Client.Do(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return nil, fmt.Errorf("Failed to send response to Teller's API: %v\n", err)
	}
	defer resp.Body.Close()

	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil, fmt.Errorf("Error reading response body\n", err)
	}

	if resp.StatusCode != 200{
		c.IndentedJSON(resp.StatusCode, gin.H{"error": body})
		return nil, fmt.Errorf("Response from Teller's API was not successful: %s\n", string(body))
	}

	log.Println("Received details from Teller API")
	return body, nil
}
