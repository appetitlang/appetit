package tools

import (
	"os"
	"testing"
)

// A simple test of the FixStringQuotations() function
func TestFixStringQuotations(t *testing.T) {
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

	test_string_complex := "\"Hello\" World\"\""
	test_pass_complex := "Hello\" World\""

	result_complex := FixStringQuotations(test_string_complex)

	if result_complex != test_pass_complex {
		t.Errorf(
			"FixStringQuotations returned %q, got %q",
			test_pass_complex,
			result_complex,
		)
	}
}

func TestFixStringEscapes(t *testing.T){
	results := FixStringEscapes("\"Hello World\"")
	pass := "\"Hello World\""
	if results != pass {
		t.Errorf(
			"FixStringEscapes returned %q, got %q",
			pass,
			results,
		)
	}

	results_newline := FixStringEscapes("\"Hello\\n World\"")
	pass_newline := "\"Hello\n World\""
	if results_newline != pass_newline {
		t.Errorf(
			"FixStringEscapes returned %q, got %q",
			pass,
			results,
		)
	}

	results_charrret := FixStringEscapes("\"Hello\\r World\"")
	pass_charrret := "\"Hello\r World\""
	if results_charrret != pass_charrret {
		t.Errorf(
			"FixStringEscapes returned %q, got %q",
			pass,
			results,
		)
	}
}

func TestFixPathSeperators(t *testing.T) {
	results := FixPathSeperators("/home/user")
	pass := "/home/user" + string(os.PathSeparator)

	if results != pass {
		t.Errorf(
			"FixPathSeperators returned %q, got %q",
			pass,
			results,
		)
	}
}

// TODO: add tests for other functions in fixer