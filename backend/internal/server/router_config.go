package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterConfig struct {
	Router *gin.Engine
	DB     *gorm.DB
}
