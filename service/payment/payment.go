package payment

import (
	"context"
	"fmt"
	"net/http"
	"process-payment/clients"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/service"
	"process-payment/storage"
	"process-payment/utils"
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type PaymentService struct {
	logger  *zap.Logger
	storage storage.Payment
}

func InitPaymentService(logger *zap.Logger, storage storage.Payment) service.Payment {
	return &PaymentService{
		logger:  logger,
		storage: storage,
	}
}

func (srv *PaymentService) CreatePayment(ctx context.Context, req models.PaymentRequest) (models.MpesaResponse, response.ErrorResponse) {
	if req.Amount < 10 {
		srv.logger.Error("amount should not be less than 10", zap.Int("amount:", req.Amount))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("amount should not be less than 10: %d", req.Amount),
		}
	}

	if err := utils.ValidatePhone(req.Phone); err != nil {
		srv.logger.Error("phone validaton failed", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	resp, err := clients.CreatePayment(ctx, req, srv.logger)
	if err.Message != "" {
		return resp, err
	}

	err = srv.storage.SaveTransaction(models.Transaction{
		Amount:    decimal.NewFromInt(int64(req.Amount)),
		Phone:     req.Phone,
		Reason:    req.Reason,
		Reference: req.Reference,
		CreatedAt: time.Now(),
	})

	if err.Message != "" {
		return resp, err
	}

	return resp, response.ErrorResponse{}
}

func (srv *PaymentService) UpdateTransactionStatus(newStatus models.PaymentStatus, reference string) response.ErrorResponse {
	err := srv.storage.UpdateTransactionStatus(newStatus, reference)
	if err.Message != "" {
		return err
	}

	return response.ErrorResponse{}
}

func (srv *PaymentService) GetTransactionByReference(reference string) (models.Transaction, response.ErrorResponse) {
	transaction, err := srv.storage.GetTransactionByReference(reference)
	if err.Message != "" {
		return models.Transaction{}, err
	}

	return transaction, err
}

func (srv *PaymentService) GetTransactionByID(ctx context.Context, id int) (models.Transaction, response.ErrorResponse) {
	transaction, err := srv.storage.GetTransactionByID(id)
	if err.Message != "" {
		return models.Transaction{}, err
	}
	return transaction, err

}

func (srv *PaymentService) GetTransactions(ctx context.Context) ([]models.Transaction, response.ErrorResponse) {
	transactions, err := srv.storage.GetTransactions()
	if err.Message != "" {
		return nil, err
	}

	return transactions, response.ErrorResponse{}

}
