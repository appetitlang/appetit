package investigator

import (
	"appetit/values"
	"testing"
)

// Test the CheckAction() function
func TestCheckAction(t *testing.T) {

	value_name := values.SYMBOL_ACTION

	result := CheckAction("1", value_name)

	if result != nil {
		t.Errorf(
			"CheckAction did not return true, invalid action symbol passed",
		)
	}
}

// Test the CheckIsStatement() function
func TestCheckIsStatement(t *testing.T) {
	// Get a valid statement name for testing
	valid_statement := values.STATEMENT_NAMES[0]
	// Set a random statement name that is invalid
	invalid_statement := "RANDOM!"

	// Get a response to the CheckIsStatement that should return true
	valid_result := CheckIsStatement("1", valid_statement)
	// Get a response to the CheckIsStatement that should return false
	invalid_result := CheckIsStatement("1", invalid_statement)

	// If a valid statement name returns false...
	if !valid_result {
		t.Errorf(
			"CheckIsStatement returned false when it should have returned " +
			"true",
		)
	}
	// If an invalid statement name returns true...
	if invalid_result {
		t.Errorf(
			"CheckIsStatement returned true when it should have returned " +
			"false",
		)
	}
}

// Test the CheckVariablePrefix() function
func TestCheckVariablePrefix(t *testing.T) {

	prefix := "x_"
	value_name := "Hello"

	result := CheckVariablePrefix("1", prefix, value_name)

	if result != nil{
		t.Errorf(
			"CheckVariablePrefix did not return true, passed %q",
			value_name,
		)
	}
}

// Test the CheckValidAssignment() function
func TestCheckValidAssignment(t *testing.T) {

	result := CheckValidAssignment("1", values.OPERATOR_ASSIGNMENT)

	if result != nil {
		t.Errorf(
			"CheckValidAssignment did not return true, invalid assignment " +
			"operator passed",
		)
	}	
}