package user

import (
	"process-payment/handlers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine, userHandler handlers.User) {
	router.POST("/user/register", userHandler.Register)
	router.POST("/user/login", userHandler.Login)
	router.POST("/user/verify/:otp", userHandler.VerifyOTP)
}
