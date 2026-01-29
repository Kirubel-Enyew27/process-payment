package main

import (
	"process-payment/db"
	"process-payment/handlers"

	"github.com/gin-gonic/gin"
)

func Init() {
	db.Connect()
}

func main() {
	router := gin.Default()
	router.POST("/payment/process", handlers.CreatePayment)

	router.Run(":8080")
}
