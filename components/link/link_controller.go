package link

import (
	"net/http"

	"github.com/Uzumachi02/link2.xyz/utils"
	"github.com/labstack/echo/v4"
)

func createLink(c echo.Context) error {
	f := new(createLinkModel)

	if err := c.Bind(f); err != nil {
		return err
	}

	hash, err := saveLink(f)
	if err != nil {
		return err
	}

	response := &utils.APIResponse{
		Response: hash,
	}

	return c.JSON(http.StatusOK, response)
}

func findLink(c echo.Context) error {
	return c.JSON(http.StatusOK, "findLink")
}
