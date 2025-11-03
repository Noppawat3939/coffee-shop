package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberController struct {
	repo repository.MemberRepo
	db   *gorm.DB
}

func NewMemberController(repo repository.MemberRepo, db *gorm.DB) *memberController {
	return &memberController{repo, db}
}

func (mc *memberController) GetMember(c *gin.Context) {
	var req dto.GetMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return

	}

	member, err := mc.repo.FindOne(req.PhoneNumber)

	if err != nil {
		util.ErrorNotFound(c)
		return
	}

	util.Success(c, member)
}

func (mc *memberController) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorBodyInvalid(c)
		return
	}

	member, err := mc.repo.Create(models.Member{
		PhoneNumber: req.PhoneNumber,
		FullName:    req.FullName,
		Provider:    "line", // default
	})

	if err != nil {
		util.Error(c, http.StatusConflict, "failed create member with line")
		return
	}

	util.Success(c, member)
}
