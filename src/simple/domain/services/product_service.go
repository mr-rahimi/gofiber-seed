package services

import (
    "errors"
    "simple/domain/entities"
)

type ProductDomainService struct {
    // Add dependencies if needed
}

func NewProductDomainService() *ProductDomainService {
    return &ProductDomainService{}
}

// ValidateProduct validates the product entity
func (s *ProductDomainService) ValidateProduct(product *entities.Product) error {
    if !product.IsValid() {
        return errors.New("invalid product entity")
    }
    
    // Add more domain-specific validation rules here
    
    return nil
}
