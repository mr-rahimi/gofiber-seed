package inMemoryRepositories

import (
    "context"
    "errors"
    "fmt"
    "sync"
    "time"
    "simple/domain/entities"
    "simple/applications/deps/repositories"
)

type InMemoryProductRepository struct {
    data map[string]*entities.Product
    mu   sync.RWMutex
}

func NewProductRepository() repositories.ProductRepository {
    return &InMemoryProductRepository{
        data: make(map[string]*entities.Product),
    }
}

func (r *InMemoryProductRepository) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    // Generate ID if not provided
    if product.ID == "" {
        product.ID = r.generateID()
    }
    
    // Set timestamps
    product.CreatedAt = time.Now()
    product.UpdatedAt = time.Now()
    
    r.data[product.ID] = product
    return product, nil
}

func (r *InMemoryProductRepository) GetByID(ctx context.Context, id string) (*entities.Product, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    product, exists := r.data[id]
    if !exists {
        return nil, errors.New("product not found")
    }
    
    return product, nil
}

func (r *InMemoryProductRepository) List(ctx context.Context) ([]*entities.Product, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    result := make([]*entities.Product, 0, len(r.data))
    for _, product := range r.data {
        result = append(result, product)
    }
    
    return result, nil
}

func (r *InMemoryProductRepository) Update(ctx context.Context, product *entities.Product) (*entities.Product, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.data[product.ID]; !exists {
        return nil, errors.New("product not found")
    }
    
    // Update timestamp
    product.UpdatedAt = time.Now()
    
    r.data[product.ID] = product
    return product, nil
}

func (r *InMemoryProductRepository) Delete(ctx context.Context, id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.data[id]; !exists {
        return errors.New("product not found")
    }
    
    delete(r.data, id)
    return nil
}

// generateID generates a simple ID
func (r *InMemoryProductRepository) generateID() string {
    return fmt.Sprintf("product_%d", time.Now().UnixNano())
}
