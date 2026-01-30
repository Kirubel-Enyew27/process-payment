package handlers

import (
	"context"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/service"
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
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.CreatePayment(ctx, req)
	if err.Message != "" {
		response.SendErrorResponse(c, &err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})

}
