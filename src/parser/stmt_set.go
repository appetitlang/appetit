/*
This module deals with the set statement and any variable related
functionality. In this way, it does more heavy lifting relative to other
statement modules which often just have single functions in them.
*/
package parser

import (
	"appetit/utils"
	"fmt"
	"go/token"
	"go/types"
	"strconv"
	"strings"
)

/*
Calculate the value variable to see if it's a math expression. Parameters
include loc, the line of the script (if an error needs to be reported and
the value to calculate. Returns the value of the calculated string.
*/
func CalculateValue(loc string, value string) string {
	// Thanks to https://stackoverflow.com/a/65484336 for the below
	// Get a file set of tokens
	expression_tokens := token.NewFileSet()
	// Get the tokens and evaluate them
	tv, err := types.Eval(expression_tokens, nil, token.NoPos, value)
	// If there's an error trying to evaluate the tokens, return the value
	if err != nil {
		return value
	}
	// Return the calculated value as a string. This has
	return tv.Value.ExactString()
}

/*
Replace variables inside of a string. Parameters include the input line of
code to fix. Returns a templated string where variables have been fixed.
*/
func VariableTemplater(input string) string {
	// For each key-value pair in the map of variables
	for key, value := range VARIABLES {
		// Get the string value of the variable
		value = string(value)
		/* Replace the value in the string if the value is found in the
		string prepended by the variable replacement symbol
		*/
		input = strings.ReplaceAll(
			input,
			SYMBOL_VARIABLE_SUBSTITUTION+key,
			value,
		)
	}
	// Return the substituted string
	return input
}

/*
Set a variable. Parameters include the tokens. Returns the variable value.
*/
func Set(tokens []Token) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("set")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("set")+" "+
				utils.ColouriseYellow("[variable name]")+" = "+
				utils.ColouriseGreen("\"[value]\"")+". An example of a "+
				"working version check might be "+utils.ColouriseCyan("set")+
				" name = "+
				utils.ColouriseGreen("\""+LANG_NAME+"\"")+"\n\n"+
				"Line of Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Set the variable name
	variable_name := tokens[2].TokenValue
	// The assignment operator
	assignment_operator := tokens[3].TokenValue
	// The variable value with fixes that need to be made
	variable_value := FixStringCombined(tokens[4].TokenValue)

	/* Get the prefix of the variable so that we can check that it isn't
	reserved
	*/
	// Hold the (possible) prefix for checking
	var variable_prefix string
	// If the length of the variable is less than the RESERVED_VARIABLE_PREFIX
	if len(variable_name) < len(SYMBOL_RESERVED_VARIABLE_PREFIX) {
		// Just set the prefix to the variable
		variable_prefix = string(variable_name)
	} else {
		// Otherwise, create a prefix to check against
		variable_prefix = string(variable_name[0:2])
	}

	// Check the variable prefix
	var_prefix_error := CheckVariablePrefix(
		loc, variable_prefix, variable_name)
	if var_prefix_error != nil {
		ReportWithFixes(
			var_prefix_error.Error(),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Check that the variable name is not one of the statement names
	statement := CheckIsStatement(loc, variable_name)
	// If it is a statement
	if statement {
		ReportWithFixes(
			"The variable - "+utils.ColouriseYellow(variable_name)+" - "+
				"is not a valid variable name as it conflicts with a statement "+
				"name.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Check for a valid assignment operator
	assignment_error := CheckValidAssignment(loc, assignment_operator)
	if assignment_error != nil {
		ReportWithFixes(
			assignment_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	templated_variable := VariableTemplater(variable_value)

	/* Get the final variable value here by checking to see if the value is
	a math expression
	*/
	final_variable_value := CalculateValue(loc, templated_variable)

	// If verbose mode is set...
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Setting"),
			utils.ColouriseYellow(variable_name),
			utils.ColouriseGreen(final_variable_value),
		)
	}
	// Set the variable
	VARIABLES[variable_name] = final_variable_value

	// If verbose mode is set, report that things are done.
	if MODE_VERBOSE {
		fmt.Println("done!")
	}

	return final_variable_value

}
