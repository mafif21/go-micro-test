package routes

import (
	"github.com/gofiber/fiber/v2"
	"product/internal/controller"
)

func ProductRoutes(router fiber.Router, controller controller.ProductController) {
	product := router.Group("/product")

	product.Get("/get", controller.FindAll)
	product.Get("/get/:id", controller.FindById)
	product.Post("/create", controller.Create)
	product.Put("/update/:id", controller.Update)
	product.Delete("/delete/:id", controller.Delete)
}
