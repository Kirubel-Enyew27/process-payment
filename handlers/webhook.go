package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"process-payment/models"

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
		c.JSON(http.StatusOK, gin.H{"msg": "webhook accepted successfully"})
		return
	}
}
