package user

import (
	"process-payment/handlers"
	"process-payment/service"
	"time"

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
