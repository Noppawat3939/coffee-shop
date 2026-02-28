package repository

import (
	"backend/models"
	"backend/util"

	"gorm.io/gorm"
)

type MemberRepo interface {
	Create(member models.Member) (models.Member, error)
	FindOne(phone_number string) (models.Member, error)
	FindAllIncluded(filter models.MemberFilter, page, limit int) ([]models.Member, error)
}

type memberRepo struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepo {
	return &memberRepo{db}
}

func (r *memberRepo) Create(member models.Member) (models.Member, error) {

	if err := r.db.Create(&member).Error; err != nil {
		return models.Member{}, err
	}

	return member, nil
}

func (r *memberRepo) FindOne(phone_number string) (models.Member, error) {
	var data models.Member
	err := r.db.Where("phone_number = ?", phone_number).First(&data).Error

	return data, err
}

func (r *memberRepo) FindAllIncluded(filter models.MemberFilter, page, limit int) ([]models.Member, error) {
	var data []models.Member
	pagination := util.Pagination{Page: page, Limit: limit}

	db := r.db.Model(&models.Member{})

	if filter.PhoneNumber != "" {
		db = db.Where("phone_number = ?", filter.PhoneNumber)
	}

	if filter.FullName != "" {
		db = db.Where("full_name ILIKE ?", "%"+filter.FullName+"%")
	}

	err := db.Preload("MemberPoint").Order("id DESC").Scopes(pagination.GetPaginationResult).Find(&data).Error

	return data, err
}
