package main

import (
	"architecture/adapters/product"
	"architecture/core/product"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(sqlite.Open("architecture.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&core.Product{})

	productRepository := adapters.NewProductRepository(db)
	productService := core.NewProductService(productRepository)
	productHandler := adapters.NewHttpProductHandler(productService)

	app.Post("/product", productHandler.CreateProduct)
	app.Get("/product/:id", productHandler.GetProductById)
	app.Get("/product", productHandler.GetProducts)

	app.Listen(":8080")
}
