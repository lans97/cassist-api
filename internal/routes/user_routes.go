package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/controllers"
)

func UserAdminRoutes(g *echo.Group) {
    g.POST("", controllers.CreateUser)
    g.GET("/:id", controllers.GetUserById)
    g.GET("", controllers.GetUsers)
    g.PATCH("/:id", controllers.UpdateUser)
    g.DELETE("/:id", controllers.SoftDeleteUser)
    g.DELETE("/:id/hard", controllers.HardDeleteUser)
}
