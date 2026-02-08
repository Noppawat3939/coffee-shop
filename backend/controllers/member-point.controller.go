package controllers

import (
	"backend/repository"
	"backend/services"

	"gorm.io/gorm"
)

type memberPointController struct {
	repo repository.MemberPointRepo
	svc  services.MemberPointService
	db   *gorm.DB
}

func NewMemberPointController(repo repository.MemberPointRepo, svc services.MemberPointService, db *gorm.DB) *memberPointController {
	return &memberPointController{repo, svc, db}
}
