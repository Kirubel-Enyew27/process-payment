package storage

import (
	"database/sql"
	"fmt"
	"process-payment/models"
)

func SaveTransaction(db *sql.DB, transaction models.Transaction) error {
	sql := `
	INSERT INTO transactions(amount, phone, reason, reference, status, created_at) 
	VALUES ( $1, $2, $3, $4, $5, $6);
	`
	_, err := db.Exec(sql, transaction.Amount, transaction.Phone, transaction.Reason, transaction.Reference, transaction.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save transaction: %v", err)
	}

	return nil

}

func UpdateTransactionStatus(db *sql.DB, newStatus models.PaymentStatus, reference string) error {
	sql := `
	 UPDATE transactions
	 SET status = $1
	 WHERE reference = $2;
	`

	_, err := db.Exec(sql, newStatus, reference)
	if err != nil {
		return fmt.Errorf("failed to update transaction status: %v", err)
	}

	return nil
}

func GetTransactionByReference(db *sql.DB, reference string) (models.Transaction, error) {
	row, err := db.Query("SELECT * FROM transactons where reference=?", reference)
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
