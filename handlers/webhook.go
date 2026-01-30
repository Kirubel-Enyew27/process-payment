package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"process-payment/db"
	"process-payment/models"
	"process-payment/service"

	"github.com/gin-gonic/gin"
)

func MpesaWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var webhook models.MpesaCallback

	if err := json.Unmarshal(body, &webhook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if webhook.Envelope.Body.StkCallback.ResultCode == 0 {
		err := service.UpdateTransactionStatus(db.DB, "completed", webhook.Envelope.Body.StkCallback.MerchantRequestID)
		if err != nil {
			fmt.Println("Failed to update transaction:", err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "webhook accepted successfully"})
		return
	}
}
