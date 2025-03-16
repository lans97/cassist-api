package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

// CRUD Operations for transaction

// Create
func CreateTransaction(c echo.Context) error {
	transaction := models.Transaction{}
    err := c.Bind(&transaction)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Bad request: Transaction struct")
    }

    res := database.DB.Create(&transaction)
    if res.Error != nil {
        if res.Error == gorm.ErrDuplicatedKey {
            return echo.NewHTTPError(http.StatusConflict, "Duplicate entry")
        }
            return echo.NewHTTPError(http.StatusInternalServerError, "Transaction creation failed")
    }

    return Response(c, http.StatusOK, "Transaction created", transaction)
}

// Read
// By ID
func GetTransactionById(c echo.Context) error {
    transaction := models.Transaction{}
    id := c.Param("id")

    res := database.DB.Where("id = ?", id).First(&transaction)
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "Transaction not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "Transaction retreived", transaction)
}

// Filters and Pagination
func GetTransactions(c echo.Context) error {
    transactions := []models.Transaction{}

    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    // Filters
    moneyBucketID := c.QueryParam("money_bucket_id")
    categoryID := c.QueryParam("category_id")

    query := database.DB.Model(&models.Transaction{})

    // Filter check

    if moneyBucketID != "" {
        query = query.Where("money_bucket_id = ?", moneyBucketID)
    }

    if categoryID != "" {
        query = query.Where("category_id = ?", categoryID)
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

    res := query.Find(&transactions)

    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, "Transactions retreived", transactions)
}

// Update
func UpdateTransaction(c echo.Context) error {
    id := c.Param("id")

    var updates map[string]any

    if err := c.Bind(&updates); err != nil {
        return err
    }

    if email, exists := updates["email"]; exists && email != "" {
        updates["email_verified"] = false
    }

    res := database.DB.Model(&models.Transaction{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update transaction")
    }

    var transaction models.Transaction
    database.DB.First(&transaction, id)

    return Response(c, http.StatusOK, "Transaction updated", transaction)
}

// Delete
// Non permanent
func SoftDeleteTransaction(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Delete(&models.Transaction{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}

// Permanent
func HardDeleteTransaction(c echo.Context) error {
    id := c.Param("id")

    res := database.DB.Unscoped().Delete(&models.Transaction{}, id)
    if res.Error != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
    }

    return Response(c, http.StatusOK, id)
}
