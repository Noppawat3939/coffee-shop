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
	NewMemberPoint(data models.MemberPoint, tx *gorm.DB) (bool, error)
	CalculateEarnPoint(total float64) int
	FormatPoint(point int) string
}

type memberPointService struct {
	pointRepo repository.MemberPointRepo
}

func NewMemberPointService(pointRepo repository.MemberPointRepo) MemberPointService {
	return &memberPointService{pointRepo}
}

func (s *memberPointService) NewMemberPoint(data models.MemberPoint, tx *gorm.DB) (bool, error) {
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

func (s *memberPointService) CalculateEarnPoint(total float64) int {
	if total < MIN_ORDER_TOTAL {
		return 0
	}

	return int(total)
}

func (s *memberPointService) FormatPoint(point int) string {
	return fmt.Sprintf("%.2f", float64(point)/100)
}
