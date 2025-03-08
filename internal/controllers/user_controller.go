package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/models"
)

// CRUD Operations for user

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, user)
}
