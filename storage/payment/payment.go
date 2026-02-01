package payment

import (
	"database/sql"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/storage"

	"go.uber.org/zap"
)

type PaymentStorage struct {
	logger *zap.Logger
	db     *sql.DB
}

func InitPaymentStorage(logger *zap.Logger, db *sql.DB) storage.Payment {
	return &PaymentStorage{
		logger: logger,
		db:     db,
	}
}

func (st *PaymentStorage) SaveTransaction(transaction models.Transaction) response.ErrorResponse {
	sql := `
	INSERT INTO transactions(user_id, amount, phone, reason, reference, status, created_at) 
	VALUES ( $1, $2, $3, $4, $5, $6);
	`
	_, err := st.db.Exec(sql, transaction.UserID, transaction.Amount, transaction.Phone, transaction.Reason, transaction.Reference, transaction.CreatedAt)
	if err != nil {
		st.logger.Error("failed to create transaction", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response.ErrorResponse{}

}

func (st *PaymentStorage) UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse {
	sql := `
	 UPDATE transactions
	 SET status = $1
	 WHERE reference = $2;
	`

	_, err := st.db.Exec(sql, newStatus, reference)
	if err != nil {
		st.logger.Error("failed to update transactio", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response.ErrorResponse{}
}

func (st *PaymentStorage) GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse) {
	row, err := st.db.Query("SELECT * FROM transactons where reference=?", reference)
	if err != nil {
		st.logger.Error("failed to get transaction by reference", zap.Error(err))
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var transaction models.Transaction

	err = row.Scan(&transaction.ID, &transaction.Amount, &transaction.Phone, &transaction.Reason, &transaction.Reference, &transaction.Status, &transaction.CreatedAt)
	if err != nil {
		st.logger.Error("failed to scan the rows", zap.Error(err))
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return transaction, response.ErrorResponse{}
}

func (st *PaymentStorage) GetTransactionByID(id int) (models.Transaction, response.ErrorResponse) {
	row, err := st.db.Query("SELECT * FROM transactons where id=?", id)
	if err != nil {
		st.logger.Error("failed to get transaction by id", zap.Error(err))
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	defer row.Close()

	var transaction models.Transaction

	err = row.Scan(&transaction.ID, &transaction.Amount, &transaction.Phone, &transaction.Reason, &transaction.Reference, &transaction.Status, &transaction.CreatedAt)
	if err != nil {
		st.logger.Error("failed to scan the rows", zap.Error(err))
		return models.Transaction{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return transaction, response.ErrorResponse{}
}

func (st *PaymentStorage) GetTransactions() ([]models.Transaction, response.ErrorResponse) {
	rows, err := st.db.Query("SELECT * FROM transactions;")
	if err != nil {
		st.logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Phone, &transaction.Reason, &transaction.Reference, &transaction.Status, &transaction.CreatedAt)
		if err != nil {
			st.logger.Error("failed to scan the rows", zap.Error(err))
			return nil, response.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		transactions = append(transactions, transaction)
	}

	return transactions, response.ErrorResponse{}

}
