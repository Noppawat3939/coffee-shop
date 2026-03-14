package server

import (
	"backend/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialEmployeeRoues(r *gin.RouterGroup) {
	db := cfg.DB
	repo := repository.NewEmployeeRepository(db)
	auditRepo := repository.NewAuditLogRepository(db)
	auditSvc := service.NewAuditLogService(auditRepo)
	service := service.NewEmployeeService(repo, auditSvc)
	controller := controllers.NewEmployeeController(repo, service, cfg.DB)

	employee := r.Group("/Employees", middleware.AuthGuard())
	{
		employee.GET("", controller.FindAll)
		employee.POST("register", controller.RegisterEmployee)
		employee.GET("/:id", controller.FindOne)
		employee.PATCH("/:id", controller.UpdateEmployee)
	}

}
