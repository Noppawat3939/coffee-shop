package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
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
