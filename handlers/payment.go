package handlers

import (
	"context"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	timeout time.Duration
	service service.Payment
}

func InitHandler(logger *zap.Logger, timeout time.Duration, service service.Payment) Payment {
	return &Handler{
		logger:  logger,
		timeout: timeout,
	}
}

func (h *Handler) CreatePayment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	resp, err := h.service.CreatePayment(ctx, req)
	if err.Message != "" {
		response.SendErrorResponse(c, &err)
		return
	}

	response.SendSuccessResponse(c, http.StatusCreated, resp, nil)

}

func (h *Handler) GetTransactionByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	if id == 0 {
		err := response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "transaction id is not provided",
		}
		response.SendErrorResponse(c, &err)
		return
	}

	resp, err := h.service.GetTransactionByID(ctx, id)
	if err.Message != "" {
		response.SendErrorResponse(c, &err)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, resp, nil)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	resp, err := h.service.GetTransactions(ctx)
	if err.Message != "" {
		response.SendErrorResponse(c, &err)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, resp, nil)
}
