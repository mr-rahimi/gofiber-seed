package product

import (
    "context"
    "simple/applications/deps/repositories"
)

type DeleteProductUseCase struct {
    productRepository repositories.ProductRepository
}

func NewDeleteProductUseCase(productRepository repositories.ProductRepository) *DeleteProductUseCase {
    return &DeleteProductUseCase{
        productRepository: productRepository,
    }
}

func (uc *DeleteProductUseCase) Execute(ctx context.Context, id string) error {
    // Check if product exists
    existing, err := uc.productRepository.GetByID(ctx, id)
    if err != nil {
        return err
    }
    if existing == nil {
        return err // Product not found
    }

    // Delete the product
    return uc.productRepository.Delete(ctx, id)
}
