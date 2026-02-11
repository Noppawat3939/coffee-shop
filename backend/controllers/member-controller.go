package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/services"
	"backend/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberController struct {
	memberSvc services.MemberService
	pointSvc  services.MemberPointService
	db        *gorm.DB
}

func NewMemberController(memberSvc services.MemberService, pointSvc services.MemberPointService, db *gorm.DB) *memberController {
	return &memberController{memberSvc, pointSvc, db}
}

func (mc *memberController) GetMember(c *gin.Context) {
	var req dto.GetMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
	}

	member, err := mc.memberSvc.FindMember(req.PhoneNumber)

	if err != nil {
		util.ErrorNotFound(c)
	}

	util.Success(c, member)
}

func (mc *memberController) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
	}

	member, err := mc.memberSvc.CreateMember(req)

	if err != nil {
		util.ErrorConflict(c)
	}

	data := models.MemberPoint{MemberID: member.ID, Member: member}
	_, err = mc.pointSvc.CreateMemberPoint(data, nil)

	if err != nil {
		util.ErrorConflict(c)
	}

	util.Success(c, member)
}
