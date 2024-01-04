package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ammarfar/mezink-golang-assignment/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo *echo.Echo
	db   *gorm.DB
}

// NewServer New Server constructor
func NewServer(db *gorm.DB) *Server {
	return &Server{
		echo: echo.New(),
		db:   db,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           config.Env.AppPort,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.echo.Logger.Infof("Server is listening on PORT: %s", config.Env.AppPort)
		if err := s.echo.StartServer(server); err != nil {
			s.echo.Logger.Infof("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.echo.Logger.Fatal("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
