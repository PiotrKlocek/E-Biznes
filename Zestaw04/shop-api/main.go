package main

import (
	"shop-api/db_config"
	"shop-api/handlers"
	"shop-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func addData(db *gorm.DB) {
	var count int64

	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return
	}

	categories := []models.Category{
		{Name: "Elektronika"},
		{Name: "Książki"},
		{Name: "Odzież"},
		{Name: "Sport"},
	}

	for i := range categories {
		db.Create(&categories[i])
	}
	//wprowadzone dane testowe
	products := []models.Product{
		{Name: "Laptop", Price: 4999.99, CategoryID: categories[0].ID},
		{Name: "Smartfon", Price: 2999.99, CategoryID: categories[0].ID},
		{Name: "Słuchawki", Price: 199.99, CategoryID: categories[0].ID},
		{Name: "Monitor", Price: 899.99, CategoryID: categories[0].ID},

		{Name: "Książka Go", Price: 79.99, CategoryID: categories[1].ID},
		{Name: "Książka Python", Price: 69.99, CategoryID: categories[1].ID},

		{Name: "Koszulka", Price: 49.99, CategoryID: categories[2].ID},
		{Name: "Bluza", Price: 129.99, CategoryID: categories[2].ID},

		{Name: "Piłka", Price: 89.99, CategoryID: categories[3].ID},
		{Name: "Rower", Price: 1999.99, CategoryID: categories[3].ID},
	}

	for i := range products {
		db.Create(&products[i])
	}

	cart := models.Cart{}
	db.Create(&cart)

	items := []models.CartItem{
		{CartID: cart.ID, ProductID: 1, Quantity: 1},
		{CartID: cart.ID, ProductID: 2, Quantity: 2},
		{CartID: cart.ID, ProductID: 5, Quantity: 1},
	}

	for i := range items {
		db.Create(&items[i])
	}
}

func main() {
	db := db_config.InitDB()

	err := db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
	)

	addData(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	productHandler := &handlers.ProductHandler{DB: db}
	categoryHandler := &handlers.CategoryHandler{DB: db}
	cartHandler := &handlers.CartHandler{DB: db}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Shop API is running")
	})

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProductByID)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	e.POST("/categories", categoryHandler.CreateCategory)
	e.GET("/categories", categoryHandler.GetCategories)
	e.GET("/categories/:id", categoryHandler.GetCategoryByID)

	e.POST("/carts", cartHandler.CreateCart)
	e.GET("/carts/:id", cartHandler.GetCart)
	e.POST("/carts/:id/items", cartHandler.AddToCart)

	e.Logger.Fatal(e.Start(":8080"))
}