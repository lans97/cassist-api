package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

// CRUD Operations for money_bucket

// Create
func CreateMoneyBucket(c echo.Context) error {
	money_bucket := models.MoneyBucket{}
    err := c.Bind(&money_bucket)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Bad request: MoneyBucket struct")
    }

    res := database.DB.Create(&money_bucket)
    if res.Error != nil {
        if res.Error == gorm.ErrDuplicatedKey {
            return echo.NewHTTPError(http.StatusConflict, "Duplicate entry")
        }
            return echo.NewHTTPError(http.StatusInternalServerError, "MoneyBucket creation failed")
    }

    return Response(c, http.StatusOK, "MoneyBucket created", money_bucket)
}

// Read
// By ID
func GetMoneyBucketById(c echo.Context) error {
    money_bucket := models.MoneyBucket{}
    id := c.Param("id")

    res := database.DB.Where("id = ?", id).First(&money_bucket)
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "MoneyBucket not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "MoneyBucket retreived", money_bucket)
}

// Filters and Pagination
func GetMoneyBuckets(c echo.Context) error {
    money_buckets := []models.MoneyBucket{}

    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    // Filters
    userID := c.QueryParam("user_id")
    name := c.QueryParam("name")

    query := database.DB.Model(&models.MoneyBucket{})

    // Filter check

    if userID != "" {
        query = query.Where("user_id = ?", userID)
    }

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

    res := query.Find(&money_buckets)

    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "MoneyBuckets retreived", money_buckets)
}

// Update
func UpdateMoneyBucket(c echo.Context) error {
    id := c.Param("id")

    var updates map[string]any

    if err := c.Bind(&updates); err != nil {
        return err
    }

    if email, exists := updates["email"]; exists && email != "" {
        updates["email_verified"] = false
    }

    res := database.DB.Model(&models.MoneyBucket{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update money_bucket")
    }

    var money_bucket models.MoneyBucket
    database.DB.First(&money_bucket, id)

    return Response(c, http.StatusOK, "MoneyBucket updated", money_bucket)
}

// Delete
// Non permanent
func SoftDeleteMoneyBucket(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Delete(&models.MoneyBucket{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}

// Permanent
func HardDeleteMoneyBucket(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Unscoped().Delete(&models.MoneyBucket{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}
