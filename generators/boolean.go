package generators

// BooleanGenerator handles test case generation for boolean parameters
type BooleanGenerator struct{}

// GenerateTestCases generates test cases for boolean parameters
func (g *BooleanGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
	return []TestCase{
		{
			ID:          baseID + "_valid_true",
			Type:        "valid",
			Description: "Valid boolean true value",
		},
		{
			ID:          baseID + "_valid_false",
			Type:        "valid",
			Description: "Valid boolean false value",
		},
		{
			ID:          baseID + "_invalid_input",
			Type:        "invalid",
			Description: "Invalid input for boolean parameter",
		},
	}
}
