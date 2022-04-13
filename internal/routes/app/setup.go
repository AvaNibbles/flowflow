package app

import (
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
)

func Setup(deps services.DependencyFactory, router *echo.Echo) error {
	conf := deps.GetConfig()

	router.Static("*", conf.Http.SiteData)

	return nil
}
