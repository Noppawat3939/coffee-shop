package controllers

import (
	"backend/internal/dto"
	"backend/models"
	"backend/pkg/pagination"
	"backend/pkg/response"
	"backend/services"

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
		response.ErrorBodyInvalid(c)
		return
	}

	member, err := mc.memberSvc.FindMember(req.PhoneNumber)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, member)
}

func (mc *memberController) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	member, err := mc.memberSvc.CreateMember(req)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	data := models.MemberPoint{MemberID: member.ID, Member: member}
	_, err = mc.pointSvc.CreateMemberPoint(data, nil)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, member)
}

func (mc *memberController) GetMembers(c *gin.Context) {
	p := pagination.NewFromQuery(c)

	filter := models.MemberFilter{
		PhoneNumber: c.Query("phone_number"),
		FullName:    c.Query("full_name"),
	}

	members, err := mc.memberSvc.FindAllMembers(filter, p.Page, p.Limit)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, members)
}
