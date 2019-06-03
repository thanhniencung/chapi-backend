package handler

import (
	"chapi-backend/chapi-internal/helper"
	"chapi-backend/product-service/repository"
	"github.com/labstack/echo"
	"net/http"
)

type CateHandler struct {
	CateRepo repository.CateRepository
}

func (m *CateHandler) Add(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}

func (m *CateHandler) Delete(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}

func (m *CateHandler) Update(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}

func (m *CateHandler) Details(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}

