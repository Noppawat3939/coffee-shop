package server

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialEmployeeRoues(r *gin.RouterGroup) {
	db := cfg.DB
	repo := repository.NewEmployeeRepository(db)
	auditRepo := repository.NewAuditLogRepository(db)
	auditSvc := service.NewAuditLogService(auditRepo)
	service := service.NewEmployeeService(repo, auditSvc)
	handler := handler.NewEmployeeHandler(repo, service, cfg.DB)

	employee := r.Group("/Employees", middleware.AuthGuard())
	{
		employee.GET("", handler.FindAll)
		employee.POST("register", handler.RegisterEmployee)
		employee.GET("/:id", handler.FindOne)
		employee.PATCH("/:id", handler.UpdateEmployee)
	}

}
