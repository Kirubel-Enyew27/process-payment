package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type PaymentRequest struct {
	Amount    int    `json:"amount"`
	Phone     string `json:"phone"`
	Reason    string `json:"reason"`
	Reference string `json:"reference"`
}

type MpesaRequest struct {
	MerchantRequestID string `json:"MerchantRequestID"`
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int    `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type MpesaResponse struct {
	MerchantRequestID string `json:"MerchantRequestID"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
	ResponseCode      string `json:"ResponseCode"`
	ResponseDesc      string `json:"ResponseDescription"`
	CustomerMessage   string `json:"CustomerMessage"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}
type TokenResponse struct {
	Token
	IssuedAt time.Time
}

type MpesaCallback struct {
	Envelope Result `json:"envelope"`
}

type Result struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  *struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata,omitempty"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

type Transaction struct {
	ID        int             `json:"id"`
	Amount    decimal.Decimal `json:"amount"`
	Phone     string          `json:"phone"`
	Reason    string          `json:"reason"`
	Reference string          `json:"reference"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
}
