package handlers

import (
	"context"
	"net/http"
	"process-payment/models"
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*30)
	defer cancel()

	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreatePayment(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})

}
