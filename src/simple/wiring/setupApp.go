package wiring

import (
	"simple/applications/deps/repositories"
	"simple/applications/deps/services"
	"simple/applications/product"
	domainServices "simple/domain/services"
	"simple/infrastructure/inMemoryRepositories"
	"simple/infrastructure/utilityServices"
)

type UseCases struct {
	CreateProductUseCase *product.CreateProductUseCase
	GetProductUseCase    *product.GetProductUseCase
	UpdateProductUseCase *product.UpdateProductUseCase
	DeleteProductUseCase *product.DeleteProductUseCase
	ListProductUseCase   *product.ListProductUseCase
}

type Repositories struct {
	ProductRepository repositories.ProductRepository
}

type Services struct {
	IDGenerator        services.IDGenerator
}

type DomainServices struct {
	ProductDomainService *domainServices.ProductDomainService
}

func SetupUseCases() (*UseCases, error) {
	// Setup repositories (using in-memory)
	repos := &Repositories{
		ProductRepository: inMemoryRepositories.NewProductRepository(),
	}

	// Setup infrastructure services
	services := &Services{
		IDGenerator:        utilityServices.NewUUIDGenerator(),
	}

	// Setup domain services
	domainServices := &DomainServices{
		ProductDomainService: domainServices.NewProductDomainService(),
	}

	// Setup use cases
	useCases := &UseCases{
		CreateProductUseCase: product.NewCreateProductUseCase(
			repos.ProductRepository,
			services.IDGenerator,
			domainServices.ProductDomainService,
		),
		GetProductUseCase: product.NewGetProductUseCase(
			repos.ProductRepository,
		),
		UpdateProductUseCase: product.NewUpdateProductUseCase(
			repos.ProductRepository,
			domainServices.ProductDomainService,
		),
		DeleteProductUseCase: product.NewDeleteProductUseCase(
			repos.ProductRepository,
		),
		ListProductUseCase: product.NewListProductUseCase(
			repos.ProductRepository,
		),
	}
	return useCases, nil
}
