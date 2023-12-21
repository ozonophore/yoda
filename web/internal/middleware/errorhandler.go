package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
)

func ErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	var e *echo.HTTPError
	if errors.As(err, &e) {
		c.JSON(e.Code, e.Message)
		return
	}
	errorData := &api.ErrorData{
		Success:     false,
		Description: err.Error(),
	}
	c.JSON(500, errorData)
}
