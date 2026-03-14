package handler

import (
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type auditLoghandler struct {
	repo repository.AuditLogRepository
	svc  service.AuditLogService
}

func NewAuditLogHandler(repo repository.AuditLogRepository, svc service.AuditLogService) *auditLoghandler {
	return &auditLoghandler{repo, svc}
}

func (h *auditLoghandler) FindAll(c *gin.Context) {
	var req dto.GetAuditLogRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		response.ErrorBodyInvalid(c)
		return
	}

	pagination := pagination.NewFromQuery(c)

	logs, err := h.svc.FindAll(req, pagination)
	if err != nil {
		response.ErrorNotFound(c)
		return
	}

	response.Success(c, logs)
}
