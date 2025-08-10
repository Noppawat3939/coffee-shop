package routes

import (
	"backend/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.NoRoute(
		func(c *gin.Context) { helpers.ErrorNotFound(c) })

	api := r.Group("/api")

	IntialMenuRoutes(api, db)
}

func SetupSwagger(r *gin.Engine, port string) {
	fmt.Println("✅ Swagger run on http://localhost:" + port + "/swagger/index.html")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
