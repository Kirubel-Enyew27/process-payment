package user

import (
	"context"
	"net/http"
	"process-payment/handlers"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger  *zap.Logger
	timeout time.Duration
	service service.User
}

func InitUserHandler(logger *zap.Logger, timeout time.Duration, service service.User) handlers.User {
	return &UserHandler{
		logger:  logger,
		timeout: timeout,
		service: service,
	}
}

func (u *UserHandler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), u.timeout)
	defer cancel()

	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Error("failed to bind register user request", zap.Error(err))
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	errResp := u.service.Register(ctx, req)
	if errResp.Message != "" {
		response.SendErrorResponse(c, &errResp)
		return
	}

	response.SendSuccessResponse(c, http.StatusCreated, nil, nil)

}

func (u UserHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), u.timeout)
	defer cancel()

	phone := c.Param("phone")
	if phone == "" {
		u.logger.Error("phone is not provided")
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "phone is not provided",
		})
	}

	otp, errResp := u.service.Login(ctx, phone)
	if errResp.Message != "" {
		response.SendErrorResponse(c, &errResp)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, otp, nil)

}

func (u UserHandler) VerifyOTP(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), u.timeout)
	defer cancel()

	otp := c.Param("otp")
	if otp == "" {
		u.logger.Error("OTP is not prvoided")
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "OTP is not provided",
		})
	}

	token, errResp := u.service.VerifyOTP(ctx, otp)
	if errResp.Message != "" {
		response.SendErrorResponse(c, &errResp)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, token, nil)
}
