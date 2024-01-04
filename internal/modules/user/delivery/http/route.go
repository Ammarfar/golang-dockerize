package user

import (
	"github.com/Ammarfar/mezink-golang-assignment/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapUserRoutes(userGroup *echo.Group, h UserHandlers, mw *middleware.MiddlewareManager) {
	userGroup.GET("/marks", h.GetAll())
}
