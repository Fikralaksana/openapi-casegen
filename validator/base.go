package validator

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// TestResult represents a single test case from JUnit XML
type TestResult struct {
	Name      string `xml:"name,attr"`
	ClassName string `xml:"classname,attr"`
	Time      string `xml:"time,attr"`
	Status    string // derived from presence of failure/error
}

// TestSuite represents a test suite from JUnit XML
type TestSuite struct {
	Name      string       `xml:"name,attr"`
	Tests     int          `xml:"tests,attr"`
	Failures  int          `xml:"failures,attr"`
	Errors    int          `xml:"errors,attr"`
	Time      string       `xml:"time,attr"`
	TestCases []TestResult `xml:"testcase"`
}

// TestSuites represents the root element of JUnit XML
type TestSuites struct {
	TestSuites []TestSuite `xml:"testsuite"`
}

// ValidationResult represents the comparison between generated and actual tests
type ValidationResult struct {
	Implemented []string // Tests that exist in both generated and actual
	Missing     []string // Tests that are generated but not implemented
	Extra       []string // Tests that exist in actual but weren't generated
}

// Validator provides test validation functionality
type Validator struct{}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// LoadTestResults loads and parses JUnit XML test results
func (v *Validator) LoadTestResults(filename string) ([]TestResult, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read test results file: %v", err)
	}

	var testSuites TestSuites
	err = xml.Unmarshal(data, &testSuites)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %v", err)
	}

	var results []TestResult
	for _, suite := range testSuites.TestSuites {
		for _, testCase := range suite.TestCases {
			// Determine status based on whether there are failure/error elements
			testCase.Status = "passed" // Assume passed unless we find failure/error
			results = append(results, testCase)
		}
	}

	return results, nil
}

// CompareTests compares generated test case IDs with actual test results
func (v *Validator) CompareTests(generatedTestIDs []string, actualTests []TestResult) *ValidationResult {
	result := &ValidationResult{}

	// Create a map of actual test names for quick lookup
	actualTestMap := make(map[string]bool)
	for _, test := range actualTests {
		actualTestMap[test.Name] = true
	}

	// Check each generated test
	for _, generatedID := range generatedTestIDs {
		if actualTestMap[generatedID] {
			result.Implemented = append(result.Implemented, generatedID)
		} else {
			result.Missing = append(result.Missing, generatedID)
		}
	}

	// Find extra tests (tests that exist but weren't generated)
	for _, test := range actualTests {
		found := false
		for _, generatedID := range generatedTestIDs {
			if test.Name == generatedID {
				found = true
				break
			}
		}
		if !found {
			result.Extra = append(result.Extra, test.Name)
		}
	}

	return result
}

// PrintReport prints a formatted validation report
func (v *Validator) PrintReport(result *ValidationResult) {
	fmt.Println("===== Test Validation Report =====")
	fmt.Printf("✅ Implemented: %d tests\n", len(result.Implemented))
	fmt.Printf("❌ Missing: %d tests\n", len(result.Missing))
	fmt.Printf("➕ Extra: %d tests\n", len(result.Extra))

	if len(result.Implemented) > 0 {
		fmt.Println("\n✅ IMPLEMENTED TESTS:")
		for _, test := range result.Implemented {
			fmt.Printf("  - %s\n", test)
		}
	}

	if len(result.Missing) > 0 {
		fmt.Println("\n❌ MISSING TESTS:")
		for _, test := range result.Missing {
			fmt.Printf("  - %s\n", test)
		}
	}

	if len(result.Extra) > 0 {
		fmt.Println("\n➕ EXTRA TESTS (not in OpenAPI spec):")
		for _, test := range result.Extra {
			fmt.Printf("  - %s\n", test)
		}
	}

	totalGenerated := len(result.Implemented) + len(result.Missing)
	if totalGenerated > 0 {
		coverage := float64(len(result.Implemented)) / float64(totalGenerated) * 100
		fmt.Printf("\n📊 Coverage: %.1f%% (%d/%d generated tests implemented)\n",
			coverage, len(result.Implemented), totalGenerated)
	}
}