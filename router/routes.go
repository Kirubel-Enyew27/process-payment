package router

import (
	"process-payment/handlers"
	"process-payment/router/payment"
	"process-payment/router/user"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, paymentHandler handlers.Payment, userHandler handlers.User) {
	// initialize paymnet routes
	payment.InitPaymentRoutes(router, paymentHandler)
	// initialize user routes
	user.InitUserRoutes(router, userHandler)

}
