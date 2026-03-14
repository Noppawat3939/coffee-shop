package repository

import (
	"backend/internal/model"
	"backend/pkg/pagination"

	"gorm.io/gorm"
)

type MemberPointRepo interface {
	// Member point
	CreateMemberPoint(data model.MemberPoint, tx *gorm.DB) (model.MemberPoint, error)
	FindOneMemberPoint(memberId uint) (model.MemberPoint, error)
	FindAllMembersPoint(q map[string]interface{}, page, limit int) ([]model.MemberPoint, error)
	IncreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (model.MemberPoint, error)
	DecreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (model.MemberPoint, error)

	// Member point log
	CreatePointLog(data model.MemberPointLog, tx *gorm.DB) (model.MemberPointLog, error)
	UpdatePointLog(id uint, data model.MemberPointLog, tx *gorm.DB) (model.MemberPointLog, error)
}

type memberPointRepo struct {
	db *gorm.DB
}

func NewMemberPointRepository(db *gorm.DB) MemberPointRepo {
	return &memberPointRepo{db}
}

func (r *memberPointRepo) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *memberPointRepo) CreateMemberPoint(data model.MemberPoint, tx *gorm.DB) (model.MemberPoint, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return model.MemberPoint{}, err
	}

	return model.MemberPoint{}, nil
}

func (r *memberPointRepo) FindOneMemberPoint(memberId uint) (model.MemberPoint, error) {
	var data model.MemberPoint

	if err := r.db.Where("member_id = ?", memberId).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) FindAllMembersPoint(q map[string]interface{}, page, limit int) ([]model.MemberPoint, error) {
	var data []model.MemberPoint

	pagination := pagination.Pagination{Page: page, Limit: limit}

	if err := r.db.Scopes(pagination.Apply).Where(q).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) IncreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (model.MemberPoint, error) {
	var data model.MemberPoint

	db := r.getDB(tx)

	if err := db.Model(&data).Where("member_id = ?", memberId).Updates(map[string]interface{}{
		"total_points": gorm.Expr("total_points + ?", point),
	}).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) DecreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (model.MemberPoint, error) {
	var data model.MemberPoint

	db := r.getDB(tx)

	if err := db.Model(&data).Where("member_id = ?", memberId).Updates(map[string]interface{}{
		"total_points": gorm.Expr("total_points - ?", point),
	}).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) CreatePointLog(data model.MemberPointLog, tx *gorm.DB) (model.MemberPointLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return model.MemberPointLog{}, err
	}
	return data, nil
}

func (r *memberPointRepo) UpdatePointLog(id uint, data model.MemberPointLog, tx *gorm.DB) (model.MemberPointLog, error) {
	db := r.getDB(tx)
	var log model.MemberPointLog

	if err := db.Model(&log).Updates(data).Error; err != nil {
		return log, err
	}

	return log, nil
}
