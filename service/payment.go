package service

import (
	"database/sql"
	"process-payment/models"
	"process-payment/storage"
)

func UpdateTransactionStatus(db *sql.DB, newStatus string, reference string) error {
	err := storage.UpdateTransactionStatus(db, newStatus, reference)
	if err != nil {
		return err
	}

	return nil
}

func GetTransactionByReference(db *sql.DB, reference string) (models.Transaction, error) {
	transaction, err := storage.GetTransactionByReference(db, reference)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err
}
