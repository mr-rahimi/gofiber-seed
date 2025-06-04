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

## Features Demonstrated

### Domain Layer

- **Entities**: Product
- **Value Objects**: Email with validation
- **Domain Services**: Product domain logic

### Application Layer

- **Use Cases**: CreateProduct, GetProduct, GetProducts
- **Error Handling**: Custom application errors
- **Dependencies**: Clean interfaces for repositories and services

### Infrastructure Layer

- **Repositories**: In-memory repository implementation
- **External Services**: Third-party API integration
- **Utility Services**: ID generation

### API Layer

- **Handlers**: RESTful product management endpoints
- **Middleware**: JWT authentication, role-based access control
- **Error Handling**: Proper HTTP status codes and error responses

## API Endpoints

### Public Endpoints

- `GET /health` - Health check
- `POST /api/v1/products` - Create product
- `GET /api/v1/docs` - API documentation

### Protected Endpoints (Requires JWT)

- `GET /api/v1/products` - Get products (paginated)
- `GET /api/v1/products/:id` - Get product by ID

### Admin Endpoints (Requires admin role)

- `GET /api/v1/admin/products` - Admin product management

## Getting Started

### Prerequisites

- Go 1.21+

### Installation

1. Clone the repository

```bash
git clone <repository-url>
cd src/simple
```

2. Install dependencies

```bash
go mod tidy
```

3. Set up environment variables

```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the application

```bash
go run main.go
```

The server will start on port 8080 by default.

### Environment Variables

```env
# Server settings
PORT=8080

# JWT settings (if using authentication)
JWT_SECRET_KEY=your-secret-key-here

# External API Settings (Example)
EXTERNAL_API_URL=https://api.example.com
EXTERNAL_API_KEY=your-api-key-here
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o bin/src/simple main.go
```

### Docker Support

```bash
docker build -t src/simple .
docker run -p 8080:8080 src/simple
```

## Architecture

This project follows Clean Architecture principles:

- **Domain**: Contains business logic, entities, and domain services
- **Application**: Contains use cases and application services
- **Infrastructure**: Contains external integrations and data persistence
- **API**: Contains HTTP handlers and routing

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
