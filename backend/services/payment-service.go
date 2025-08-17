package services

import (
	"encoding/base64"
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
