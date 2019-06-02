package main

import (
	"chapi-backend/user-service/router"
	"chapi-backend/chapi-internal/db"
	"github.com/labstack/echo"
)

func main() {
	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "ryan",
		Password: "postgres",
		DbName: "user-service",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	router.Router(e, sql)
	e.Logger.Fatal(e.Start(":3000"))
}