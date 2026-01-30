package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func GenerateMpesaHashedPassword(shortcode, passkey, timestamp string) (string, error) {
	if shortcode == "" || passkey == "" || timestamp == "" {
		return "", errors.New("shortcode, passkey, and timestamp must not be empty")
	}

	concat := shortcode + passkey + timestamp
	hash := sha256.Sum256([]byte(concat))

	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", hash)))

	return encoded, nil
}

func ValidatePhone(phone string) error {
	var trimmedPhone string
	if strings.HasPrefix(phone, "+2517") {
		trimmedPhone = phone[1:]
	} else if strings.HasPrefix(phone, "2517") {
		trimmedPhone = phone
	} else if strings.HasPrefix(phone, "07") {
		trimmedPhone = "251" + phone[1:]
	} else if strings.HasPrefix(phone, "7") {
		trimmedPhone = "251" + phone
	}

	re := regexp.MustCompile(`^[0-9]+$`)

	if !re.MatchString(phone) || !strings.HasPrefix(trimmedPhone, "2517") || len(trimmedPhone) != 12 {
		return fmt.Errorf("invalid phone: %s", phone)
	}

	return nil
}
