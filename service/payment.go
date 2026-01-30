package service

import (
	"context"
	"fmt"
	"process-payment/clients"
	"process-payment/models"
	"process-payment/storage"
	"process-payment/utils"
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Service struct {
	logger  *zap.Logger
	storage storage.Payment
}

func InitService(logger *zap.Logger, storage storage.Payment) Payment {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}

func (srv *Service) CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, error) {
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

	err = srv.storage.SaveTransaction(models.Transaction{
		Amount:    decimal.NewFromInt(int64(req.Amount)),
		Phone:     req.Phone,
		Reason:    req.Reason,
		Reference: req.Reference,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (srv *Service) UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) error {
	err := srv.storage.UpdateTransactionStatus(newStatus, reference)
	if err != nil {
		return err
	}

	return nil
}

func (srv *Service) GetTransactionByReference(reference string) (models.Transaction, error) {
	transaction, err := srv.storage.GetTransactionByReference(reference)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err
}
