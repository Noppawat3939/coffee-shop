package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialEmployeeRoues(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewEmployeeRepository(db)
	controller := controllers.NewEmployeeController(repo, db)

	employee := r.Group("/Employees")
	{
		employee.POST("", controller.CreateEmployee)
		employee.GET("", controller.FindAll)
		employee.GET("/:id", controller.FindOne)
		employee.PATCH("/:id", controller.UpdateEmployee)
	}

}
