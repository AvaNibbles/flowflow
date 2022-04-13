package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avanibbles/flowflow/internal/routes/api/v1/apimodels"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type HttpHandleableError interface {
	error
	Handle(c echo.Context)
}

type StatusError struct {
	code    int
	message string
	inner   error
}

func MakeStatusErr(code int, message string, inner error) error {
	return StatusError{
		code:    code,
		message: message,
		inner:   inner,
	}
}

func (se StatusError) Error() string {
	return fmt.Sprintf("http err (%d) - %s: %s", se.code, se.message, se.inner)
}

func (se StatusError) Handle(c echo.Context) {
	body := apimodels.HttpError{Message: se.message, Code: se.code}
	if err := c.JSON(se.code, body); err != nil {
		c.Logger().Error(err)
	}
}

func (se StatusError) Code() int {
	return se.code
}

func MakeEchoErrorHandler(logger *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var hhe HttpHandleableError
		var echoErr *echo.HTTPError
		switch {
		case errors.As(err, &hhe):
			hhe.Handle(c)
		case errors.As(err, &echoErr):
			logger.Error("echo http error", zap.Error(echoErr.Internal), zap.Any("message", echoErr.Message))
			if err := c.NoContent(echoErr.Code); err != nil {
				logger.Error("error responding with an error", zap.Error(err))
			}
		default:
			logger.Error("unhandled http error", zap.Error(err))
			body := apimodels.HttpError{Message: err.Error(), Code: 500}
			if err = c.JSON(http.StatusInternalServerError, body); err != nil {
				logger.Error("error responding with an error", zap.Error(err))
			}
		}
	}
}
