package controllers

import (
	"backend/dto"
	"backend/repository"
	"backend/services"
	"backend/util"

	"github.com/gin-gonic/gin"
)

type auditLogController struct {
	repo repository.AuditLogRepository
	svc  services.AuditLogService
}

func NewAuditLogController(repo repository.AuditLogRepository, svc services.AuditLogService) *auditLogController {
	return &auditLogController{repo, svc}
}

func (c *auditLogController) FindAll(ctx *gin.Context) {
	var req dto.GetAuditLogRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		util.ErrorBodyInvalid(ctx)
		return
	}

	pagination := util.NewPaginationFromQuery(ctx)

	logs, err := c.svc.FindAll(req, pagination)
	if err != nil {
		util.ErrorNotFound(ctx)
		return
	}

	util.Success(ctx, logs)
}
