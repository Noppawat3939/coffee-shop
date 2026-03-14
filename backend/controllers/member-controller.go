package controllers

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberController struct {
	memberSvc service.MemberService
	pointSvc  service.MemberPointService
	db        *gorm.DB
}

func NewMemberController(memberSvc service.MemberService, pointSvc service.MemberPointService, db *gorm.DB) *memberController {
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

	data := model.MemberPoint{MemberID: member.ID, Member: member}
	_, err = mc.pointSvc.CreateMemberPoint(data, nil)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, member)
}

func (mc *memberController) GetMembers(c *gin.Context) {
	p := pagination.NewFromQuery(c)

	filter := model.MemberFilter{
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
