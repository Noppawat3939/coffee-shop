package server

import (
	"backend/internal/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) *Server {
	r := gin.New()

	r.RedirectTrailingSlash = true
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		CORS(),
		middleware.ReqID(),
		middleware.ReqLogger(),
	)

	s := &Server{
		db:     db,
		router: r,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Start(port string) error {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      s.router,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// run server (non-blocking)
	go func() {
		fmt.Println("Server running on port:", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error:", err)
		}
	}()

	// listen shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shutdown http server
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server force shutdown:", err)
		return err
	}

	// close DB
	if err := s.closeDB(); err != nil {
		log.Println("DB close error:", err)
	}

	fmt.Println("Server exited")
	return nil
}

func (s *Server) closeDB() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
