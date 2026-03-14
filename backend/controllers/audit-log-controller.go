package controllers

import (
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/pkg/pagination"
	"backend/pkg/response"
	"backend/services"

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
		response.ErrorBodyInvalid(ctx)
		return
	}

	pagination := pagination.NewFromQuery(ctx)

	logs, err := c.svc.FindAll(req, pagination)
	if err != nil {
		response.ErrorNotFound(ctx)
		return
	}

	response.Success(ctx, logs)
}
