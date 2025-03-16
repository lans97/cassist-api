package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/gorm"
)

// CRUD Operations for category

// Create
func CreateCategory(c echo.Context) error {
	category := models.Category{}
	err := c.Bind(&category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request: Category struct")
	}

	res := database.DB.Create(&category)
	if res.Error != nil {
		if res.Error == gorm.ErrDuplicatedKey {
			return echo.NewHTTPError(http.StatusConflict, "Duplicate entry")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Category creation failed")
	}

	return Response(c, http.StatusOK, "Category created", category)
}

// Read
// By ID
func GetCategoryById(c echo.Context) error {
	category := models.Category{}
	id := c.Param("id")

	res := database.DB.Where("id = ?", id).First(&category)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Category not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return Response(c, http.StatusOK, "Category retreived", category)
}

// Filters and Pagination
func GetCategories(c echo.Context) error {
	categories := []models.Category{}

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	// Filters
	userID := c.QueryParam("user_id")
	name := c.QueryParam("name")

	query := database.DB.Model(&models.Category{})

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
		query = query.Limit(limit_n).Offset(limit_n * page_n)
	}

	res := query.Find(&categories)

	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return Response(c, http.StatusOK, "Categories retreived", categories)
}

// Update
func UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	var updates map[string]any

	if err := c.Bind(&updates); err != nil {
		return err
	}

	res := database.DB.Model(&models.Category{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update category")
	}

	var category models.Category
	database.DB.First(&category, id)

	return Response(c, http.StatusOK, "Category updated", category)
}

// Delete
// Non permanent
func SoftDeleteCategory(c echo.Context) error {
	id := c.Param("id")

	res := database.DB.Delete(&models.Category{}, id)
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return Response(c, http.StatusOK, id)
}

// Permanent
func HardDeleteCategory(c echo.Context) error {
	id := c.Param("id")

	res := database.DB.Unscoped().Delete(&models.Category{}, id)
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return Response(c, http.StatusOK, id)
}
