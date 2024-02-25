package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"product/internal/config"
	"product/internal/controller"
	"product/internal/repositories"
	"product/internal/routes"
	"product/internal/services"
)

func StartApp() {
	app := fiber.New()

	app.Use(cors.New())

	validator := validator.New()

	db := config.OpenConnection()

	repository := repositories.NewProductRepository()
	service := services.NewProductService(db, repository)
	productController := controller.NewProductController(service, validator)

	api := app.Group("api")
	app.Static("/public/product", "./public/product")
	routes.ProductRoutes(api, productController)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
