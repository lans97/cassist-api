package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/controllers"
)

func TransactionRoutes(g *echo.Group) {
    g.POST("", controllers.CreateTransaction)
    g.GET("/:id", controllers.GetTransactionById)
    g.GET("", controllers.GetTransactions)
    g.PATCH("/:id", controllers.UpdateTransaction)
    g.DELETE("/:id", controllers.SoftDeleteTransaction)
    g.DELETE("/:id/hard", controllers.HardDeleteTransaction)
}
