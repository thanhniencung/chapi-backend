package repo_impl

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/product-service/model"
	"chapi-backend/product-service/repository"
	"context"
)

type CateRepoImpl struct {
	sql *db.Sql
}

func NewCateRepo(sql *db.Sql) repository.CateRepository {
	return &CateRepoImpl{
		sql: sql,
	}
}

func (c *CateRepoImpl) AddCate(context context.Context, cate model.Cate) (model.Cate, error) {
	return model.Cate{}, nil
}

func (c *CateRepoImpl) UpdateCate(context context.Context, cate model.Cate) (model.Cate, error) {
	return model.Cate{}, nil
}

func (c *CateRepoImpl) DeleteCate(context context.Context, cateId string) (error) {
	return nil
}

func (c *CateRepoImpl) SelectCateById(context context.Context, userId string) (model.Cate, error) {
	return model.Cate{}, nil
}

