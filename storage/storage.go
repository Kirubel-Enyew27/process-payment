package storage

import "process-payment/models"

type Payment interface {
	SaveTransaction(transaction models.Transaction) error
	UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) error
	GetTransactionByReference(reference string) (models.Transaction, error)
}
