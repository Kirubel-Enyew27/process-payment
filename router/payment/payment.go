package payment

import (
	"process-payment/handlers"

	"github.com/gin-gonic/gin"
)

func InitPaymentRoutes(router *gin.Engine, paymentHandler handlers.Payment) {
	router.POST("/payment/process", paymentHandler.CreatePayment)
	router.GET("/transactions/:id", paymentHandler.GetTransactionByID)
	router.GET("/transactons", paymentHandler.GetTransactions)

}
