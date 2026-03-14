package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialAuditLogRoutes(r *gin.RouterGroup) {
	repo := repository.NewAuditLogRepository(cfg.DB)
	svc := service.NewAuditLogService(repo)
	controller := controllers.NewAuditLogController(repo, svc)

	admin := r.Group("/Admin", middleware.AuthGuard())

	admin.GET("/audit-logs", controller.FindAll)
}
