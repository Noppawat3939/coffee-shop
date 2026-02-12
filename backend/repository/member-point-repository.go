package repository

import (
	"backend/models"
	"backend/util"

	"gorm.io/gorm"
)

type MemberPointRepo interface {
	// Member point
	CreateMemberPoint(data models.MemberPoint, tx *gorm.DB) (models.MemberPoint, error)
	FindOneMemberPoint(memberId uint) (models.MemberPoint, error)
	FindAllMembersPoint(q map[string]interface{}, page, limit int) ([]models.MemberPoint, error)
	IncreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (models.MemberPoint, error)
	DecreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (models.MemberPoint, error)

	// Member point log
	CreatePointLog(data models.MemberPointLog, tx *gorm.DB) (models.MemberPointLog, error)
	UpdatePointLog(id uint, data models.MemberPointLog, tx *gorm.DB) (models.MemberPointLog, error)
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

func (r *memberPointRepo) CreateMemberPoint(data models.MemberPoint, tx *gorm.DB) (models.MemberPoint, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return models.MemberPoint{}, err
	}

	return models.MemberPoint{}, nil
}

func (r *memberPointRepo) FindOneMemberPoint(memberId uint) (models.MemberPoint, error) {
	var data models.MemberPoint

	if err := r.db.Where("member_id = ?", memberId).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) FindAllMembersPoint(q map[string]interface{}, page, limit int) ([]models.MemberPoint, error) {
	var data []models.MemberPoint

	pagination := util.Pagination{Page: page, Limit: limit}

	if err := r.db.Scopes(pagination.GetPaginationResult).Where(q).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) IncreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (models.MemberPoint, error) {
	var data models.MemberPoint

	db := r.getDB(tx)

	if err := db.Model(&data).Where("member_id = ?", memberId).Updates(map[string]interface{}{
		"total_points": gorm.Expr("total_points + ?", point),
	}).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) DecreaseMemberPoint(memberId uint, point int, tx *gorm.DB) (models.MemberPoint, error) {
	var data models.MemberPoint

	db := r.getDB(tx)

	if err := db.Model(&data).Where("member_id = ?", memberId).Updates(map[string]interface{}{
		"total_points": gorm.Expr("total_points - ?", point),
	}).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *memberPointRepo) CreatePointLog(data models.MemberPointLog, tx *gorm.DB) (models.MemberPointLog, error) {
	db := r.getDB(tx)
	if err := db.Create(&data).Error; err != nil {
		return models.MemberPointLog{}, err
	}
	return data, nil
}

func (r *memberPointRepo) UpdatePointLog(id uint, data models.MemberPointLog, tx *gorm.DB) (models.MemberPointLog, error) {
	db := r.getDB(tx)
	var log models.MemberPointLog

	if err := db.Model(&log).Updates(data).Error; err != nil {
		return log, err
	}

	return log, nil
}
