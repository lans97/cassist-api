package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/middlewares"
	"github.com/lans97/cassist-api/internal/routes"
)

func main() {
    database.InitDB()

	e := echo.New()

    e.Use(middleware.Logger())

    e.HTTPErrorHandler = middlewares.CustomErrorHandler

    routes.UserRoutes(e.Group("/users"))

    e.Logger.Fatal(e.Start(":42069"))
}
