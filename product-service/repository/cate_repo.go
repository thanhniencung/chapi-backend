package repository

import (
	"chapi-backend/product-service/model"
	"context"
)

type CateRepository interface {
	AddCate(context context.Context, cate model.Cate) (model.Cate, error)
	UpdateCate(context context.Context, cate model.Cate) (model.Cate, error)
	DeleteCate(context context.Context, cateId string) (error)
	SelectCateById(context context.Context, userId string) (model.Cate, error)
}
