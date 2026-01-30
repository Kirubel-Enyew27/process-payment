package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"process-payment/models"

	"github.com/spf13/viper"
)

func SendSMS(data models.SMSData) error {
	token := viper.GetString("SMS_API_KEY")
	shortcode := viper.GetString("SMS_SHORT_CODE")
	url := viper.GetString("SMS_API_URL")

	if token == "" || shortcode == "" {
		fmt.Println("SMS_TOKEN or SMS_SHORTCODE is not available!")
		return errors.New("SMS_TOKEN or SMS_SHORTCODE is not available!")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("token", token)
	_ = writer.WriteField("phone", data.Phone)
	_ = writer.WriteField("msg", data.Message)
	_ = writer.WriteField("shortcode_id", shortcode)

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		fmt.Printf("Failed to create SMS request: %v", err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send SMS request: %v", err)
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read SMS response: %v", err)
		return err
	}

	var smsResp models.SMSResponse
	if err := json.Unmarshal(respBody, &smsResp); err != nil {
		fmt.Printf("Failed to parse SMS response: %v", err)
		return fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Printf("SMS Response: %+v\n", smsResp)

	if smsResp.Error {
		fmt.Printf("SMS failed: %v", smsResp.Message)
		return fmt.Errorf("SMS failed:%s", smsResp.Message)
	}

	return nil
}
