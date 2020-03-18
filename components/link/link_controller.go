package link

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func createLink(c echo.Context) error {
	f := new(createLinkModel)

	if err := c.Bind(f); err != nil {
		return err
	}

	saveLink(f)

	return c.JSON(http.StatusOK, f)
}

func findLink(c echo.Context) error {
	return c.JSON(http.StatusOK, "findLink")
}
