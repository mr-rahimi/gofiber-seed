# Src/simple Back-end

This is a src/simple back-end project demonstrating clean architecture principles with Go, Fiber, and in-memory data storage.

## Project Structure

```
src/
├── domain/                          # Domain layer (business entities, value objects, domain services)
│   ├── entities/                    # Business entities
│   ├── value_objects/               # Value objects (email, etc.)
│   └── services/                    # Domain services
├── applications/                    # Application layer (use cases, business logic)
│   ├── deps/                        # Dependencies interfaces
│   │   ├── repositories/            # Repository interfaces
│   │   └── services/                # Service interfaces
│   └── product/    # Product-related use cases
├── infrastructure/                  # Infrastructure layer (external concerns)
│   ├── inMemoryRepositories/       # In-memory data storage implementations
│   ├── externalServices/            # External API integrations
│   └── utilityServices/             # Utility services (ID generation, etc.)
├── api/                            # API layer (HTTP handlers, middleware)
│   ├── middleware/                  # HTTP middleware
│   └── product/    # Product API handlers
└── wiring/                         # Dependency injection setup
```
