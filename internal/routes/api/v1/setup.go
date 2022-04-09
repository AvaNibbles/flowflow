package v1

import (
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
)

func Setup(deps services.DependencyFactory, router *echo.Group) error {
	router.GET("/version", versionHandler())

	if err := setupHackRoutes(deps, router.Group("/hack")); err != nil {
		return err
	}

	return nil
}
