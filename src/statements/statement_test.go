package statements

import (
	"appetit/values"
	"testing"
)

// Test the minver statement call.
func TestMinver(t *testing.T) {

	// Set a valid version number
	correct_version := 1

	// Run the MinVer() statement function with the sample tokens
	result_minver := MinVer(values.TEST_MINVER)

	// Compare the resuls
	if result_minver != correct_version {
		t.Errorf(
			"MinVer() did not return %d, returned %d",
			correct_version, result_minver,
		)
	}
}

// Test the set statement call.
func TestSet(t *testing.T) {

	correct_variable_value := "Hello World!"

	result_set := Set(values.TEST_SET)

	if result_set != correct_variable_value {
		t.Errorf(
			"Set did not return %s, returned %s",
			correct_variable_value, result_set,
		)
	}
}

// Test both the write and writeln statement calls.
func TestWrite(t *testing.T) {

	// Correct output for writeln calls
	correct_newline := "Hello World!\n"

	// Correct output for write calls
	correct_nonewline := "Hello World!"

	// Get the result of the Writeln() function for a writeln call
	result_newline := Writeln(values.TEST_WRITELN, true)

	// Compare the results
	if result_newline != correct_newline {
		t.Errorf(
			"Writeln did not return %s, returned %s",
			correct_newline, result_newline,
		)
	}

	// Get the result of the Writeln() function for a write call
	result_nonewline := Writeln(values.TEST_WRITE, false)

	// Compare the results
	if result_nonewline != correct_nonewline {
		t.Errorf(
			"Writeln did not return %s, returned %s",
			correct_newline, result_newline,
		)
	}
}