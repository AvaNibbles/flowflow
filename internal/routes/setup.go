package routes

import (
	"github.com/avanibbles/flowflow/internal/routes/api"
	"github.com/avanibbles/flowflow/internal/routes/health"
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
)

func Setup(deps services.DependencyFactory, router *echo.Echo) error {
	if err := health.Setup(router.Group("/health")); err != nil {
		return err
	}

	if err := api.Setup(deps, router.Group("/api")); err != nil {
		return err
	}

	return nil
}
