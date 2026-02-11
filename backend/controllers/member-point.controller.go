package controllers

import (
	"backend/models"
	"backend/services"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberPointController struct {
	svc       services.MemberPointService
	memberSvc services.MemberService
	db        *gorm.DB
}

func NewMemberPointController(svc services.MemberPointService, memberSvc services.MemberService, db *gorm.DB) *memberPointController {
	return &memberPointController{svc, memberSvc, db}
}

func (mc *memberPointController) CreateMemberPoint(c *gin.Context) {
	phone_number := c.Param("phone_number")

	member, err := mc.memberSvc.FindMember(phone_number)
	if err != nil {
		util.ErrorNotFound(c)
	}

	ok, _ := mc.svc.CreateMemberPoint(models.MemberPoint{MemberID: member.ID, TotalPoints: 0}, nil)

	if !ok {
		util.Error(c, http.StatusConflict, "failed register a new member point")
	}

	util.Success(c)
}
