package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/controllers"
)

func MoneyBucketRoutes(g *echo.Group) {
	g.POST("", controllers.CreateMoneyBucket)
	g.GET("/:id", controllers.GetMoneyBucketById)
	g.GET("", controllers.GetMoneyBuckets)
	g.PATCH("/:id", controllers.UpdateMoneyBucket)
	g.DELETE("/:id", controllers.SoftDeleteMoneyBucket)
	g.DELETE("/:id/hard", controllers.HardDeleteMoneyBucket)
}
