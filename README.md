# Go Project Template Generator

A powerful code generator for creating Go backend projects with Clean Architecture principles. This project uses [Plop.js](https://plopjs.com/) to scaffold complete CRUD applications with a well-structured, production-ready architecture.

## 🚀 What is this project?

This is a **project template generator** that creates complete Go backend applications following Clean Architecture principles. Instead of manually setting up the same boilerplate code every time you start a new Go project, this generator creates a fully functional backend with:

- **Clean Architecture** structure (Domain, Application, Infrastructure, API layers)
- **Complete CRUD operations** for any entity
- **RESTful API endpoints** with proper HTTP handlers
- **In-memory repository** implementation (easily extensible to databases)
- **Dependency injection** setup
- **Configuration management**
- **Project structure** following Go best practices

## 🏗️ Generated Project Structure

When you use this generator, it creates a project with this structure:

```
your-project/
├── api/                             # API layer (HTTP handlers, routing)
│   ├── api_router.go               # Main API router
│   └── {entity}/                   # Entity-specific handlers
│       ├── handler.go              # Base handler struct
│       ├── create_{entity}_handler.go
│       ├── get_{entity}_handler.go
│       ├── list_{entity}_handler.go
│       ├── update_{entity}_handler.go
│       ├── delete_{entity}_handler.go
│       └── {entity}_router.go      # Entity routes
├── applications/                    # Application layer (use cases)
│   ├── deps/                       # Dependency interfaces
│   │   ├── repositories/           # Repository interfaces
│   │   └── services/              # Service interfaces
│   └── {entity}/                   # Entity use cases
│       ├── create_{entity}.go
│       ├── get_{entity}.go
│       ├── list_{entity}.go
│       ├── update_{entity}.go
│       └── delete_{entity}.go
├── domain/                         # Domain layer (business logic)
│   ├── entities/                   # Business entities
│   │   └── {entity}.go
│   └── services/                   # Domain services
│       └── {entity}_service.go
├── infrastructure/                 # Infrastructure layer
│   ├── inMemoryRepositories/      # Data persistence
│   │   └── {entity}_repository.go
│   └── utilityServices/           # Utilities
│       └── uuid_generator.go
├── wiring/                         # Dependency injection
│   └── setupApp.go
├── main.go                         # Application entry point
├── go.mod                          # Go module file
├── .env                           # Environment variables
└── README.md                      # Project documentation
```

## 🛠️ How to Use

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- Go 1.21+ (for running generated projects)

### Step 1: Install Dependencies

Navigate to the generator directory and install the required packages:

```bash
cd generator
npm install
```

### Step 2: Generate a New Project

Run the generator to create a new Go project:

```bash
npm run generate
# or use the seed command directly
npm run seed
# or with npx
npx plop seed
# or just
npx plop
```

You'll be prompted for:

1. **Directory name**: The name of your project directory (e.g., `src/my-ecommerce-api`, `src/user-management-service`)
   - **💡 Recommended**: Use `src/your-project-name` to keep generated projects organized
2. **Module/Feature name**: The main entity for your CRUD operations (e.g., `product`, `user`, `order`)

### Available Commands

- `npm run generate` - Interactive mode to select and run any generator
- `npm run seed` - Directly run the seed generator to create a complete project
- `npx plop seed` - Alternative way to run the seed generator directly
- `npx plop` - Interactive mode using npx

### Step 3: Example Usage

**💡 Best Practice**: It's recommended to generate projects in the `src/` directory to keep your generated projects organized and separate from the generator itself.

```bash
$ npm run seed
# or
$ npx plop seed

? Directory name where layers will be created: src/my-shop-api
? Module/Feature name: product
```

This creates a complete project at `src/my-shop-api/` with full CRUD operations for `product` entities.

### Step 4: Run Your Generated Project

```bash
cd src/my-shop-api
go mod tidy
air
```

Your API will be available at `http://localhost:3000` with these endpoints:

- `POST /api/v1/products` - Create a product
- `GET /api/v1/products` - List all products
- `GET /api/v1/products/:id` - Get a specific product
- `PUT /api/v1/products/:id` - Update a product
- `DELETE /api/v1/products/:id` - Delete a product

## 📁 What Gets Generated

### Core Files

- **main.go**: Application entry point with server setup
- **go.mod**: Go module with necessary dependencies
- **.env**: Environment configuration
- **README.md**: Comprehensive project documentation

### Clean Architecture Layers

1. **Domain Layer**: Business entities and domain services
2. **Application Layer**: Use cases and business logic
3. **Infrastructure Layer**: Data persistence and external services
4. **API Layer**: HTTP handlers and routing

### Complete CRUD Operations

For each entity, the generator creates:

- Create, Read, Update, Delete, List operations
- HTTP handlers for each operation
- Use cases with proper error handling
- Repository interfaces and implementations
- Proper routing and middleware setup

## 🔧 Customization

### Adding New Templates

1. Create new `.hbs` templates in the `templates/` directory
2. Add the template to the `actions` array in `plopfile.js`
3. Use Handlebars helpers like `{{pascalCase}}`, `{{kebabCase}}`, `{{snakeCase}}`

### Available Handlebars Helpers

- `{{pascalCase text}}`: Converts to PascalCase
- `{{camelCase text}}`: Converts to camelCase
- `{{snakeCase text}}`: Converts to snake_case
- `{{kebabCase text}}`: Converts to kebab-case
- `{{getModuleName directoryName}}`: Extracts module name from path

### Extending Templates

The templates use Handlebars syntax and have access to these variables:

- `directoryName`: The project directory name
- `moduleName`: The entity/feature name
- `modules`: Array of module names
- `layers`: Array of architectural layers

## 🌟 Features

- **🏗️ Clean Architecture**: Proper separation of concerns
- **📦 Complete CRUD**: All basic operations included
- **🔧 Dependency Injection**: Proper DI setup with interfaces
- **📝 Documentation**: Auto-generated README and API docs
- **🎯 Production Ready**: Follows Go best practices
- **🔄 Extensible**: Easy to add new templates and customize
- **⚡ Fast Setup**: Get a working API in minutes

## 📋 Best Practices

### Project Organization

- **Generate in src/ directory**: Always create new projects in the `src/` directory to keep them organized and separate from the generator code
- **Consistent naming**: Use descriptive names like `src/ecommerce-api`, `src/user-service`, etc.
- **Module naming**: Choose singular entity names for modules (e.g., `product`, `user`, not `products` or `users`)

### Example Directory Structure

```
go-project-template/
├── generator/          # Generator source code
├── src/               # Generated projects (recommended)
│   ├── ecommerce-api/ # Your first project
│   ├── user-service/  # Your second project
│   └── inventory-api/ # Your third project
└── README.md
```

## 🚀 Examples of Generated Projects

In the `src/` directory, you can see examples of generated projects that demonstrate best practices:

- `src/simple/`: A complete example project with `product` entities showing the full Clean Architecture implementation

This serves as a working example of what the generator produces and demonstrates the recommended project organization structure. You can explore this project to understand the generated code structure before creating your own projects.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/new-template`)
3. Add your templates or improvements
4. Commit your changes (`git commit -m 'Add new template'`)
5. Push to the branch (`git push origin feature/new-template`)
6. Open a Pull Request
