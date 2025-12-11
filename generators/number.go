package generators

// NumberGenerator handles test case generation for number (float) parameters
type NumberGenerator struct{}

// GenerateTestCases generates test cases for number parameters
func (g *NumberGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
	return []TestCase{
		{
			ID:          baseID + "_valid_input",
			Type:        "valid",
			Description: "Valid number input",
		},
		{
			ID:          baseID + "_invalid_input",
			Type:        "invalid",
			Description: "Invalid input for number parameter",
		},
		{
			ID:          baseID + "_boundary_min",
			Type:        "boundary_min",
			Description: "Minimum boundary value for number",
		},
		{
			ID:          baseID + "_boundary_max",
			Type:        "boundary_max",
			Description: "Maximum boundary value for number",
		},
	}
}
