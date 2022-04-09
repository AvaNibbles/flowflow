package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(router *echo.Group) error {
	router.GET("/live", livenessHandler())
	router.GET("/ready", readinessHandler())
	return nil
}

func livenessHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

func readinessHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
