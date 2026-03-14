package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberPointHandler struct {
	svc       service.MemberPointService
	memberSvc service.MemberService
	db        *gorm.DB
}

func NewMemberPointHandler(svc service.MemberPointService, memberSvc service.MemberService, db *gorm.DB) *memberPointHandler {
	return &memberPointHandler{svc, memberSvc, db}
}

func (h *memberPointHandler) CreateMemberPoint(c *gin.Context) {
	phone_number := c.Param("phone_number")

	member, err := h.memberSvc.FindMember(phone_number)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	ok, _ := h.svc.CreateMemberPoint(model.MemberPoint{MemberID: member.ID, TotalPoints: 0}, nil)

	if !ok {
		response.Error(c, http.StatusConflict, "failed register a new member point")
		return
	}

	response.Success(c)
}
