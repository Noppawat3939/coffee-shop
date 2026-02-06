package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MemberPointRepo interface {
	// Member point
	CreateMemberPoint(data models.MemberPoint) (models.MemberPoint, error)
	FindOneMemberPoint(memberId uint) (models.MemberPoint, error)
	FindMembersPoint() ([]models.MemberPoint, error)
	UpdateMemberPoint(memberId uint, point int) (models.MemberPoint, error)

	// Member point log
	CreatePointLog(data models.MemberPointLog) (models.MemberPointLog, error)
	UpdatePointLog(id uint, data models.MemberPointLog) (models.MemberPointLog, error)
}

type memberPointRepo struct {
	db *gorm.DB
}

func NewMemberPointRepository(db *gorm.DB) MemberPointRepo {
	return &memberPointRepo{db}
}

func (r *memberPointRepo) CreateMemberPoint(data models.MemberPoint) (models.MemberPoint, error) {
	return models.MemberPoint{}, nil
}

func (r *memberPointRepo) FindOneMemberPoint(memberId uint) (models.MemberPoint, error) {
	return models.MemberPoint{}, nil
}

func (r *memberPointRepo) FindMembersPoint() ([]models.MemberPoint, error) {
	return []models.MemberPoint{}, nil
}

func (r *memberPointRepo) UpdateMemberPoint(memberId uint, point int) (models.MemberPoint, error) {
	return models.MemberPoint{}, nil
}

func (r *memberPointRepo) CreatePointLog(data models.MemberPointLog) (models.MemberPointLog, error) {
	return models.MemberPointLog{}, nil
}

func (r *memberPointRepo) UpdatePointLog(id uint, data models.MemberPointLog) (models.MemberPointLog, error) {
	return models.MemberPointLog{}, nil
}
