package main

import (
	"architecture/controller/product"
	"architecture/data/entities"
	"architecture/data/repositories"
	"architecture/domain/usecases"
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

	db.AutoMigrate(&entities.Product{})

	productRepository := repositories.NewProductRepository(db)
	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	getProductByIdUseCase := usecases.NewGetProductByIdUseCase(productRepository)
	getProductsUseCase := usecases.NewGetProductsUseCase(productRepository)
	productController := adapters.NewProductController(
		createProductUseCase,
		getProductByIdUseCase,
		getProductsUseCase,
	)

	app.Post("/product", productController.CreateProduct)
	app.Get("/product/:id", productController.GetProductById)
	app.Get("/product", productController.GetProducts)

	app.Listen(":8080")
}
