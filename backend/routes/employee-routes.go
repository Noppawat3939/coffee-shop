package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
)

func (cfg *RouterConfig) IntialEmployeeRoues(r *gin.RouterGroup) {
	repo := repository.NewEmployeeRepository(cfg.DB)
	controller := controllers.NewEmployeeController(repo, cfg.DB)

	employee := r.Group("/Employees")
	{
		employee.POST("register", controller.RegisterEmployee)
		employee.GET("", controller.FindAll)
		employee.GET("/:id", controller.FindOne)
		employee.PATCH("/:id", controller.UpdateEmployee)
	}

}
