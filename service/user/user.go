package user

import (
	"context"
	"fmt"
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

	user, errResp := u.storage.GetUserByPhone(phone)
	if errResp.Message != "" {
		return "", errResp
	}

	otp, err := utils.GenerateUniqueOTP("0123456789", 4)
	if err != nil {
		u.logger.Error("failed to generate otp", zap.Error(err))
		return otp, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to generate otp " + err.Error(),
		}
	}

	sms := models.SMSData{
		Phone:   phone,
		Message: fmt.Sprintf("Your 4-digit verification code is %v, enter the code to login", otp),
	}

	if err := utils.SendSMS(sms); err != nil {
		u.logger.Error("failed to send otp", zap.Error(err))
		return otp, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to send otp" + err.Error(),
		}
	}

	models.UserSession[otp] = user

	return otp, response.ErrorResponse{}
}

func (u *UserService) VerifyOTP(ctx context.Context, otp string) (string, response.ErrorResponse) {
	user, exists := models.UserSession[otp]
	if !exists {
		u.logger.Error("user not found associated with this OTP", zap.String("OTP", otp))
		return "", response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "user not found associated wuth this otp",
		}
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		u.logger.Error("failed to generate access token")
		return "", response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to generate access token",
		}
	}

	session := models.Session{
		ID:        user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24),
		CreatedAt: time.Now(),
	}

	errResp := u.storage.LoginSession(session)
	if errResp.Message != "" {
		return "", errResp
	}

	return token, response.ErrorResponse{}

}
