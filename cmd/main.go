package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/firebase"
	"github.com/lans97/cassist-api/internal/middlewares"
	"github.com/lans97/cassist-api/internal/routes"
	"golang.org/x/time/rate"
)

func main() {
    database.InitDB()
    firebase.InitFirebase()

	e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
    e.Use(middlewares.FirebaseAuth)

    e.HTTPErrorHandler = middlewares.CustomErrorHandler

    adminGroup := e.Group("/admin", middlewares.RequireRole("admin"))

    routes.UserAdminRoutes(adminGroup.Group("/users"))
    routes.MoneyBucketAdminRoutes(adminGroup.Group("/money_buckets"))
    routes.CategoryAdminRoutes(adminGroup.Group("/categories"))
    routes.TransactionAdminRoutes(adminGroup.Group("/transactions"))

    e.Logger.Fatal(e.Start(":42069"))
}
