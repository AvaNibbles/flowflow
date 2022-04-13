package v1

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/avanibbles/flowflow/internal/util"

	"github.com/avanibbles/flowflow/internal/services/storage"
	"go.uber.org/zap"

	"github.com/avanibbles/flowflow/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

func setupHackRoutes(deps services.DependencyFactory, router *echo.Group) error {
	router.POST("/s3/objects", uploadObject(deps))
	router.GET("/s3/objects/:key", downloadObject(deps))
	router.GET("/err/:code", makeStatusError())
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

// MakeError godoc
// @Summary      Make an error
// @Description  Make an error
// @Tags         hack
// @Produce      json
// @Param        code  path      int  true  "http code"
// @Success      200   {object}  apimodels.HttpError
// @Failure      400   {object}  apimodels.HttpError
// @Failure      401   {object}  apimodels.HttpError
// @Failure      403   {object}  apimodels.HttpError
// @Failure      404   {object}  apimodels.HttpError
// @Failure      409   {object}  apimodels.HttpError
// @Failure      500   {object}  apimodels.HttpError
// @Router       /api/v1/hack/err/{code} [get]
func makeStatusError() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")
		codeNum, err := strconv.Atoi(code)
		if err != nil {
			return err
		}

		return util.MakeStatusErr(codeNum, "test", nil)
	}
}
