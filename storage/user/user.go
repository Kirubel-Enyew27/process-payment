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

func (usr *UserStorage) GetUserByPhone(phone string) (models.User, response.ErrorResponse) {
	row, err := usr.db.Query("SELECT * FROM users where phone=?", phone)
	if err != nil {
		return models.User{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user by phone",
		}
	}
	defer row.Close()

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Phone, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		return user, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to scan user row",
		}
	}

	return user, response.ErrorResponse{}
}
