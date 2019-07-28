package router

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/chapi-internal/middleware"
	"chapi-backend/user-service/handler"
	repo "chapi-backend/user-service/repository/repo_impl"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, sql *db.Sql) {

	handler := handler.UserHandler{
		UserRepo: repo.NewUserRepo(sql),
	}

	e.POST("/sign-in", handler.SignIn)
	e.POST("/sign-up", handler.SignUp)

	e.GET("/profile", handler.Profile, middleware.JWTMiddleware())
	e.GET("/list", handler.List, middleware.JWTMiddleware())
}
