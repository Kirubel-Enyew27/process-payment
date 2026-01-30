package router

import (
	"process-payment/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, handler handlers.Payment) {
	router.POST("/payment/process", handler.CreatePayment)

}
