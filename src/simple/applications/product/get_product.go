package product

import (
    "context"
    "simple/applications/deps/repositories"
    "simple/domain/entities"
)

type GetProductUseCase struct {
    productRepository repositories.ProductRepository
}

func NewGetProductUseCase(
    productRepository repositories.ProductRepository,
) *GetProductUseCase {
    return &GetProductUseCase{
        productRepository: productRepository,
    }
}

func (uc *GetProductUseCase) Execute(ctx context.Context, id string) (*entities.Product, error) {
    // Get from repository
    product, err := uc.productRepository.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    return product, nil
}
