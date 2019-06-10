package repository

import (
	"chapi-backend/order-service/model"
	"context"
)

type OrderRepository interface {
	UpdateStateOrder(context context.Context, order model.Order) error
	AddToCard(context context.Context, userId string, card model.Card) error
	CountShoppingCard(context context.Context, userId string) (model.OrderCount, error)
	ShoppingCard(context context.Context, userId string, orderId string) (model.Order, error)
}
