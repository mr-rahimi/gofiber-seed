package api

import (
	"simple/wiring"
	"simple/api/product"

	"github.com/gofiber/fiber/v2"
)

func ConfigApiRoutes(app *fiber.App, useCases *wiring.UseCases) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Src/simple API is running",
		})
	})

	// API version 1
	v1 := app.Group("/api/v1")

	// Product routes
	product.SetupProductRoutes(v1, useCases)

	// API documentation endpoint
	app.Get("/api/v1/docs", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Src/simple API Documentation",
			"version": "1.0.0",
			"endpoints": map[string]interface{}{
				"product": map[string]string{
					"GET /api/v1/product":        "List products",
					"GET /api/v1/product/:id":    "Get product by ID",
					"POST /api/v1/product":       "Create new product",
					"PUT /api/v1/product/:id":    "Update product",
					"DELETE /api/v1/product/:id": "Delete product",
				},
			},
		})
	})
}
