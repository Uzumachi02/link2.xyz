package server

import (
	"github.com/Uzumachi02/link2.xyz/components/link"
	"github.com/Uzumachi02/link2.xyz/components/page"
	"github.com/Uzumachi02/link2.xyz/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Start ...
func Start(config *Config) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = utils.NewTemplate("views/*.html")

	page.RegisterRoutes(e)
	link.RegisterRoutes(e)

	return e.Start(config.BindAddr)
}
