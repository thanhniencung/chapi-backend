package handler

import (
	"chapi-backend/chapi-internal/encrypt"
	"chapi-backend/chapi-internal/helper"
	"chapi-backend/product-service/model"
	"chapi-backend/product-service/repository"
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepository
}

func (p *ProductHandler) Add(c echo.Context) error {
	req := model.Product{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _,err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	req.ProductId = encrypt.UUIDV1()
	product, err := p.ProductRepo.AddProduct(ctx, req)
	if err != nil {
		fmt.Println(err)
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, product)
}

func (p *ProductHandler) Delete(c echo.Context) error {
	defer c.Request().Body.Close()

	productId := c.Param("product_id")
	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)

	err := p.ProductRepo.DeletePRoduct(ctx, productId)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, nil)
}

func (p *ProductHandler) Update(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}

func (p *ProductHandler) Details(c echo.Context) error {
	return helper.ResponseErr(c, http.StatusBadRequest)
}