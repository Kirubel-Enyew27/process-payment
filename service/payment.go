package service

import (
	"context"
	"database/sql"
	"fmt"
	"process-payment/clients"
	"process-payment/models"
	"process-payment/storage"
	"process-payment/utils"
)

func CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, error) {
	if req.Amount < 10 {
		return models.MpesaResponse{}, fmt.Errorf("amount should not be less than 10: %d", req.Amount)
	}

	if err := utils.ValidatePhone(req.Phone); err != nil {
		return models.MpesaResponse{}, err
	}

	resp, err := clients.CreatePayment(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func UpdateTransactionStatus(db *sql.DB, newStatus models.PaymentStatus, reference string) error {
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
