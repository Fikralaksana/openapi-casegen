# Example API Specifications

This directory contains sample API specifications to demonstrate the capabilities of the openapi-casegen tool.

## Files

### `openapi.yaml`
- **Format**: OpenAPI 3.0 (YAML)
- **Description**: A simple user management API with GET, POST endpoints
- **Features Demonstrated**:
  - Query parameters (limit, offset)
  - Path parameters (user ID)
  - Request body schemas with multiple properties
  - Enum parameters

### `swagger.json`
- **Format**: Swagger 2.0 (JSON)
- **Description**: Sample Petstore API (subset)
- **Features Demonstrated**:
  - Complex object schemas with references (`$ref`)
  - Multiple HTTP methods (GET, POST, PUT, DELETE)
  - Form data parameters
  - Header parameters
  - Array parameters

## Usage

```bash
# Generate test cases from OpenAPI 3.0 spec
../openapi-casegen examples/openapi.yaml

# Generate test cases from Swagger 2.0 spec
../openapi-casegen examples/swagger.json
```

## Expected Output

Both examples will generate comprehensive test case IDs covering:
- Basic endpoint accessibility
- Parameter validation (valid/invalid inputs)
- Boundary testing for numeric parameters
- Enum value testing
- Individual property validation for request bodies
