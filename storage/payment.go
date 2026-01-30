package storage

import (
	"database/sql"
	"fmt"
	"process-payment/models"

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

func (st *Storage) SaveTransaction(transaction models.Transaction) error {
	sql := `
	INSERT INTO transactions(amount, phone, reason, reference, status, created_at) 
	VALUES ( $1, $2, $3, $4, $5, $6);
	`
	_, err := st.db.Exec(sql, transaction.Amount, transaction.Phone, transaction.Reason, transaction.Reference, transaction.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save transaction: %v", err)
	}

	return nil

}

func (st *Storage) UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) error {
	sql := `
	 UPDATE transactions
	 SET status = $1
	 WHERE reference = $2;
	`

	_, err := st.db.Exec(sql, newStatus, reference)
	if err != nil {
		return fmt.Errorf("failed to update transaction status: %v", err)
	}

	return nil
}

func (st *Storage) GetTransactionByReference(reference string) (models.Transaction, error) {
	row, err := st.db.Query("SELECT * FROM transactons where reference=?", reference)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("failed to get transaction by reference: %v", err)
	}

	var transaction models.Transaction

	err = row.Scan(&transaction.ID, &transaction.Amount, &transaction.Phone, &transaction.Reason, &transaction.Reference, &transaction.Status, &transaction.CreatedAt)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("failed to scan transaction: %v", err)
	}

	return transaction, nil
}
