package repository

import (
	"backend/internal/model"
	"backend/pkg/pagination"

	"gorm.io/gorm"
)

type MemberRepo interface {
	Create(member model.Member) (model.Member, error)
	FindOne(phone_number string) (model.Member, error)
	FindAllIncluded(filter model.MemberFilter, page, limit int) ([]model.Member, error)
}

type memberRepo struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepo {
	return &memberRepo{db}
}

func (r *memberRepo) Create(member model.Member) (model.Member, error) {

	if err := r.db.Create(&member).Error; err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func (r *memberRepo) FindOne(phone_number string) (model.Member, error) {
	var data model.Member
	err := r.db.Where("phone_number = ?", phone_number).First(&data).Error

	return data, err
}

func (r *memberRepo) FindAllIncluded(filter model.MemberFilter, page, limit int) ([]model.Member, error) {
	var data []model.Member
	pagination := pagination.Pagination{Page: page, Limit: limit}

	db := r.db.Model(&model.Member{})

	if filter.PhoneNumber != "" {
		db = db.Where("phone_number = ?", filter.PhoneNumber)
	}

	if filter.FullName != "" {
		db = db.Where("full_name ILIKE ?", "%"+filter.FullName+"%")
	}

	err := db.Preload("MemberPoint").Order("id DESC").Scopes(pagination.Apply).Find(&data).Error

	return data, err
}
