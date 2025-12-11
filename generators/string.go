package generators

// StringGenerator handles test case generation for string parameters
type StringGenerator struct{}

// GenerateTestCases generates test cases for string parameters
func (g *StringGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
	return []TestCase{
		{
			ID:          baseID + "_valid_input",
			Type:        "valid",
			Description: "Valid string input",
		},
		{
			ID:          baseID + "_invalid_input",
			Type:        "invalid",
			Description: "Invalid input for string parameter",
		},
	}
}
