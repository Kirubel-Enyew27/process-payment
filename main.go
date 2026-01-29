package main

import (
	"process-payment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/payment/process", handlers.CreatePayment)

	router.Run(":8080")
}
