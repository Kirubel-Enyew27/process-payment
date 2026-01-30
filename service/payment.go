package service

import (
	"database/sql"
	"process-payment/storage"
)

func UpdateTransactionStatus(db *sql.DB, newStatus string, reference string) error {
	err := storage.UpdateTransactionStatus(db, newStatus, reference)
	if err != nil {
		return err
	}

	return nil
}
