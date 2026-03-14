package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/middleware"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) InitialAuditLogRoutes(r *gin.RouterGroup) {
	repo := repository.NewAuditLogRepository(cfg.DB)
	svc := services.NewAuditLogService(repo)
	controller := controllers.NewAuditLogController(repo, svc)

	admin := r.Group("/Admin", middleware.AuthGuard())

	admin.GET("/audit-logs", controller.FindAll)
}
