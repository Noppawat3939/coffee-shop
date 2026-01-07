package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"

	pp "github.com/Frontware/promptpay"
	qrcode "github.com/skip2/go-qrcode"
)

func GeneratePromptPayQR(amount float64) (string, error) {
	paymentInfo := pp.PromptPay{
		PromptPayID: os.Getenv("PROMPTPAY_PHONE"),
		Amount:      amount,
	}

	qrStr, err := paymentInfo.Gen()
	if err != nil {
		return "", err
	}

	png, err := qrcode.Encode(qrStr, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	result := base64.StdEncoding.EncodeToString(png)

	return result, nil
}

func GeneratePaymentCodePromptPayment(amount float64) (string, error) {
	paymentInfo := pp.PromptPay{
		PromptPayID: os.Getenv("PROMPTPAY_PHONE"),
		Amount:      amount,
	}

	qrStr, err := paymentInfo.Gen()
	if err != nil {
		return "", err
	}

	return qrStr, nil
}

func SignPayload(payload string) string {
	secret := os.Getenv("QR_SECRET")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
