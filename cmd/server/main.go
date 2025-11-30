package main

import (
	"Luma/internal/handler"
	"Luma/internal/service"

	"github.com/gin-gonic/gin"
)


func main(){
	
	tellerService := service.NewTellerService()

	tellerHandler := handler.NewHandler(tellerService)

	router := gin.Default()
	router.GET("/", tellerHandler.EnrollmentDetails)

	router.Run(":8080")

}
