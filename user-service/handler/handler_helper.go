package handler

import (
	internalModel "chapi-backend/chapi-internal/model"
	"github.com/labstack/echo"
	"net/http"
)

func ResponseErr(c echo.Context, code int, errMsg ...string) error {
	var msg = errMsg[0]
	if len(errMsg) == 0 {
		msg = http.StatusText(code)
	}
	return c.JSON(code, internalModel.Response{
		StatusCode:  code,
		Message: msg,
	})
}

func ResponseData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, internalModel.Response{
		StatusCode:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: data,
	})
}
