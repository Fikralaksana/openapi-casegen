package processor

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// OpenAPI3Processor handles OpenAPI 3.0 specifications
type OpenAPI3Processor struct{}

// ProcessFile loads and processes an OpenAPI 3.0 specification file
func (p *OpenAPI3Processor) ProcessFile(filename string) ([]EndpointCases, error) {
	doc, err := loadOpenAPI3Spec(filename)
	if err != nil {
		return nil, err
	}

	return extractEndpointsOpenAPI3(doc), nil
}

// loadOpenAPI3Spec loads an OpenAPI 3.0 specification
func loadOpenAPI3Spec(path string) (*openapi3.T, error) {
	loader := &openapi3.Loader{
		IsExternalRefsAllowed: true,
	}
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, err
	}
	err = doc.Validate(loader.Context)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// extractEndpointsOpenAPI3 extracts endpoints from OpenAPI 3.0 specification
func extractEndpointsOpenAPI3(doc *openapi3.T) []EndpointCases {
	results := []EndpointCases{}

	for path, pathItem := range doc.Paths {
		// Get all operations for this path
		operations := pathItem.Operations()

		for method, operation := range operations {
			ec := EndpointCases{
				Endpoint: path,
				Method:   method,
				Cases:    []ParameterCase{},
			}

			// 1. Extract parameters (query, path, header, cookie)
			for _, paramRef := range operation.Parameters {
				if paramRef.Value == nil {
					continue
				}
				p := paramRef.Value

				pc := ParameterCase{
					ParamName:   p.Name,
					ParamIn:     p.In,
					Required:    p.Required,
					Description: p.Description,
					EnumValues:  enumFromSchema(p.Schema),
					DataType:    extractDataTypeFromOpenAPI3Schema(p.Schema),
				}

				ec.Cases = append(ec.Cases, pc)
			}

			// 2. Extract request body (JSON schema)
			if operation.RequestBody != nil &&
				operation.RequestBody.Value != nil &&
				operation.RequestBody.Value.Content["application/json"] != nil {

				bodySchema := operation.RequestBody.Value.Content["application/json"].Schema
				if bodySchema != nil {
					bodyCases := extractRequestBodyCasesOpenAPI3(path, method, bodySchema)
					ec.Cases = append(ec.Cases, bodyCases...)
				}
			}

			results = append(results, ec)
		}
	}

	return results
}

// extractRequestBodyCasesOpenAPI3 extracts request body cases from OpenAPI 3.0 schema
func extractRequestBodyCasesOpenAPI3(path, method string, schemaRef *openapi3.SchemaRef) []ParameterCase {
	out := []ParameterCase{}
	if schemaRef.Value == nil {
		return out
	}

	for name, s := range schemaRef.Value.Properties {
		pc := ParameterCase{
			ParamName:   name,
			ParamIn:     "body",
			Required:    contains(schemaRef.Value.Required, name),
			EnumValues:  s.Value.Enum,
			Description: s.Value.Description,
			DataType:    extractDataTypeFromOpenAPI3Schema(s),
		}
		out = append(out, pc)
	}

	return out
}

// extractDataTypeFromOpenAPI3Schema extracts the data type from OpenAPI 3.0 schema
func extractDataTypeFromOpenAPI3Schema(schemaRef *openapi3.SchemaRef) string {
	if schemaRef == nil || schemaRef.Value == nil {
		return "string" // default to string if no schema
	}

	schema := schemaRef.Value

	// Handle array types
	if schema.Type == "array" {
		if schema.Items != nil && schema.Items.Value != nil {
			return "array[" + extractDataTypeFromOpenAPI3Schema(schema.Items) + "]"
		}
		return "array"
	}

	// Handle object types
	if schema.Type == "object" {
		return "object"
	}

	// Return the primitive type
	if schema.Type != "" {
		return schema.Type
	}

	return "string" // default fallback
}

// safely extract enum values
func enumFromSchema(ref *openapi3.SchemaRef) []interface{} {
	if ref == nil || ref.Value == nil {
		return nil
	}
	return ref.Value.Enum
}

func contains(arr []string, k string) bool {
	for _, x := range arr {
		if x == k {
			return true
		}
	}
	return false
}
