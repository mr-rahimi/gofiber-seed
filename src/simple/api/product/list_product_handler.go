package product

import (
    "context"
    "time"
    
    "github.com/gofiber/fiber/v2"
)

// ListProduct handles GET /product
func (h *Handler) ListProduct(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    results, err := h.listProductUC.Execute(ctx)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to retrieve products",
            "details": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "data": results,
    })
}
