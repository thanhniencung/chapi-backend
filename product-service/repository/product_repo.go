package repository

import (
	"chapi-backend/product-service/model"
	"context"
)

type ProductRepository interface {
	AddProduct(context context.Context, product model.Product) (model.Product, error)
	UpdateProduct(context context.Context, product model.Product) (model.Product, error)
	DeletePRoduct(context context.Context, productId string) (error)
	SelectProductById(context context.Context, userId string) (model.Product, error)
}
