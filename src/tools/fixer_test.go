package tools

import (
	"testing"
)

// A simple test of the FixStringQuotations() function
func TestFixStringQuotationsSimple(t *testing.T) {
	test_string := "\"Hello World\""
	test_pass := "Hello World"

	result := FixStringQuotations(test_string)

	if result != test_pass {
		t.Errorf(
			"FixStringQuotations returned %q, got %q",
			test_pass,
			result,
		)
	}
}

// A more complex test of the FixStringQuotations() function
func TestFixStringQuotationsComplex(t *testing.T) {
	test_string := "\"Hello\" World\"\""
	test_pass := "Hello\" World\""

	result := FixStringQuotations(test_string)

	if result != test_pass {
		t.Errorf(
			"FixStringQuotations returned %q, got %q",
			test_pass,
			result,
		)
	}
}

// TODO: add tests for other functions in fixer