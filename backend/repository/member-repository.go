package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type MemberRepo interface {
	Create(member models.Member, tx *gorm.DB) (models.Member, error)
	FindOne(phone_number string) (models.Member, error)
}

type memberRepo struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepo {
	return &memberRepo{db}
}

func (r *memberRepo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *memberRepo) Create(member models.Member, tx *gorm.DB) (models.Member, error) {
	db := r.getDB(tx)
	if err := db.Create(&member).Error; err != nil {
		return models.Member{}, err
	}

	return member, nil
}

func (r *memberRepo) FindOne(phone_number string) (models.Member, error) {
	var data models.Member
	err := r.db.Where("phone_number = ?", phone_number).First(&data).Error

	return data, err
}
