# openapi-casegen

openapi-casegen is a developer tool that automatically generates test case definitions from your OpenAPI specification. Its goal is to ensure that every documented endpoint, parameter, and request field in your API has a corresponding test — eliminating the "did I test this input?" uncertainty in backend development.

## Features

- ✅ **Multi-format support**: OpenAPI 3.0 (YAML/JSON) and Swagger 2.0 (JSON)
- ✅ **Data type aware**: Reads actual schema types, not just descriptions
- ✅ **Modular generators**: Easy to extend with new data types
- ✅ **Comprehensive coverage**: Generates valid, invalid, boundary, and basic access test cases
- ✅ **Complete endpoint coverage**: Every endpoint gets at least one test case

## Quick Start

```bash
# Generate test cases from an OpenAPI 3.0 specification
./openapi-casegen examples/openapi.yaml

# Generate test cases from a Swagger 2.0 specification
./openapi-casegen examples/swagger.json

# Validate generated test cases against JUnit XML results
./openapi-casegen examples/openapi.yaml examples/results.xml
```

## Architecture

The tool is organized into three main modules for clean separation of concerns:

### 1. **Spec Module** (`spec/`)
Handles loading and parsing API specifications:

- `spec/base.go` - Interfaces and format detection
- `spec/openapi3.go` - OpenAPI 3.0 specification processing
- `spec/swagger2.go` - Swagger 2.0 specification processing

### 2. **Generators Module** (`generators/`)
Handles test case generation for different data types:

- `generators/base.go` - Main dispatcher and interfaces
- `generators/enum.go` - Enum parameter test cases
- `generators/integer.go` - Integer parameter test cases
- `generators/number.go` - Float parameter test cases
- `generators/string.go` - String parameter test cases
- `generators/boolean.go` - Boolean parameter test cases

### 3. **Validators Module** (`validator/`)
Validates test implementation against JUnit XML results:

- `validator/base.go` - JUnit XML parser and test comparison logic

### 4. **Main Application** (`main.go`)
Orchestrates the processing pipeline: Spec → Generators → Validator → Output

### 5. **Examples** (`examples/`)
Sample API specifications and test results for testing:

- `examples/openapi.yaml` - OpenAPI 3.0 specification
- `examples/swagger.json` - Swagger 2.0 specification
- `examples/results.xml` - Sample JUnit XML test results

## Test Case Types

The tool generates comprehensive test cases for different scenarios:

### **Endpoint Access**
- `{endpoint}.{method}_basic_access` - Basic endpoint accessibility test

### **Parameter Validation**
- `{endpoint}.{param}_valid_input` - Valid parameter values
- `{endpoint}.{param}_invalid_input` - Invalid parameter values

### **Boundary Testing**
- `{endpoint}.{param}_boundary_min` - Minimum boundary values (for numbers)
- `{endpoint}.{param}_boundary_max` - Maximum boundary values (for numbers)

### **Enum Testing**
- `{endpoint}.{param}_valid_{enum_value}` - Each allowed enum value
- `{endpoint}.{param}_invalid_input` - Invalid enum values

## Adding New Components

### New Data Type Generator
```go
// Create generators/newtype.go
type NewTypeGenerator struct{}

func (g *NewTypeGenerator) GenerateTestCases(baseID, paramType string, enumValues []interface{}) []TestCase {
    // Your test case logic here
}

// Add to generators/base.go switch statement
case "newtype":
    generator = &NewTypeGenerator{}
```

### New Specification Processor
```go
// Create processor/newformat.go
type NewFormatProcessor struct{}

func (p *NewFormatProcessor) ProcessFile(filename string) ([]EndpointCases, error) {
    // Your processing logic here
}

// Add to processor/base.go GetProcessor function
case "newformat":
    return &NewFormatProcessor{}
```

## Future Extensions

The modular design makes it easy to add:
- **`result/` Module**: Test result matching and validation
- **`reporter/` Module**: Different output formats (JUnit, HTML, etc.)
- **`validator/` Module**: API specification validation
- **`transformer/` Module**: Convert between specification formats

Each module follows a clear naming convention based on its primary responsibility.

---

## 🤖 AI-Generated Tool

This tool was created with the assistance of AI through an interactive development process. The AI helped design implement features, and organize the codebase while ensuring best practices and comprehensive functionality.

*Built with Cursor IDE and AI collaboration*
