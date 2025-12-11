package generators

// TestCase represents a generated test case
type TestCase struct {
	ID          string
	Type        string // valid, invalid, boundary_min, boundary_max, enum_value
	Description string
}

// Generator defines the interface for test case generators
type Generator interface {
	GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase
}

// GenerateTestCasesForType returns appropriate test cases based on the data type
func GenerateTestCasesForType(baseID string, dataType string, enumValues []interface{}) []TestCase {
	var generator Generator

	switch dataType {
	case "integer":
		generator = &IntegerGenerator{}
	case "number":
		generator = &NumberGenerator{}
	case "string":
		generator = &StringGenerator{}
	case "boolean":
		generator = &BooleanGenerator{}
	default:
		// Default to string generator for unknown types
		generator = &StringGenerator{}
	}

	// If enum values exist, use enum generator instead
	if len(enumValues) > 0 {
		generator = &EnumGenerator{}
	}

	return generator.GenerateTestCases(baseID, dataType, enumValues)
}
