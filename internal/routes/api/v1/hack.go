package v1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/avanibbles/flowflow/internal/services/storage"
	"go.uber.org/zap"

	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

func setupHackRoutes(deps services.DependencyFactory, router *echo.Group) error {
	router.POST("/s3/objects", uploadObject(deps))
	router.GET("/s3/objects/:key", downloadObject(deps))
	return nil
}

func uploadObject(deps services.DependencyFactory) echo.HandlerFunc {
	st := deps.GetStorage()
	return func(c echo.Context) error {
		objName := shortid.MustGenerate()

		req := storage.PutRequest{Key: objName, Body: c.Request().Body}
		resp, err := st.Put(req)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func downloadObject(deps services.DependencyFactory) echo.HandlerFunc {
	st := deps.GetStorage()
	logger := deps.GetLogger("downloadObject")
	return func(c echo.Context) error {
		key := c.Param("key")

		req := storage.GetRequest{Key: key}
		resp, err := st.Get(req)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		httpResp := c.Response()
		httpResp.Status = http.StatusOK

		respHeaders := httpResp.Writer.Header()
		respHeaders.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "my-file.csv"))
		respHeaders.Set("Cache-Control", "no-store")

		defer httpResp.Flush()
		bytesWritten, err := io.Copy(httpResp.Writer, resp.Body)
		if err != nil {
			return err
		}

		logger.Info("download object complete", zap.Int64("bytes", bytesWritten))
		return nil
	}
}
