package link

import "github.com/labstack/echo/v4"

// ILinkRepository ...
type ILinkRepository interface {
	Create(*LinkModel) error
	Find(int) (*LinkModel, error)
}

// RegisterRoutes ...
func RegisterRoutes(e *echo.Echo) {
	e.GET("/api/link", findLink)
	e.POST("/api/link", createLink)
}
