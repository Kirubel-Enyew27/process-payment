package user

import (
	"database/sql"
	"process-payment/storage"

	"go.uber.org/zap"
)

type UserStorage struct {
	logger *zap.Logger
	db     *sql.DB
}

func InitUserStorage(logger *zap.Logger, db *sql.DB) storage.User {
	return &UserStorage{
		logger: logger,
		db:     db,
	}
}
