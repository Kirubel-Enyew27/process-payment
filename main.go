package main

import (
	"log"
	"process-payment/db"
	paymentHandler "process-payment/handlers/payment"
	"process-payment/router"
	paymentSvc "process-payment/service/payment"
	"process-payment/storage/payment"

	userHandler "process-payment/handlers/user"
	userSvc "process-payment/service/user"
	"process-payment/storage/user"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	viper.AutomaticEnv()
	DB, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	logger, _ := zap.NewDevelopment()
	timeout := time.Second * 30

	// Payment
	paymentStorage := payment.InitPaymentStorage(logger, DB)
	paymentService := paymentSvc.InitPaymentService(logger, paymentStorage)
	paymentHandler := paymentHandler.InitPaymentHandler(logger, timeout, paymentService)

	// User
	userStorage := user.InitUserStorage(logger, DB)
	userService := userSvc.InitUserService(logger, userStorage)
	userHandler := userHandler.InitUserHandler(logger, timeout, userService)

	router.SetUpRoutes(r, paymentHandler, userHandler)

	r.Run(":8080")
}
