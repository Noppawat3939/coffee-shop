package services

import (
	"backend/models"
	"backend/repository"
	"fmt"
)

const MIN_ORDER_TOTAL = 100

type MemberPointService interface {
	NewMemberPoint(data models.MemberPoint) (bool, error)
	CalculateEarnPoint(total float64) int
	FormatPoint(point int) string
}

type memberPointService struct {
	pointRepo repository.MemberPointRepo
}

func NewMemberPointService(pointRepo repository.MemberPointRepo) MemberPointService {
	return &memberPointService{pointRepo}
}

func (s *memberPointService) NewMemberPoint(data models.MemberPoint) (bool, error) {

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
