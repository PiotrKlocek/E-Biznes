package handlers

import (
	"net/http"
	"strconv"

	"shop-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

type ProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req ProductRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "name is required",
		})
	}

	product := models.Product{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	if req.CategoryID != 0 {
		var category models.Category
		if err := h.DB.First(&category, req.CategoryID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "category not found",
			})
		}
	}

	if err := h.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create product",
		})
	}

	h.DB.Preload("Category").First(&product, product.ID)

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	var products []models.Product

	minPriceStr := c.QueryParam("min_price")
	categoryIDStr := c.QueryParam("category_id")
	name := c.QueryParam("name")

	var minPrice float64
	var categoryID uint

	if minPriceStr != "" {
		parsed, err := strconv.ParseFloat(minPriceStr, 64)
		if err == nil {
			minPrice = parsed
		}
	}

	if categoryIDStr != "" {
		parsed, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err == nil {
			categoryID = uint(parsed)
		}
	}

	query := h.DB.Preload("Category").Scopes(
		withMinPrice(minPrice),
		withCategory(categoryID),
		withNameLike(name),
	)

	if err := query.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not fetch products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := h.DB.Preload("Category").First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found",
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found",
		})
	}

	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "name is required",
		})
	}

	if req.CategoryID != 0 {
		var category models.Category
		if err := h.DB.First(&category, req.CategoryID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "category not found",
			})
		}
	}

	product.Name = req.Name
	product.Price = req.Price
	product.CategoryID = req.CategoryID

	if err := h.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not update product",
		})
	}

	h.DB.Preload("Category").First(&product, product.ID)

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found",
		})
	}

	if err := h.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not delete product",
		})
	}

	return c.NoContent(http.StatusNoContent)
}


func withMinPrice(minPrice float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if minPrice <= 0 {
			return db
		}
		return db.Where("price >= ?", minPrice)
	}
}

func withCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID == 0 {
			return db
		}
		return db.Where("category_id = ?", categoryID)
	}
}

func withNameLike(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		return db.Where("name LIKE ?", "%"+name+"%")
	}
}