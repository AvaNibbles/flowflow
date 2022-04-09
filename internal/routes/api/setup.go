package api

import (
	v1 "github.com/avanibbles/flowflow/internal/routes/api/v1"
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
)

func Setup(deps services.DependencyFactory, router *echo.Group) error {
	if err := v1.Setup(deps, router.Group("/v1")); err != nil {
		return err
	}

	return nil
}
