/*
The checks module provides checks for various parts of a script.
*/
package investigator

import (
	"appetit/tools"
	"appetit/values"
	"errors"
	"fmt"
	"os"
	"slices"
)

/*
Check that an action is signaled by an appropriate
values.SYMBOL_ACTION. Parameters include loc, the line of the script
(if an error needs to be reported and the value to check. Returns
error if the action is incorrect.
*/
func CheckAction(loc string, action_value string) error {
	if action_value != values.SYMBOL_ACTION {
		return fmt.Errorf(
			"an action was made using an invalid action statement (%s), "+
				"please ensure that you use %s",
			tools.ColouriseMagenta(action_value),
			tools.ColouriseMagenta(values.SYMBOL_ACTION),
		)
	}

	return nil
}

/*
Check that something is a valid statement. Parameters include loc, the line
of the script (if an error needs to be reported) and the value to check.
Returns bool, true if value_name is a valid statement name.
*/
func CheckIsStatement(loc string, value_name string) bool {
	return slices.Contains(values.STATEMENT_NAMES, value_name)
}

/*
Check that the prefix of the variable isn't the reserved one. Parameters
include loc, the line of the script (if an error needs to be reported and
the value to check. Returns bool, true if the variable does not start
with the reserved variable prefix.
*/
func CheckVariablePrefix(
	loc string, prefix string, variable_name string) error {
	// If the prefix is the reserved variable prefix...
	if prefix != values.RESERVED_VARIABLE_PREFIX {
		return nil
		// Report back an error while providing the operator to use
	}

	return fmt.Errorf(
		"you've named a variable %s which starts with %s (this is not "+
			"allowed). If you were trying to use a reserved variable, "+
			"consult the following list: %s",
		tools.ColouriseYellow(variable_name),
		tools.ColouriseYellow(values.RESERVED_VARIABLE_PREFIX),
		values.ListReservedVariables(),
	)
}

/*
Check that an assignment operator is a valid assignment operator.
Parameters include loc, the line of the script (if an error needs to be
reported and the value to check. Returns bool, true if the assignment
operator is valid.
*/
func CheckValidAssignment(loc string, value_name string) error {
	// If the value_name is not the ASSIGNMENT_OPERATOR...
	if value_name != values.OPERATOR_ASSIGNMENT {
		// Report back an error while providing the operator to use
		return fmt.Errorf(
			"an assignment was made using an invalid operator (%s), please ensure that you use %s",
			value_name,
			tools.ColouriseMagenta(values.OPERATOR_ASSIGNMENT),
		)
	}
	return nil
}

/*
Report whether there is an appropriate number of tokens. Parameters include
the tokens and the required_number which is the required number of tokens
to have in the line. Returns a bool, true if there is a valid number of
tokens and false otherwise. Additionally, an error is reported to make the
gopher happy.
*/
func ValidNumberOfTokens(
	tokens []values.Token, required_number int) (bool, error) {
	/* Get the token count and subtracting one to account for the fact that the
	line number is included.
	*/
	token_count := len(tokens) - 1

	// If the token_count does not equal the required number of tokens
	if token_count != required_number {
		// Return false an an error message
		return false, fmt.Errorf("invalid number of tokens")
	}

	// If we got here, we can
	return true, nil
}

/*
Check to ensure that a file exists. Parameters include the file name
itself. Returns a boolean, true if the file exists, false if it does not.
*/
func FileExists(file_name string) bool {
	// Check whether the file exists
	_, err := os.Stat(file_name)
	/* Return whether it does or does not exist. Here, we're returning true if
	there is not an error (and thus the file is presumed to exist) and false if
	there is an error (and thus the file is presumed not to exist).
	*/
	return !errors.Is(err, os.ErrNotExist)
}
