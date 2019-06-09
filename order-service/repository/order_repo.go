package repository

import (
	"chapi-backend/order-service/model"
	"context"
)

type OrderRepository interface {
	CreateOrder(context context.Context, order model.Order) error
	UpdateStateOrder(context context.Context, stateOrder string) error
	AddToCard(context context.Context, card model.Card) error
	OrderInfo() model.Order
}
