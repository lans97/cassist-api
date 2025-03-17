package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/controllers"
)

func CategoryAdminRoutes(g *echo.Group) {
    g.POST("", controllers.CreateCategory)
    g.GET("/:id", controllers.GetCategoryById)
    g.GET("", controllers.GetCategories)
    g.PATCH("/:id", controllers.UpdateCategory)
    g.DELETE("/:id", controllers.SoftDeleteCategory)
    g.DELETE("/:id/hard", controllers.HardDeleteCategory)
}
