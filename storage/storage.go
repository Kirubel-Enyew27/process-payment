package storage

import (
	"process-payment/models"
	"process-payment/pkg/response"
)

type Payment interface {
	SaveTransaction(transaction models.Transaction) response.ErrorResponse
	UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse
	GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse)
}
