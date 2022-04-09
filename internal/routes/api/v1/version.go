package v1

import (
	"net/http"

	"github.com/avanibbles/flowflow/internal"
	"github.com/labstack/echo/v4"
)

type VersionResponse struct {
	Version        string `json:"version"`
	CommitHash     string `json:"commit_hash"`
	BuildTimestamp string `json:"build_timestamp"`
}

// GetVersion godoc
// @Summary      Get flowflow version
// @Description  Get flowflow version
// @Tags         version
// @Produce      json
// @Success      200  {object}  VersionResponse
// @Router       /api/v1/version [get]
func versionHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := &VersionResponse{
			Version:        internal.Version,
			CommitHash:     internal.CommitHash,
			BuildTimestamp: internal.BuildTimestamp,
		}
		return c.JSON(http.StatusOK, resp)
	}
}
