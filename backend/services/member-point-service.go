package services

import (
	"backend/models"
	"backend/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const MIN_ORDER_TOTAL = 100

type MemberPointService interface {
	CreateMemberPoint(data models.MemberPoint, tx *gorm.DB) (bool, error)
	CalculateEarnPoint(total float64) (bool, int)
	FormatPoint(point int) string
	EarnPointFromOrder(order models.Order, tx *gorm.DB) error
}

type memberPointService struct {
	pointRepo repository.MemberPointRepo
}

func NewMemberPointService(pointRepo repository.MemberPointRepo) MemberPointService {
	return &memberPointService{pointRepo}
}

func (s *memberPointService) CreateMemberPoint(data models.MemberPoint, tx *gorm.DB) (bool, error) {
	_, err := s.pointRepo.FindOneMemberPoint(data.MemberID)

	if err == nil {
		// already exists
		return false, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// real DB error
		return false, err
	}

	_, err = s.pointRepo.CreateMemberPoint(models.MemberPoint{MemberID: data.MemberID, TotalPoints: data.TotalPoints}, tx)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (s *memberPointService) CalculateEarnPoint(total float64) (bool, int) {
	if total < MIN_ORDER_TOTAL {
		return false, 0
	}

	return true, int(total)
}

func (s *memberPointService) FormatPoint(point int) string {
	return fmt.Sprintf("%.2f", float64(point)/100)
}

func (s *memberPointService) EarnPointFromOrder(order models.Order, tx *gorm.DB) error {
	ok, points := s.CalculateEarnPoint(order.Total)

	if !ok {
		return nil
	}

	log := models.MemberPointLog{
		MemberID: order.MemberID,
		OrderID:  &order.ID,
		Type:     models.MemberPointLogType.Earn,
		Points:   points,
	}

	_, err := s.pointRepo.CreatePointLog(log, tx)

	if err != nil {
		return err
	}

	_, err = s.pointRepo.IncreaseMemberPoint(log.MemberID, points, tx)
	if err != nil {
		return err
	}

	return nil
}
