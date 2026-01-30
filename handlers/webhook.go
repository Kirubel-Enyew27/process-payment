package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"process-payment/db"
	"process-payment/models"
	"process-payment/service"
	"process-payment/utils"

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
			fmt.Println("failed to update transaction:", err)
		}

		transaction, err := service.GetTransactionByReference(db.DB, webhook.Envelope.Body.StkCallback.MerchantRequestID)
		if err != nil {
			fmt.Println("failed to get updated transaction:", err)
		}

		sms := models.SMSData{
			Phone:   transaction.Phone,
			Message: fmt.Sprintf("You have transaferred amount of %s via M-Pesa successfully ", transaction.Amount),
		}

		if err := utils.SendSMS(sms); err != nil {
			fmt.Printf("failed to send sms: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "webhook accepted successfully"})
		return
	}
}
