package storage

import (
	"process-payment/models"
	"process-payment/pkg/response"
)

type Payment interface {
	SaveTransaction(transaction models.Transaction) response.ErrorResponse
	UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse
	GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse)
	GetTransactionByID(id int) (models.Transaction, response.ErrorResponse)
	GetTransactions() ([]models.Transaction, response.ErrorResponse)
}

type User interface{}
