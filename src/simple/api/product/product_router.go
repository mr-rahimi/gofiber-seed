package product

import (
    "simple/wiring"
    "github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(router fiber.Router, useCases *wiring.UseCases) {
    // Product handler
    handler := NewHandler(useCases)
    
    // Product routes
    router.Post("/product", handler.CreateProduct)
    router.Get("/product", handler.ListProduct)
    router.Get("/product/:id", handler.GetProduct)
    router.Put("/product/:id", handler.UpdateProduct)
    router.Delete("/product/:id", handler.DeleteProduct)
}
