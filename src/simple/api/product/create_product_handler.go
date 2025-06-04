package product

import (
    "context"
    "time"
    
    "simple/domain/entities"
    
    "github.com/gofiber/fiber/v2"
)

// CreateProduct handles POST /product
func (h *Handler) CreateProduct(c *fiber.Ctx) error {
    var entity entities.Product
    if err := c.BodyParser(&entity); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error":   "Invalid request body",
            "details": err.Error(),
        })
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    result, err := h.createProductUC.Execute(ctx, &entity)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to create product",
            "details": err.Error(),
        })
    }

    return c.Status(201).JSON(fiber.Map{
        "message": "Product created successfully",
        "data":    result,
    })
}
