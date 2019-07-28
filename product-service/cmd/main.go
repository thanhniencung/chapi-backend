package main

import (
	"github.com/thanhniencung/chapi-internal/db"
	"product-service/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	sql := &db.Sql{
		Host:     "host.docker.internal", //host.docker.internal
		Port:     5432,
		UserName: "ryan",
		Password: "postgres",
		DbName:   "product-service",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	
	router.ProductRouter(e, sql)
	router.CateRouter(e, sql)

	e.Logger.Fatal(e.Start(":3001"))
}
