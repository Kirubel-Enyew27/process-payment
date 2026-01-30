package user

import (
	"context"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/service"
	"process-payment/storage"
	"process-payment/utils"
	"time"

	"go.uber.org/zap"
)

type UserService struct {
	logger  *zap.Logger
	storage storage.User
}

func InitUserService(logger *zap.Logger, storage storage.User) service.User {
	return &UserService{
		logger:  logger,
		storage: storage,
	}
}

func (u *UserService) Register(ctx context.Context, payload models.RegisterRequest) response.ErrorResponse {
	user := models.User{
		Username:  payload.Username,
		Phone:     payload.Phone,
		Role:      payload.Role,
		CreatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		u.logger.Error("failed to validate user data", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	return u.storage.Register(user)

}

func (u *UserService) Login(ctx context.Context, phone string) (string, response.ErrorResponse) {
	if err := utils.ValidatePhone(phone); err != nil {
		u.logger.Error("failed to validate phone", zap.Error(err))
		return "", response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	
	return "", response.ErrorResponse{}
}
