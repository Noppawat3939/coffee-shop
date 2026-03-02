package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repository"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialEmployeeRoues(r *gin.RouterGroup) {
	db := cfg.DB
	repo := repository.NewEmployeeRepository(db)
	auditRepo := repository.NewAuditLogRepository(db)
	auditSvc := services.NewAuditLogService(auditRepo)
	service := services.NewEmployeeService(repo, auditSvc)
	controller := controllers.NewEmployeeController(repo, service, cfg.DB)

	employee := r.Group("/Employees", middleware.AuthGuard())
	{
		employee.GET("", controller.FindAll)
		employee.POST("register", controller.RegisterEmployee)
		employee.GET("/:id", controller.FindOne)
		employee.PATCH("/:id", controller.UpdateEmployee)
	}

}
