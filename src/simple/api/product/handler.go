package product

import (
    "simple/applications/product"
    "simple/wiring"
)

type Handler struct {
    createProductUC *product.CreateProductUseCase
    getProductUC    *product.GetProductUseCase
    updateProductUC *product.UpdateProductUseCase
    deleteProductUC *product.DeleteProductUseCase
    listProductUC   *product.ListProductUseCase
}

func NewHandler(useCases *wiring.UseCases) *Handler {
    return &Handler{
        createProductUC: useCases.CreateProductUseCase,
        getProductUC:    useCases.GetProductUseCase,
        updateProductUC: useCases.UpdateProductUseCase,
        deleteProductUC: useCases.DeleteProductUseCase,
        listProductUC:   useCases.ListProductUseCase,
    }
}
