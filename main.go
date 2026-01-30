package main

import (
	"process-payment/db"
	"process-payment/handlers"
	"process-payment/service"
	"process-payment/storage"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	db.Connect()

	router := gin.Default()
	logger, _ := zap.NewDevelopment()
	timeout := time.Second * 30
	storage := storage.InitStorage(logger, db.DB)
	service := service.InitService(logger, storage)
	handler := handlers.InitHandler(logger, timeout, service)
	router.POST("/payment/process", handler.CreatePayment)

	router.Run(":8080")
}
