package router

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/chapi-internal/middleware"
	"github.com/labstack/echo"
	"chapi-backend/product-service/handler"
	repo "chapi-backend/product-service/repository/repo_impl"
)

func ProductRouter(e *echo.Echo, sql *db.Sql) {
	handler := handler.ProductHandler{
		ProductRepo: repo.NewProductRepo(sql),
	}

	p := e.Group("/product")
	p.Use(middleware.JWTMiddleware())

	p.POST("/add", handler.Add, )
	p.DELETE("/delete/:product_id", handler.Delete)
	p.PUT("/update", handler.Update)
	p.GET("/detail/:product_id", handler.Details)
}
