package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/controllers"
)

func UserRoutes(g echo.Group) {
    g.POST("/", controllers.CreateUser)
    g.GET("/id/:id", controllers.GetUserById)
    g.GET("/username/:username", controllers.GetUserByUsername)
    g.GET("/", controllers.GetUsers)
    g.PUT("/id/:id", controllers.UpdateUser)
    g.DELETE("/soft/:id", controllers.SoftDeleteUser)
    g.DELETE("/hard/:id", controllers.HardDeleteUser)
}
