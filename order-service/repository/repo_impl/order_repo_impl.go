package repo_impl

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/order-service/model"
	"chapi-backend/order-service/repository"
	"context"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepository {
	return &OrderRepoImpl{
		sql: sql,
	}
}

func (o *OrderRepoImpl) CreateOrder(context context.Context, order model.Order) error {
	return nil;
}

func (o *OrderRepoImpl) UpdateStateOrder(context context.Context, stateOrder string) error {
	return nil
}

func (o *OrderRepoImpl) AddToCard(context context.Context, card model.Card) error {
	return nil
}

func (o *OrderRepoImpl) OrderInfo() model.Order {
	return model.Order{}
}
