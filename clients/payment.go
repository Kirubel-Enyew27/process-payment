package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"process-payment/models"
	"process-payment/pkg/response"
	"process-payment/utils"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func CreatePayment(ctx context.Context, payload models.PaymentRequest, logger *zap.Logger) (models.MpesaResponse, response.ErrorResponse) {
	request, err := prepareRequest(payload)
	if err != nil {
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to marshal request payload",
		}
	}

	mpesa_base_url := viper.GetString("MPESA_BASE_URL")
	url := mpesa_base_url + "/mpesa/stkpush/v3/processrequest"

	token, err := GetAccessToken()
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		}
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Error("failed to create request", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create request",
		}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("failed to send request", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to send request",
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to read response body",
		}
	}

	logger.Info("response", zap.Any("Response Body", string(respBody)))

	var response2 models.MpesaResponse

	if err := json.Unmarshal(respBody, &response2); err != nil {
		logger.Error("failed to unmarshal response", zap.Error(err))
		return models.MpesaResponse{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to unmarshal response",
		}
	}

	return response2, response.ErrorResponse{}
}

func prepareRequest(payload models.PaymentRequest) (models.MpesaRequest, error) {
	password := viper.GetString("PASSWORD")
	shortCode := viper.GetString("SHORT_CODE")
	callBack := viper.GetString("CALL_BACK")
	timestamp := time.Now().Format("20060102150405")
	hashedPassword, err := utils.GenerateMpesaHashedPassword(shortCode, password, timestamp)
	if err != nil {
		return models.MpesaRequest{}, err
	}

	return models.MpesaRequest{
		MerchantRequestID: payload.Reference,
		BusinessShortCode: shortCode,
		Password:          hashedPassword,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            payload.Amount,
		PartyA:            payload.Phone,
		PartyB:            shortCode,
		PhoneNumber:       payload.Phone,
		CallBackURL:       callBack,
		AccountReference:  payload.Reference,
		TransactionDesc:   "C2B payment",
	}, nil
}

func GetAccessToken() (models.TokenResponse, error) {
	consumerKey := viper.GetString("MPESA_CONSUMER_KEY")
	consumerSecret := viper.GetString("MPESA_CONSUMER_SECRET")
	mpesa_base_url := viper.GetString("MPESA_BASE_URL")
	url := mpesa_base_url + "/v1/token/generate?grant_type=client_credentials"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return models.TokenResponse{}, errors.New("failed to create access token request")
	}
	req.SetBasicAuth(consumerKey, consumerSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.TokenResponse{}, errors.New("failed to send access token response")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TokenResponse{}, errors.New("failed to read access token response body")
	}

	var response models.Token

	if err := json.Unmarshal(respBody, &response); err != nil {
		return models.TokenResponse{}, errors.New("failed to unmarshal access token response")
	}

	token := models.TokenResponse{
		Token:    response,
		IssuedAt: time.Now(),
	}

	return token, nil
}
