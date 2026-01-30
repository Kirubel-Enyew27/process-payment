package user

import (
	"process-payment/storage"

	"go.uber.org/zap"
)

type UserService struct {
	logger  *zap.Logger
	storage storage.User
}

func InitUserService(logger *zap.Logger, storage storage.User) storage.User {
	return &UserService{
		logger:  logger,
		storage: storage,
	}
}
