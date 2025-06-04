package product

import (
    "context"
    "time"
    
    "simple/domain/entities"
    
    "github.com/gofiber/fiber/v2"
)

// UpdateProduct handles PUT /product/{id}
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID parameter is required",
        })
    }

    var entity entities.Product
    if err := c.BodyParser(&entity); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error":   "Invalid request body",
            "details": err.Error(),
        })
    }

    entity.ID = id

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    result, err := h.updateProductUC.Execute(ctx, &entity)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to update product",
            "details": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Product updated successfully",
        "data":    result,
    })
}
