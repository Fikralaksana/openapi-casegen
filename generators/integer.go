package generators

// IntegerGenerator handles test case generation for integer parameters
type IntegerGenerator struct{}

// GenerateTestCases generates test cases for integer parameters
func (g *IntegerGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
	return []TestCase{
		{
			ID:          baseID + "_valid_input",
			Type:        "valid",
			Description: "Valid integer input",
		},
		{
			ID:          baseID + "_invalid_input",
			Type:        "invalid",
			Description: "Invalid input for integer parameter",
		},
		{
			ID:          baseID + "_boundary_min",
			Type:        "boundary_min",
			Description: "Minimum boundary value for integer",
		},
		{
			ID:          baseID + "_boundary_max",
			Type:        "boundary_max",
			Description: "Maximum boundary value for integer",
		},
	}
}
