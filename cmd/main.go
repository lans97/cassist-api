package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lans97/cassist-api/internal/database"
)

func main() {
    database.InitDB()
    defer database.DB.Close()

	e := echo.New()

    e.Use(middleware.Logger())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, world!")
    })

    e.Logger.Fatal(e.Start(":42069"))
}
