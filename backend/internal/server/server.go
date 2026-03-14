package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) *Server {
	r := gin.Default()

	r.RedirectTrailingSlash = true
	r.Use(gin.Logger(), gin.Recovery(), CORSMiddleware())

	s := &Server{
		db:     db,
		router: r,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}
