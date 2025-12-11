# Spec Module

This module handles loading and parsing API specifications from different formats.

## Architecture

- `base.go` - Core interfaces and format detection
- `openapi3.go` - OpenAPI 3.0 specification processing
- `swagger2.go` - Swagger 2.0 specification processing

## Interface

```go
type SpecProcessor interface {
    ProcessFile(filename string) ([]EndpointCases, error)
}
```

## Data Types

- `EndpointCases` - Collection of test cases for an endpoint
- `ParameterCase` - Individual parameter with metadata

## Adding New Formats

1. Implement the `SpecProcessor` interface
2. Add format detection in `DetectSpecVersion()`
3. Register in `GetProcessor()` function

## Supported Formats

- **OpenAPI 3.0**: YAML and JSON
- **Swagger 2.0**: JSON

Format is automatically detected from file content.
