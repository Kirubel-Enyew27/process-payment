package user

import (
	"database/sql"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
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

func (usr *UserStorage) Register(user models.User) response.ErrorResponse {
	sql := `INSERT INTO users(username, phone, role, status, created_at) 
	VALUES (&1, &2, &3, $4, $5);`

	_, err := usr.db.Exec(sql, user.Username, user.Phone, user.Role, user.Status, user.CreatedAt)
	if err != nil {
		usr.logger.Error("failed to register user", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response.ErrorResponse{}
}
