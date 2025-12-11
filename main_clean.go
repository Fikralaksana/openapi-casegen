package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"openapi-tester/generators"
	"openapi-tester/processor"
)

// --------------------------
// Processing logic is now in the processor package
// --------------------------

// generateTestCaseIDs creates test case identifiers for a parameter
func generateTestCaseIDs(endpoint, method string, param processor.ParameterCase) []string {
	// Clean endpoint path for use in test ID (remove leading slash, replace slashes with dots)
	endpointClean := strings.TrimPrefix(endpoint, "/")
	endpointClean = strings.ReplaceAll(endpointClean, "/", ".")
	endpointClean = strings.ReplaceAll(endpointClean, "{", "")
	endpointClean = strings.ReplaceAll(endpointClean, "}", "")

	// Base test ID: endpoint.paramname
	baseID := fmt.Sprintf("%s.%s", endpointClean, param.ParamName)

	// Use the generators package to create test cases
	testCases := generators.GenerateTestCasesForType(baseID, param.DataType, param.EnumValues)

	// Extract just the IDs from the test cases
	var testIDs []string
	for _, tc := range testCases {
		testIDs = append(testIDs, tc.ID)
	}

	return testIDs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: openapi-casegen <openapi-spec-file>")
		fmt.Println("Supports OpenAPI 3.0 and Swagger 2.0 specifications")
		fmt.Println("Example: openapi-casegen openapi.yaml")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Detect specification format and get appropriate processor
	version, err := processor.DetectSpecVersion(filename)
	if err != nil {
		log.Fatalf("failed to detect specification format: %v", err)
	}

	processor := processor.GetProcessor(version)
	if processor == nil {
		log.Fatalf("unsupported specification format: %s", version)
	}

	endpoints, err := processor.ProcessFile(filename)
	if err != nil {
		log.Fatalf("failed to process specification: %v", err)
	}

	fmt.Println("===== Generated Test Case IDs =====")
	for _, ep := range endpoints {
		fmt.Printf("\n[%s] %s\n", ep.Method, ep.Endpoint)
		for _, c := range ep.Cases {
			testCaseIDs := generateTestCaseIDs(ep.Endpoint, ep.Method, c)
			for _, testID := range testCaseIDs {
				fmt.Printf("- %s\n", testID)
			}
		}
	}
}
