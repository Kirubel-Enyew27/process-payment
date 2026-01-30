package storage

import (
	"database/sql"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"

	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
	db     *sql.DB
}

func InitStorage(logger *zap.Logger, db *sql.DB) Payment {
	return &Storage{
		logger: logger,
		db:     db,
	}
}

func (st *Storage) SaveTransaction(transaction models.Transaction) response.ErrorResponse {
	sql := `
	INSERT INTO transactions(amount, phone, reason, reference, status, created_at) 
	VALUES ( $1, $2, $3, $4, $5, $6);
	`
	_, err := st.db.Exec(sql, transaction.Amount, transaction.Phone, transaction.Reason, transaction.Reference, transaction.CreatedAt)
	if err != nil {
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response.ErrorResponse{}

}

func (st *Storage) UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse {
	sql := `
	 UPDATE transactions
	 SET status = $1
	 WHERE reference = $2;
	`

	_, err := st.db.Exec(sql, newStatus, reference)
	if err != nil {
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response.ErrorResponse{}
}

func (st *Storage) GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse) {
	row, err := st.db.Query("SELECT * FROM transactons where reference=?", reference)
	if err != nil {
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var transaction models.Transaction

	err = row.Scan(&transaction.ID, &transaction.Amount, &transaction.Phone, &transaction.Reason, &transaction.Reference, &transaction.Status, &transaction.CreatedAt)
	if err != nil {
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return transaction, response.ErrorResponse{}
}
