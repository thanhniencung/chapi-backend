package router

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/chapi-internal/middleware"
	"chapi-backend/user-service/handler"
	repo "chapi-backend/user-service/repository/repo_impl"
	"github.com/labstack/echo"
	"net"
	"net/http"
)

func Router(e *echo.Echo, sql *db.Sql) {
	handler := handler.UserHandler{
		UserRepo: repo.NewUserRepo(sql),
	}

	e.GET("/", func(c echo.Context) error {

		conn, _ := net.Dial("udp", "8.8.8.8:80")
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)

		return c.JSON(http.StatusOK, echo.Map{
			"ip":  localAddr.IP,
		})
	})

	e.POST("/sign-in", handler.SignIn)
	e.POST("/sign-up", handler.SignUp)

	e.GET("/profile", handler.Profile, middleware.JWTMiddleware())
}
