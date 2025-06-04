package product

import (
    "context"
    "simple/applications/deps/repositories"
    "simple/applications/deps/services"
    "simple/domain/entities"
    domainServices "simple/domain/services"
)

type CreateProductUseCase struct {
    productRepository repositories.ProductRepository
    idGenerator             services.IDGenerator
    productDomainService *domainServices.ProductDomainService
}

func NewCreateProductUseCase(
    productRepository repositories.ProductRepository,
    idGenerator services.IDGenerator,
    productDomainService *domainServices.ProductDomainService,
) *CreateProductUseCase {
    return &CreateProductUseCase{
        productRepository:    productRepository,
        idGenerator:             idGenerator,
        productDomainService: productDomainService,
    }
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, product *entities.Product) (*entities.Product, error) {
    // Generate ID if not provided
    if product.ID == "" {
        product.ID = uc.idGenerator.GenerateID()
    }

    // Domain validation
    if err := uc.productDomainService.ValidateProduct(product); err != nil {
        return nil, err
    }

    // Create the product
    createdProduct, err := uc.productRepository.Create(ctx, product)
    if err != nil {
        return nil, err
    }

    return createdProduct, nil
}
