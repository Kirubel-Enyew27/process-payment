package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
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
