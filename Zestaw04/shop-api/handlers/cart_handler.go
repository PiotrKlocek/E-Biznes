package handlers

import (
	"net/http"
	"strconv"

	"shop-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartHandler struct {
	DB *gorm.DB
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (h *CartHandler) CreateCart(c echo.Context) error {
	cart := models.Cart{}

	if err := h.DB.Create(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create cart",
		})
	}

	return c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) GetCart(c echo.Context) error {
	id := c.Param("id")

	var cart models.Cart
	if err := h.DB.
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "cart not found",
		})
	}

	return c.JSON(http.StatusOK, cart)
}
func (h *CartHandler) AddToCart(c echo.Context) error {
	cartIDStr := c.Param("id")
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid cart id",
		})
	}

	var req AddToCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}

	if req.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "quantity must be greater than 0",
		})
	}

	var cart models.Cart
	if err := h.DB.First(&cart, uint(cartID)).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "cart not found",
		})
	}

	var product models.Product
	if err := h.DB.First(&product, req.ProductID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found",
		})
	}

	item := models.CartItem{
		CartID:    uint(cartID),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.DB.Create(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not add item to cart",
		})
	}

	h.DB.Preload("Product").First(&item, item.ID)

	return c.JSON(http.StatusCreated, item)
}