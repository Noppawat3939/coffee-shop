package services

import (
	"backend/models"
	"backend/repository"
	"fmt"
	"slices"

	"gorm.io/gorm"
)

type OrderService interface {
	CreateMenuVariations(data []models.OrderMenuVariation, order models.Order, tx *gorm.DB) (bool, error)
	CreateLog(order models.Order, tx *gorm.DB) (bool, error)
	UpdateOrderStatusAndLog(odNumber, status string, tx *gorm.DB) (models.Order, error)
}

type orderService struct {
	repo repository.OrderRepo
}

func NewOrderService(repo repository.OrderRepo) OrderService {
	return &orderService{repo}
}

func (s *orderService) CreateMenuVariations(data []models.OrderMenuVariation, order models.Order, tx *gorm.DB) (bool, error) {
	for i := range data {
		data[i].OrderID = order.ID
		if _, err := s.repo.CreateOrderMenuVariation(data[i], tx); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *orderService) CreateLog(order models.Order, tx *gorm.DB) (bool, error) {
	if _, err := s.repo.CreateOrderStatusLog(models.OrderStatusLog{
		OrderID: order.ID,
		Status:  order.Status,
	}, tx); err != nil {
		return false, err
	}

	return true, nil
}

func (s *orderService) UpdateOrderStatusAndLog(odNumber, status string, tx *gorm.DB) (models.Order, error) {
	q := map[string]interface{}{"order_number": odNumber}
	data := models.Order{Status: status}

	order, err := s.repo.FindOneOrderByOrderNumber(odNumber)
	if err != nil {
		return order, err
	}

	allowed, ok := mappingAllowedStatusToUpdate[order.Status]
	if !ok || !slices.Contains(allowed, status) {
		return order, fmt.Errorf("current status not allowed to update")
	}

	_, err = s.repo.UpdateOrder(q, data, tx)
	if err != nil {
		return order, err
	}

	// create order status log
	if _, err := s.repo.CreateOrderStatusLog(models.OrderStatusLog{OrderID: order.ID, Status: status}, tx); err != nil {
		return order, err
	}

	return order, nil
}

var mappingAllowedStatusToUpdate = map[string][]string{
	models.OrderStatus.ToPay:    {models.OrderStatus.Paid, models.OrderStatus.Canceled},
	models.OrderStatus.Paid:     {models.OrderStatus.Paid},
	models.OrderStatus.Canceled: {models.OrderStatus.Canceled},
}
