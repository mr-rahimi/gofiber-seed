package product

import (
    "context"
    "simple/domain/entities"
    "simple/applications/deps/repositories"
)

type ListProductUseCase struct {
    productRepository repositories.ProductRepository
}

func NewListProductUseCase(productRepository repositories.ProductRepository) *ListProductUseCase {
    return &ListProductUseCase{
        productRepository: productRepository,
    }
}

func (uc *ListProductUseCase) Execute(ctx context.Context) ([]*entities.Product, error) {
    return uc.productRepository.List(ctx)
}
