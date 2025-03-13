package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

// CRUD Operations for user

// Create
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

    res := database.DB.Create(&user)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

	return c.JSON(http.StatusOK, user)
}

// Read
// By ID
func GetUserById(c echo.Context) error {
    user := models.User{}
    id := c.Param("id")

    res := database.DB.Where("id = ?", id).First(&user)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, user)
}

// By Username
func GetUserByUsername(c echo.Context) error {
    user := models.User{}
    email := c.Param("username")

    res := database.DB.Where("username = ?", email).First(&user)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, &user)
}

// With Pagination
func GetUsers(c echo.Context) error {
    users := []models.User{}

    res := database.DB.Scopes(database.Pagination(c)).Find(&users)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, &users)

}

// Update
func UpdateUser(c echo.Context) error {
    user_db := models.User{}
    user_req := models.User{}
    id := c.Param("id")
    c.Bind(&user_req)

    res := database.DB.Where("id = ?", id).First(&user_db)
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, res.Error)
        }
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    res = database.DB.Model(&user_db).Updates(user_req)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, user_req)
}

// Delete
// Non permanent
func SoftDeleteUser(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Delete(&models.User{}, id)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, id)
}

// Permanent
func HardDeleteUser(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Unscoped().Delete(&models.User{}, id)
    if res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, id)
}
