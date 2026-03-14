package controllers

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberPointController struct {
	svc       service.MemberPointService
	memberSvc service.MemberService
	db        *gorm.DB
}

func NewMemberPointController(svc service.MemberPointService, memberSvc service.MemberService, db *gorm.DB) *memberPointController {
	return &memberPointController{svc, memberSvc, db}
}

func (mc *memberPointController) CreateMemberPoint(c *gin.Context) {
	phone_number := c.Param("phone_number")

	member, err := mc.memberSvc.FindMember(phone_number)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	ok, _ := mc.svc.CreateMemberPoint(model.MemberPoint{MemberID: member.ID, TotalPoints: 0}, nil)

	if !ok {
		response.Error(c, http.StatusConflict, "failed register a new member point")
		return
	}

	response.Success(c)
}
