package product

import (
    "context"
    "time"
    
    "github.com/gofiber/fiber/v2"
)

// DeleteProduct handles DELETE /product/{id}
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID parameter is required",
        })
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err := h.deleteProductUC.Execute(ctx, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to delete product",
            "details": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Product deleted successfully",
    })
}
