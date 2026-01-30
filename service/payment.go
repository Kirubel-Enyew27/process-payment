package service

import (
	"context"
	"database/sql"
	"fmt"
	"process-payment/clients"
	"process-payment/models"
	"process-payment/storage"
	"regexp"
	"strings"
)

func CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, error) {
	if req.Amount < 10 {
		return models.MpesaResponse{}, fmt.Errorf("amount should not be less than 10: %d", req.Amount)
	}

	var trimmedPhone string
	if strings.HasPrefix(req.Phone, "+2517") {
		trimmedPhone = req.Phone[1:]
	} else if strings.HasPrefix(req.Phone, "2517") {
		trimmedPhone = req.Phone
	} else if strings.HasPrefix(req.Phone, "07") {
		trimmedPhone = "251" + req.Phone[1:]
	} else if strings.HasPrefix(req.Phone, "7") {
		trimmedPhone = "251" + req.Phone
	}

	re := regexp.MustCompile(`^[0-9]+$`)

	if !re.MatchString(req.Phone) || !strings.HasPrefix(trimmedPhone, "2517") || len(trimmedPhone) != 12 {
		return models.MpesaResponse{}, fmt.Errorf("invalid phone: %s", req.Phone)
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
