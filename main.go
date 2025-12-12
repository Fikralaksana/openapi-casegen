package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"openapi-tester/generators"
	"openapi-tester/spec"
	"openapi-tester/validator"
)

// --------------------------
// Processing logic is now in the processor package
// --------------------------

// generateEndpointAccessTestID creates a basic test case ID for accessing an endpoint
func generateEndpointAccessTestID(endpoint, method string) string {
	// Clean endpoint path for use in test ID (remove leading slash, replace slashes with underscores)
	endpointClean := strings.TrimPrefix(endpoint, "/")
	endpointClean = strings.ReplaceAll(endpointClean, "/", "_")
	endpointClean = strings.ReplaceAll(endpointClean, "{", "")
	endpointClean = strings.ReplaceAll(endpointClean, "}", "")

	return fmt.Sprintf("%s_%s_basic_access", endpointClean, strings.ToLower(method))
}

// generateTestCaseIDs creates test case identifiers for a parameter
func generateTestCaseIDs(endpoint, method string, param processor.ParameterCase) []string {
	// Clean endpoint path for use in test ID (remove leading slash, replace slashes with underscores)
	endpointClean := strings.TrimPrefix(endpoint, "/")
	endpointClean = strings.ReplaceAll(endpointClean, "/", "_")
	endpointClean = strings.ReplaceAll(endpointClean, "{", "")
	endpointClean = strings.ReplaceAll(endpointClean, "}", "")

	// Base test ID: endpoint_paramname
	baseID := fmt.Sprintf("%s_%s", endpointClean, param.ParamName)

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
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage:")
		fmt.Println("  openapi-casegen <openapi-spec-file>                    # Generate test cases")
		fmt.Println("  openapi-casegen <openapi-spec-file> <junit-xml-file>   # Validate tests against JUnit XML")
		fmt.Println("")
		fmt.Println("Supports OpenAPI 3.0 and Swagger 2.0 specifications")
		fmt.Println("Examples:")
		fmt.Println("  openapi-casegen openapi.yaml")
		fmt.Println("  openapi-casegen openapi.yaml results.xml")
		os.Exit(1)
	}

	specFile := os.Args[1]

	// Detect specification format and get appropriate processor
	version, err := processor.DetectSpecVersion(specFile)
	if err != nil {
		log.Fatalf("failed to detect specification format: %v", err)
	}

	processor := processor.GetProcessor(version)
	if processor == nil {
		log.Fatalf("unsupported specification format: %s", version)
	}

	endpoints, err := processor.ProcessFile(specFile)
	if err != nil {
		log.Fatalf("failed to process specification: %v", err)
	}

	// Collect all generated test case IDs
	var generatedTestIDs []string
	for _, ep := range endpoints {
		// Generate basic endpoint access test case
		endpointTestID := generateEndpointAccessTestID(ep.Endpoint, ep.Method)
		generatedTestIDs = append(generatedTestIDs, endpointTestID)

		// Generate parameter-specific test cases
		for _, c := range ep.Cases {
			testCaseIDs := generateTestCaseIDs(ep.Endpoint, ep.Method, c)
			generatedTestIDs = append(generatedTestIDs, testCaseIDs...)
		}
	}

	// Check if validation mode is requested
	if len(os.Args) == 3 {
		xmlFile := os.Args[2]
		validateTests(generatedTestIDs, xmlFile)
	} else {
		printGeneratedTests(endpoints)
	}
}

func printGeneratedTests(endpoints []processor.EndpointCases) {
	fmt.Println("===== Generated Test Case IDs =====")
	for _, ep := range endpoints {
		fmt.Printf("\n[%s] %s\n", ep.Method, ep.Endpoint)

		// Generate basic endpoint access test case
		endpointTestID := generateEndpointAccessTestID(ep.Endpoint, ep.Method)
		fmt.Printf("- %s\n", endpointTestID)

		// Generate parameter-specific test cases
		for _, c := range ep.Cases {
			testCaseIDs := generateTestCaseIDs(ep.Endpoint, ep.Method, c)
			for _, testID := range testCaseIDs {
				fmt.Printf("- %s\n", testID)
			}
		}
	}
}

func validateTests(generatedTestIDs []string, xmlFile string) {
	v := validator.NewValidator()

	// Load test results from XML
	actualTests, err := v.LoadTestResults(xmlFile)
	if err != nil {
		log.Fatalf("failed to load test results: %v", err)
	}

	// Compare generated tests with actual tests
	result := v.CompareTests(generatedTestIDs, actualTests)

	// Print validation report
	v.PrintReport(result)
}
