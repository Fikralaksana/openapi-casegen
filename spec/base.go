package processor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// SpecProcessor defines the interface for processing API specifications
type SpecProcessor interface {
	ProcessFile(filename string) ([]EndpointCases, error)
}

// EndpointCases represents a collection of test cases for an endpoint
type EndpointCases struct {
	Endpoint string
	Method   string
	Cases    []ParameterCase
}

// ParameterCase represents a single parameter with its test case information
type ParameterCase struct {
	ParamName   string
	ParamIn     string // path, query, header, cookie, body, formData
	Required    bool
	EnumValues  []interface{}
	Description string
	DataType    string
}

// DetectSpecVersion determines if a file is Swagger 2.0 or OpenAPI 3.0
func DetectSpecVersion(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Try to parse as JSON first
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err == nil {
		// Successfully parsed as JSON
		if _, ok := raw["swagger"]; ok {
			return "swagger2", nil
		}
		if _, ok := raw["openapi"]; ok {
			return "openapi3", nil
		}
	} else {
		// Not valid JSON, check if it's YAML with openapi
		content := string(data)
		if len(content) > 10 && (content[:7] == "openapi" || (len(content) > 10 && content[:10] == "openapi:")) {
			return "openapi3", nil
		}
	}

	return "", fmt.Errorf("unable to determine specification format")
}

// GetProcessor returns the appropriate processor for the spec version
func GetProcessor(version string) SpecProcessor {
	switch version {
	case "openapi3":
		return &OpenAPI3Processor{}
	case "swagger2":
		return &Swagger2Processor{}
	default:
		return nil
	}
}
