package product

import (
    "context"
    "time"
    
    "github.com/gofiber/fiber/v2"
)

// GetProduct handles GET /product/{id}
func (h *Handler) GetProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID parameter is required",
        })
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    result, err := h.getProductUC.Execute(ctx, id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Product not found",
            "details": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "data": result,
    })
}
