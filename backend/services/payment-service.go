package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/util"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"os"
	"time"

	pp "github.com/Frontware/promptpay"
	"gorm.io/gorm"
)

type PaymentService interface {
	CreatePaymentTransactionLog(
		req dto.CreateTxnLogRequest,
	) (*dto.CreateTxnResponse, error)
	GeneratePaymentCodePromptPayment(amount float64) (string, error)
	FindOnePaymentLog(q map[string]interface{}) (*dto.EnquireTxnResponse, error)
	UpdatePaymentStatus(odNumber, status string, tx *gorm.DB) (bool, error)
}

type paymentService struct {
	odRepo  repository.OrderRepo
	payRepo repository.PaymentRepo
}

func NewPaymentService(odRepo repository.OrderRepo, payRepo repository.PaymentRepo) PaymentService {
	return &paymentService{odRepo, payRepo}
}

func (s *paymentService) CreatePaymentTransactionLog(
	req dto.CreateTxnLogRequest,
) (*dto.CreateTxnResponse, error) {

	order, err := s.odRepo.FindOneOrderByOrderNumber(req.OrderNumber)
	if err != nil {
		return nil, err
	}

	err = s.payRepo.CancelActivePaymentLog(order.OrderNumber, nil)
	if err != nil {
		return nil, err
	}

	payload, err := s.GeneratePaymentCodePromptPayment(order.Total)
	if err != nil {
		return nil, err
	}

	signature := signPayload(payload)

	log, err := s.payRepo.CreatePaymentLog(models.PaymentOrderTransactionLog{
		OrderID:           order.ID,
		OrderNumberRef:    order.OrderNumber,
		Amount:            order.Total,
		TransactionNumber: util.GenerateTransactionNumber(req.OrderNumber),
		Status:            models.OrderStatus.ToPay, // initial status
		PaymentCode:       payload,
		QRSignature:       signature,
		ExpiredAt:         time.Now().Add(10 * time.Minute), // expired in 10 min
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

func (s *paymentService) FindOnePaymentLog(q map[string]interface{}) (*dto.EnquireTxnResponse, error) {
	log, err := s.payRepo.FindOneTransaction(q)
	if err != nil {
		return nil, err
	}

	return &dto.EnquireTxnResponse{
		TransactionNumber: log.TransactionNumber,
		OrderNumberRef:    log.OrderNumberRef,
		Amount:            log.Amount,
		Status:            log.Status,
		PaymentCode:       log.PaymentCode,
		ExpiredAt:         log.ExpiredAt,
		CreatedAt:         log.CreatedAt,
		Order: dto.EnquireTxnWithOrderResponse{
			ID:          log.Order.ID,
			OrderNumber: log.Order.OrderNumber,
			Status:      log.Order.Status,
			Total:       log.Order.Total,
		},
	}, nil
}

func (s *paymentService) UpdatePaymentStatus(odNumber, status string, tx *gorm.DB) (bool, error) {
	q := map[string]interface{}{
		"order_number_ref": odNumber,
		"status":           models.OrderStatus.ToPay,
	}

	_, err := s.payRepo.UpdatePaymentLog(q, models.PaymentOrderTransactionLog{
		Status:    status,
		ExpiredAt: time.Now(),
	}, tx)

	if err != nil {
		return false, err
	}

	return true, nil
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
		Amount:      math.Round(amount*100) / 100,
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
