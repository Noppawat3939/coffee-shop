package server

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialAuditLogRoutes(r *gin.RouterGroup) {
	repo := repository.NewAuditLogRepository(cfg.DB)
	svc := service.NewAuditLogService(repo)
	handler := handler.NewAuditLogHandler(repo, svc)

	admin := r.Group("/Admin", middleware.AuthGuard())

	admin.GET("/audit-logs", handler.FindAll)
}
