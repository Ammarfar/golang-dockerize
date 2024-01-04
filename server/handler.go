package server

import (
	"net/http"

	gorm_repository "github.com/Ammarfar/mezink-golang-assignment/internal/database/gorm"
	"github.com/Ammarfar/mezink-golang-assignment/internal/logger"
	apiMiddleware "github.com/Ammarfar/mezink-golang-assignment/internal/middleware"
	userHttp "github.com/Ammarfar/mezink-golang-assignment/internal/modules/user/delivery/http"
	userUseCase "github.com/Ammarfar/mezink-golang-assignment/internal/modules/user/usecase"
	"github.com/Ammarfar/mezink-golang-assignment/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init db repositories
	userRepo := gorm_repository.NewUserRepositoryGorm(s.db)

	// logger
	logger := logger.NewEchoLogger(s.echo.Logger)

	// Init useCases
	userUC := userUseCase.NewUserUseCase(logger, userRepo)

	// Init handlers
	userHandler := userHttp.NewUserHandler(logger, userUC)

	mw := apiMiddleware.NewMiddlewareManager()

	// global mdlwr
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	userGroup := v1.Group("/users")

	userHttp.MapUserRoutes(userGroup, userHandler, mw)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.Response{
			Code: 0,
			Msg:  "success",
		})
	})

	health.GET("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, utils.Response{})
	})

	return nil
}
