package entities

import (
    "time"
)

type Product struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Name      string    `json:"name" bson:"name"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewProduct(name string) *Product {
    now := time.Now()
    return &Product{
        Name:      name,
        CreatedAt: now,
        UpdatedAt: now,
    }
}

// IsValid validates the product entity
func (product *Product) IsValid() bool {
    return product.Name != ""
}

// Update updates the product fields
func (product *Product) Update(name string) {
    product.Name = name
    product.UpdatedAt = time.Now()
}
