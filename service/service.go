package service

import (
	"context"
	"process-payment/models"
	"process-payment/pkg/response"
)

type Payment interface {
	CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, response.ErrorResponse)
	UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse
	GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse)
	GetTransactionByID(ctx context.Context, id int) (models.Transaction, response.ErrorResponse)
	GetTransactions(ctx context.Context) ([]models.Transaction, response.ErrorResponse)
}

type User interface {
	Register(ctx context.Context, payload models.RegisterRequest) response.ErrorResponse
	Login(ctx context.Context, phone string) (string, response.ErrorResponse)
}
