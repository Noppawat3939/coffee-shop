package services

import (
	"backend/models"
	"backend/repository"

	"gorm.io/gorm"
)

type OrderService interface {
	UpdateOrderStatusAndLog(odNumber, status string, tx *gorm.DB) (bool, error)
}

type orderService struct {
	repo repository.OrderRepo
}

func NewOrderService(repo repository.OrderRepo) OrderService {
	return &orderService{repo}
}

func (s *orderService) UpdateOrderStatusAndLog(odNumber, status string, tx *gorm.DB) (bool, error) {
	q := map[string]interface{}{"order_number": odNumber}
	data := models.Order{Status: status}

	order, err := s.repo.UpdateOrder(q, data, tx)
	if err != nil {
		return false, err
	}

	// create order status log
	if _, err := s.repo.CreateOrderStatusLog(models.OrderStatusLog{OrderID: order.ID, Status: status}, tx); err != nil {
		return false, err
	}

	return true, nil
}
