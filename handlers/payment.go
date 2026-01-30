package handlers

import (
	"context"
	"net/http"
	"process-payment/models"
	"process-payment/service"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*30)
	defer cancel()

	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.CreatePayment(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp})

}
