package routes

import (
	"backend/controllers"
	"backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntialOrderRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewOrderRepository(db)
	controller := controllers.NewOrderController(repo, db)

	order := r.Group("/orders")
	{
		order.POST("", controller.CreateOrder)
	}
}
