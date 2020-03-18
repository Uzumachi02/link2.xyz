package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", "")
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
