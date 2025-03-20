package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

// CRUD Operations for role

// Create
func CreateRole(c echo.Context) error {
	role := models.Role{}
    err := c.Bind(&role)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Bad request: Role struct")
    }

    res := database.DB.Create(&role)
    if res.Error != nil {
        if res.Error == gorm.ErrDuplicatedKey {
            return echo.NewHTTPError(http.StatusConflict, "Duplicate entry")
        }
            return echo.NewHTTPError(http.StatusInternalServerError, "Role creation failed")
    }

    return Response(c, http.StatusOK, "Role created", role)
}

// Read
// By ID
func GetRoleById(c echo.Context) error {
    role := models.Role{}
    id := c.Param("id")

    res := database.DB.Where("id = ?", id).First(&role)
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "Role not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "Role retreived", role)
}

// Filters and Pagination
func GetRoles(c echo.Context) error {
    roles := []models.Role{}

    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    // Filters
    name := c.QueryParam("name")

    query := database.DB.Model(&models.Role{})

    // Filter check

    if name != "" {
        query = query.Where("name = ?", name)
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

    res := query.Find(&roles)

    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "Roles retreived", roles)
}

// Update
func UpdateRole(c echo.Context) error {
    id := c.Param("id")

    var updates map[string]any

    if err := c.Bind(&updates); err != nil {
        return err
    }

    if email, exists := updates["email"]; exists && email != "" {
        updates["email_verified"] = false
    }

    res := database.DB.Model(&models.Role{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update role")
    }

    var role models.Role
    database.DB.First(&role, id)

    return Response(c, http.StatusOK, "Role updated", role)
}

// Delete
// Non permanent
func SoftDeleteRole(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Delete(&models.Role{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}

// Permanent
func HardDeleteRole(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Unscoped().Delete(&models.Role{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}
