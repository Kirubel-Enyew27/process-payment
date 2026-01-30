package service

import (
	"context"
	"process-payment/models"
)

type Payment interface {
	CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, error)
	UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) error
	GetTransactionByReference(reference string) (models.Transaction, error)
}
