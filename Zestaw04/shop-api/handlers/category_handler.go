package handlers

import (
	"net/http"

	"shop-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var category models.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}

	if category.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "name is required",
		})
	}

	if err := h.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create category",
		})
	}

	return c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategories(c echo.Context) error {
	var categories []models.Category

	if err := h.DB.Preload("Products").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not fetch categories",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategoryByID(c echo.Context) error {
	id := c.Param("id")

	var category models.Category
	if err := h.DB.Preload("Products").First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "category not found",
		})
	}

	return c.JSON(http.StatusOK, category)
}