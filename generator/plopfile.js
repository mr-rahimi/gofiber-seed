module.exports = function (plop) {
  // Helper to convert to PascalCase
  plop.setHelper("pascalCase", function (text) {
    return (
      text.charAt(0).toUpperCase() +
      text.slice(1).replace(/[-_\s]+(.)?/g, function (match, chr) {
        return chr ? chr.toUpperCase() : "";
      })
    );
  });

  // Helper to convert to camelCase
  plop.setHelper("camelCase", function (text) {
    const pascal =
      text.charAt(0).toUpperCase() +
      text.slice(1).replace(/[-_\s]+(.)?/g, function (match, chr) {
        return chr ? chr.toUpperCase() : "";
      });
    return pascal.charAt(0).toLowerCase() + pascal.slice(1);
  });

  // Helper to convert to snake_case
  plop.setHelper("snakeCase", function (text) {
    return text
      .replace(/([A-Z])/g, "_$1")
      .toLowerCase()
      .replace(/^_/, "");
  });
  // Helper to convert to kebab-case
  plop.setHelper("kebabCase", function (text) {
    return text
      .replace(/([A-Z])/g, "-$1")
      .toLowerCase()
      .replace(/^-/, "");
  });

  // Helper to extract module name from directory path
  plop.setHelper("getModuleName", function (directoryName) {
    // Remove leading/trailing slashes and split by path separators
    const pathParts = directoryName
      .replace(/^[\/\\]+|[\/\\]+$/g, "")
      .split(/[\/\\]+/);

    // Return the last part of the path
    return pathParts[pathParts.length - 1];
  });
  // Generator for creating complete CRUD feature structure
  plop.setGenerator("seed", {
    description: "Create complete CRUD feature with all architectural layers",
    prompts: [
      {
        type: "input",
        name: "directoryName",
        message:
          "Directory name where layers will be created (e.g., my-app, backend-service):",
        validate: function (value) {
          if (/.+/.test(value)) {
            return true;
          }
          return "Directory name is required";
        },
      },
      {
        type: "input",
        name: "moduleName",
        message: "Module/Feature name (e.g., product, order, user):",
        validate: function (value) {
          if (/.+/.test(value)) {
            return true;
          }
          return "Module name is required";
        },
      },
    ],
    actions: function (data) {
      // Automatically set all required components for complete CRUD functionality
      data.layers = ["api", "application", "domain", "infrastructure"];
      data.apiComponents = ["handler", "router", "policy"];
      data.applicationComponents = ["usecases", "deps"];
      data.domainComponents = ["entities", "valueObjects", "services"];
      data.infrastructureComponents = ["inMemoryRepositories"];

      // Add modules array for template iteration
      data.modules = [data.moduleName];
      const actions = []; // Add foundational files first
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/.env",
        templateFile: "templates/env.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/main.go",
        templateFile: "templates/main.go.hbs",
        skipIfExists: true,
      });
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/wiring/setupApp.go",
        templateFile: "templates/wiring_setup.go.hbs",
        skipIfExists: true,
      });
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/go.mod",
        templateFile: "templates/go.mod.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/README.md",
        templateFile: "templates/README.md.hbs",
        skipIfExists: true,
      }); // Application layer - Dependencies interfaces
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/deps/services/services.go",
        templateFile: "templates/application_services_interface.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/deps/repositories/repositories.go",
        templateFile: "templates/application_repository_interface.go.hbs",
        skipIfExists: true,
      }); // Domain layer
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/domain/entities/{{snakeCase moduleName}}.go",
        templateFile: "templates/domain_entity.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/domain/services/{{snakeCase moduleName}}_service.go",
        templateFile: "templates/domain_service.go.hbs",
        skipIfExists: true,
      }); // Application layer - Use cases
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/{{snakeCase moduleName}}/create_{{snakeCase moduleName}}.go",
        templateFile: "templates/application_create_usecase.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/{{snakeCase moduleName}}/get_{{snakeCase moduleName}}.go",
        templateFile: "templates/application_get_usecase.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/{{snakeCase moduleName}}/update_{{snakeCase moduleName}}.go",
        templateFile: "templates/application_update_usecase.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/{{snakeCase moduleName}}/delete_{{snakeCase moduleName}}.go",
        templateFile: "templates/application_delete_usecase.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/applications/{{snakeCase moduleName}}/list_{{snakeCase moduleName}}.go",
        templateFile: "templates/application_list_usecase.go.hbs",
        skipIfExists: true,
      }); // Infrastructure layer
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/infrastructure/inMemoryRepositories/{{snakeCase moduleName}}_repository.go",
        templateFile: "templates/infrastructure_inmemory_repository.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/infrastructure/utilityServices/uuid_generator.go",
        templateFile: "templates/infrastructure_uuid_generator.go.hbs",
        skipIfExists: true,
      }); // API layer
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/api_router.go",
        templateFile: "templates/api_main_router.go.hbs",
        skipIfExists: true,
      }); // Base handler with struct and constructor
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/handler.go",
        templateFile: "templates/api_handler_base.go.hbs",
        skipIfExists: true,
      });

      // Individual handler files for each operation
      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/create_{{snakeCase moduleName}}_handler.go",
        templateFile: "templates/api_create_handler.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/get_{{snakeCase moduleName}}_handler.go",
        templateFile: "templates/api_get_handler.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/list_{{snakeCase moduleName}}_handler.go",
        templateFile: "templates/api_list_handler.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/update_{{snakeCase moduleName}}_handler.go",
        templateFile: "templates/api_update_handler.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/delete_{{snakeCase moduleName}}_handler.go",
        templateFile: "templates/api_delete_handler.go.hbs",
        skipIfExists: true,
      });

      actions.push({
        type: "add",
        path: "../{{kebabCase directoryName}}/api/{{snakeCase moduleName}}/{{snakeCase moduleName}}_router.go",
        templateFile: "templates/api_router.go.hbs",
        skipIfExists: true,
      });

      return actions;
    },
  });
};
