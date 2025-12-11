package generators

// EnumGenerator handles test case generation for enum parameters
type EnumGenerator struct{}

// GenerateTestCases generates test cases for enum parameters
func (g *EnumGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
	var testCases []TestCase

	// Generate test case for each enum value
	for _, enumVal := range enumValues {
		testCases = append(testCases, TestCase{
			ID:          baseID + "_valid_" + enumVal.(string),
			Type:        "enum_value",
			Description: "Valid enum value: " + enumVal.(string),
		})
	}

	// Add invalid input test case
	testCases = append(testCases, TestCase{
		ID:          baseID + "_invalid_input",
		Type:        "invalid",
		Description: "Invalid input for enum parameter",
	})

	return testCases
}
