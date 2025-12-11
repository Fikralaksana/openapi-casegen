package processor

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/spec"
)

// Swagger2Processor handles Swagger 2.0 specifications
type Swagger2Processor struct{}

// ProcessFile loads and processes a Swagger 2.0 specification file
func (p *Swagger2Processor) ProcessFile(filename string) ([]EndpointCases, error) {
	swagger, err := loadSwagger2Spec(filename)
	if err != nil {
		return nil, err
	}

	return extractEndpointsSwagger2(swagger), nil
}

// loadSwagger2Spec loads a Swagger 2.0 specification
func loadSwagger2Spec(path string) (*spec.Swagger, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var swagger spec.Swagger
	if err := json.Unmarshal(data, &swagger); err != nil {
		return nil, err
	}

	return &swagger, nil
}

// extractEndpointsSwagger2 extracts endpoints from Swagger 2.0 specification
func extractEndpointsSwagger2(swagger *spec.Swagger) []EndpointCases {
	results := []EndpointCases{}

	for path, pathItem := range swagger.Paths.Paths {
		// Handle all operations (GET, POST, PUT, DELETE, etc.)
		operations := map[string]*spec.Operation{
			"get":     pathItem.Get,
			"post":    pathItem.Post,
			"put":     pathItem.Put,
			"delete":  pathItem.Delete,
			"options": pathItem.Options,
			"head":    pathItem.Head,
			"patch":   pathItem.Patch,
		}

		for method, operation := range operations {
			if operation == nil {
				continue
			}

			ec := EndpointCases{
				Endpoint: path,
				Method:   method,
				Cases:    []ParameterCase{},
			}

			// 1. Extract parameters (query, path, header, formData)
			allParams := append(operation.Parameters, pathItem.Parameters...)
			for _, param := range allParams {
				// Skip body parameters - we'll handle them separately
				if param.In == "body" {
					continue
				}

				pc := ParameterCase{
					ParamName:   param.Name,
					ParamIn:     param.In,
					Required:    param.Required,
					Description: param.Description,
					EnumValues:  extractEnumFromSwaggerParam(param),
					DataType:    extractDataTypeFromSwaggerParam(param),
				}
				ec.Cases = append(ec.Cases, pc)
			}

			// 2. Extract request body (body parameters in Swagger 2.0)
			for _, param := range allParams {
				if param.In == "body" && param.Schema != nil {
					bodyCases := extractSwaggerRequestBodyCases(path, method, param.Schema, swagger)
					// If we successfully extracted individual properties, use them
					// Otherwise, fall back to treating it as a generic body parameter
					if len(bodyCases) > 0 {
						ec.Cases = append(ec.Cases, bodyCases...)
					} else {
						// Fallback: treat as generic body parameter
						pc := ParameterCase{
							ParamName:   param.Name,
							ParamIn:     param.In,
							Required:    param.Required,
							Description: param.Description,
							EnumValues:  nil,
							DataType:    "object",
						}
						ec.Cases = append(ec.Cases, pc)
					}
				}
			}

			results = append(results, ec)
		}
	}

	return results
}

// extractSwaggerRequestBodyCases extracts request body cases from Swagger 2.0 schema
func extractSwaggerRequestBodyCases(path, method string, schema *spec.Schema, swagger *spec.Swagger) []ParameterCase {
	out := []ParameterCase{}
	if schema == nil {
		return out
	}

	actualSchema := schema

	// If this is a reference, resolve it
	if schema.Ref.String() != "" {
		ref := schema.Ref.String()
		if strings.HasPrefix(ref, "#/definitions/") {
			definitionName := strings.TrimPrefix(ref, "#/definitions/")
			if def, ok := swagger.Definitions[definitionName]; ok {
				actualSchema = &def
			} else {
				// Reference not found, return empty
				return out
			}
		}
	}

	for name, propSchema := range actualSchema.Properties {
		pc := ParameterCase{
			ParamName:   name,
			ParamIn:     "body",
			Required:    contains(actualSchema.Required, name),
			EnumValues:  extractEnumFromSwaggerSchema(propSchema),
			Description: propSchema.Description,
			DataType:    extractDataTypeFromSwaggerSchema(propSchema),
		}
		out = append(out, pc)
	}

	return out
}

// Extract enum values from Swagger 2.0 parameter
func extractEnumFromSwaggerParam(param spec.Parameter) []interface{} {
	if param.Enum != nil {
		result := make([]interface{}, len(param.Enum))
		for i, v := range param.Enum {
			result[i] = v
		}
		return result
	}

	// Handle array items enum
	if param.Items != nil && param.Items.Enum != nil {
		result := make([]interface{}, len(param.Items.Enum))
		for i, v := range param.Items.Enum {
			result[i] = v
		}
		return result
	}

	return nil
}

// extractDataTypeFromSwaggerParam extracts data type from Swagger 2.0 parameter
func extractDataTypeFromSwaggerParam(param spec.Parameter) string {
	// Check if it's a body parameter with schema
	if param.In == "body" && param.Schema != nil {
		return extractDataTypeFromSwaggerSchema(*param.Schema)
	}

	// For other parameters, check the Type field
	if param.Type != "" {
		if param.Type == "array" && param.Items != nil {
			if param.Items.Type != "" {
				return "array[" + param.Items.Type + "]"
			}
			return "array"
		}
		return param.Type
	}

	return "string" // default fallback
}

// Extract enum values from Swagger 2.0 schema
func extractEnumFromSwaggerSchema(schema spec.Schema) []interface{} {
	if len(schema.Enum) > 0 {
		result := make([]interface{}, len(schema.Enum))
		for i, v := range schema.Enum {
			result[i] = v
		}
		return result
	}

	return nil
}

// extractDataTypeFromSwaggerSchema extracts data type from Swagger 2.0 schema
func extractDataTypeFromSwaggerSchema(schema spec.Schema) string {
	// Handle array types
	if len(schema.Type) > 0 && schema.Type[0] == "array" {
		if schema.Items != nil && schema.Items.Schema != nil {
			return "array[" + extractDataTypeFromSwaggerSchema(*schema.Items.Schema) + "]"
		}
		// For now, just return "array" without item type
		return "array"
	}

	// Handle object types
	if len(schema.Type) > 0 && schema.Type[0] == "object" {
		return "object"
	}

	// Return the primitive type
	if len(schema.Type) > 0 {
		return schema.Type[0]
	}

	return "string" // default fallback
}
