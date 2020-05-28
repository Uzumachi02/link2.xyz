package link

import "github.com/labstack/echo/v4"

// RegisterRoutes ...
func RegisterRoutes(e *echo.Echo) {
	e.GET("/api/link", findLink)
	e.POST("/api/link", createLink)
}
