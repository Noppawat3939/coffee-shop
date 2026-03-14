package handler

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberHandler struct {
	memberSvc service.MemberService
	pointSvc  service.MemberPointService
	db        *gorm.DB
}

func NewMemberHandler(memberSvc service.MemberService, pointSvc service.MemberPointService, db *gorm.DB) *memberHandler {
	return &memberHandler{memberSvc, pointSvc, db}
}

func (h *memberHandler) GetMember(c *gin.Context) {
	var req dto.GetMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	member, err := h.memberSvc.FindMember(req.PhoneNumber)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, member)
}

func (h *memberHandler) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	member, err := h.memberSvc.CreateMember(req)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	data := model.MemberPoint{MemberID: member.ID, Member: member}
	_, err = h.pointSvc.CreateMemberPoint(data, nil)

	if err != nil {
		response.ErrorConflict(c)
		return
	}

	response.Success(c, member)
}

func (h *memberHandler) GetMembers(c *gin.Context) {
	p := pagination.NewFromQuery(c)

	filter := model.MemberFilter{
		PhoneNumber: c.Query("phone_number"),
		FullName:    c.Query("full_name"),
	}

	members, err := h.memberSvc.FindAllMembers(filter, p.Page, p.Limit)

	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, members)
}
