package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/util"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	pp "github.com/Frontware/promptpay"
)

type PaymentService interface {
	CreatePaymentTransactionLog(
		req dto.CreateTxnLogRequest,
	) (*dto.CreateTxnResponse, error)
	GeneratePaymentCodePromptPayment(amount float64) (string, error)
}

type paymentService struct {
	repo repository.OrderRepo
}

func NewPaymentService(repo repository.OrderRepo) PaymentService {
	return &paymentService{
		repo: repo,
	}
}

func (s *paymentService) CreatePaymentTransactionLog(
	req dto.CreateTxnLogRequest,
) (*dto.CreateTxnResponse, error) {

	order, err := s.repo.FindOneOrderByOrderNumber(req.OrderNumber)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CancelActivePaymentLog(int(order.ID)); err != nil {
		return nil, err
	}

	payload, err := s.GeneratePaymentCodePromptPayment(order.Total)
	if err != nil {
		return nil, err
	}

	signature := signPayload(payload)

	log, err := s.repo.CreatePaymentLog(models.PaymentOrderTransactionLog{
		OrderID:           order.ID,
		Amount:            order.Total,
		TransactionNumber: util.GenerateTransactionNumber(req.OrderNumber),
		Status:            models.OrderStatus.ToPay, // initial status
		PaymentCode:       payload,
		QRSignature:       signature,
		ExpiredAt:         time.Now().Add(2 * time.Minute), // expired in 2 min
	}, nil)

	if err != nil {
		return nil, err
	}

	return &dto.CreateTxnResponse{
		TransactionNumber: log.TransactionNumber,
		Amount:            log.Amount,
		Status:            log.Status,
		PaymentCode:       log.PaymentCode,
		ExpiredAt:         log.ExpiredAt,
		CreatedAt:         log.CreatedAt,
	}, nil

}

// func GeneratePromptPayQR(amount float64) (string, error) {
// 	paymentInfo := pp.PromptPay{
// 		PromptPayID: os.Getenv("PROMPTPAY_PHONE"),
// 		Amount:      amount,
// 	}

// 	qrStr, err := paymentInfo.Gen()
// 	if err != nil {
// 		return "", err
// 	}

// 	png, err := qrcode.Encode(qrStr, qrcode.Medium, 256)
// 	if err != nil {
// 		return "", err
// 	}

// 	result := base64.StdEncoding.EncodeToString(png)

// 	return result, nil
// }

func (s *paymentService) GeneratePaymentCodePromptPayment(amount float64) (string, error) {
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

func signPayload(payload string) string {
	secret := os.Getenv("QR_SECRET")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
