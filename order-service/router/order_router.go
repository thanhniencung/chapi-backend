package router

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/chapi-internal/middleware"
	"chapi-backend/order-service/handler"
	repo "chapi-backend/order-service/repository/repo_impl"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, sql *db.Sql) {
	handler := handler.OrderHandler{
		OrderRepo: repo.NewOrderRepo(sql),
	}

	c := e.Group("/order")
	c.Use(middleware.JWTMiddleware())

	e.POST("/add", handler.AddToCard)
	e.POST("/confirm", handler.ConfirmOrder)
	e.POST("/info", handler.OrderCountItem)
	e.POST("/details", handler.OrderDetails)
}
