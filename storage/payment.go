package storage

import (
	"database/sql"
	"fmt"
	"process-payment/models"
)

func SaveTransaction(db *sql.DB, transaction models.Transaction) error {
	sql := `
	INSERT INTO transactions(amount, phone, reason, reference, created_at) VALUES (
     $1, $2, $3, $4, $5
	)
	`
	_, err := db.Exec(sql, transaction.Amount, transaction.Phone, transaction.Reason, transaction.Reference, transaction.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save transaction: %v", err)
	}

	return nil

}
