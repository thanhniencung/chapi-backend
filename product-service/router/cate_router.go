package router

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/chapi-internal/middleware"
	"chapi-backend/product-service/handler"
	repo "chapi-backend/product-service/repository/repo_impl"
	"github.com/labstack/echo"
)

func CateRouter(e *echo.Echo, sql *db.Sql) {
	handler := handler.CateHandler{
		CateRepo: repo.NewCateRepo(sql),
	}

	c := e.Group("/cate/")
	c.Use(middleware.JWTMiddleware())

	c.POST("/cate/add", handler.Add)
	c.DELETE("/cate/delete", handler.Delete)
	c.PUT("/cate/update", handler.Update)
	c.GET("/cate/detail/:cate_id", handler.Details)
}