package models

type SMSData struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
type SMSResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}
