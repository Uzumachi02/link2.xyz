package page

import "github.com/labstack/echo/v4"

// RegisterRoutes ...
func RegisterRoutes(e *echo.Echo) {
	e.GET("/ping", ping)
	e.GET("/", indexPage)
}
