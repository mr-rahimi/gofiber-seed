package product

import (
    "context"
    "simple/applications/deps/repositories"
    "simple/domain/entities"
    domainServices "simple/domain/services"
)

type UpdateProductUseCase struct {
    productRepository repositories.ProductRepository
    productDomainService *domainServices.ProductDomainService
}

func NewUpdateProductUseCase(
    productRepository repositories.ProductRepository,
    productDomainService *domainServices.ProductDomainService,
) *UpdateProductUseCase {
    return &UpdateProductUseCase{
        productRepository:    productRepository,
        productDomainService: productDomainService,
    }
}

func (uc *UpdateProductUseCase) Execute(ctx context.Context, product *entities.Product) (*entities.Product, error) {
    // Check if product exists
    existing, err := uc.productRepository.GetByID(ctx, product.ID)
    if err != nil {
        return nil, err
    }
    if existing == nil {
        return nil, err // Product not found
    }

    // Domain validation
    if err := uc.productDomainService.ValidateProduct(product); err != nil {
        return nil, err
    }

    // Update the product
    updatedProduct, err := uc.productRepository.Update(ctx, product)
    if err != nil {
        return nil, err
    }

    return updatedProduct, nil
}
