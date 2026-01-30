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

func UpdateTransactionStatus(db *sql.DB, newStatus string, reference string) error {
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
