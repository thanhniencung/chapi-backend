package main

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/order-service/router"
	"github.com/labstack/echo"
		"github.com/labstack/echo/middleware"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "ryan",
		Password: "postgres",
		DbName:   "order-service",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	router.Router(e, sql)
	e.Logger.Fatal(e.Start(":3003"))
}
