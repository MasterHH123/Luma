package handler

import (
	"Luma/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	clientService *service.TellerSevice
}

func NewHandler(svc *service.TellerSevice) *Handler {
	return &Handler{
		clientService: svc,
	}
}

func (h *Handler) EnrollmentDetails(c *gin.Context){
	details, err := h.clientService.GetEnrollmentDetails(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Data(http.StatusOK, "application/json", details)

}
