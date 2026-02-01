package payment

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/utils"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) MpesaWebhook(c *gin.Context) {
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
		err := h.service.UpdateTransactionStatus(models.StatusCompleted, webhook.Envelope.Body.StkCallback.MerchantRequestID)
		if err.Message != "" {
			fmt.Println("failed to update transaction:", err)
		}

		transaction, err := h.service.GetTransactionByReference(webhook.Envelope.Body.StkCallback.MerchantRequestID)
		if err.Message != "" {
			fmt.Println("failed to get updated transaction:", err)
		}

		sms := models.SMSData{
			Phone:   transaction.Phone,
			Message: fmt.Sprintf("You have transaferred amount of %s ETB via M-Pesa successfully ", transaction.Amount),
		}

		if err := utils.SendSMS(sms); err != nil {
			fmt.Printf("failed to send sms: %v", err)
		}

		go response.SendSuccessResponse(c, http.StatusOK, webhook, nil)

	} else {
		err := h.service.UpdateTransactionStatus(models.StatusFailed, webhook.Envelope.Body.StkCallback.MerchantRequestID)
		if err.Message != "" {
			fmt.Println("failed to update transaction:", err)
		}

		go response.SendErrorResponse(c, &err)
	}
}
