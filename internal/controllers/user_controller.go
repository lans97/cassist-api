package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

var Response = models.JSONResponse

// CRUD Operations for user

// Create
func CreateUser(c echo.Context) error {
	user := models.User{}
    err := c.Bind(&user)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Bad request: User struct")
    }

    res := database.DB.Create(&user)
    if res.Error != nil {
        if res.Error == gorm.ErrDuplicatedKey {
            return echo.NewHTTPError(http.StatusConflict, "Duplicate entry")
        }
            return echo.NewHTTPError(http.StatusInternalServerError, "User creation failed")
    }

    return Response(c, http.StatusOK, "User created", user)
}

// Read
// By ID
func GetUserById(c echo.Context) error {
    user := models.User{}
    id := c.Param("id")

    res := database.DB.Where("id = ?", id).First(&user)
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "User retreived", user)
}

// Filters and Pagination
func GetUsers(c echo.Context) error {
    users := []models.User{}

    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    // Filters
    uuid := c.QueryParam("uuid")
    email := c.QueryParam("email")
    displayName := c.QueryParam("display_name")

    query := database.DB.Model(&models.User{})

    // Filter check

    if uuid != "" {
        query = query.Where("uuid = ?", uuid)
    }

    if email != "" {
        query = query.Where("email = ?", email)
    }

    if displayName != "" {
        query = query.Where("display_name = ?", displayName)
    }

    limit_n := 0
    page_n := 0

    var err error

    // Determine limit and page
    if limit != "" {
        limit_n, err = strconv.Atoi(limit)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Bad request: limit NaN")
        }
    }

    if page != "" {
        page_n, err = strconv.Atoi(page)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Bad request: page NaN")
        }
    }

    offset := page_n * limit_n // If any are 0 just show everything

    if offset != 0 {
        query = query.Limit(limit_n).Offset(limit_n*page_n)
    }

    res := query.Find(&users)

    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "Users retreived", users)
}

// Update
func UpdateUser(c echo.Context) error {
    id := c.Param("id")

    var updates map[string]any

    if err := c.Bind(&updates); err != nil {
        return err
    }

    if email, exists := updates["email"]; exists && email != "" {
        updates["email_verified"] = false
    }

    res := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
    }

    var user models.User
    database.DB.First(&user, id)

    return Response(c, http.StatusOK, "User updated", user)
}

// Delete
// Non permanent
func SoftDeleteUser(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Delete(&models.User{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return echo.NewHTTPError(http.StatusOK, id)
}

// Permanent
func HardDeleteUser(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Unscoped().Delete(&models.User{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return echo.NewHTTPError(http.StatusOK, id)
}
