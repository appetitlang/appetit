/*
This module deals with the ask statement.
*/
package parser

import (
	"appetit/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Set a variable. Parameters include the tokens. Returns the final value of
the variable.
*/
func Ask(tokens []Token) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("ask")+" statement needs "+
				"to follow the form:\n\n\t"+utils.ColouriseCyan("ask")+" "+
				utils.ColouriseGreen("\"[question/prompt]\"")+
				utils.ColouriseMagenta(" to ")+
				utils.ColouriseYellow("\"[variable name]\"")+"\n\nAn example "+
				"of a working version check might be:\n\n\t"+
				utils.ColouriseCyan("ask")+" "+
				utils.ColouriseGreen("\"What is your name?\"")+
				utils.ColouriseMagenta(" to ")+
				utils.ColouriseGreen("\"name\"")+"\n\n"+
				"Your line of code looks like the following:\n\n\t"+
				utils.ColouriseRed(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Fix the prompt to ensure that quotation marks and escapes are handled
	properly.
	*/
	prompt := FixStringCombined(tokens[2].TokenValue)
	// Get a templated value for the prompt
	prompt = VariableTemplater(prompt)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	if action_error != nil {
		Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	/* Fix the variable name to ensure that quotation marks and escapes are
	handled properly.
	*/
	variable_name := FixStringCombined(tokens[4].TokenValue)

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
		variable_prefix = string(
			variable_name[0:len(SYMBOL_RESERVED_VARIABLE_PREFIX)],
		)
	}

	// Check the variable prefix
	var_prefix_error := CheckVariablePrefix(
		loc, variable_prefix, variable_name)
	if var_prefix_error != nil {
		ReportWithFixes(
			var_prefix_error.Error(),
			loc,
			tokens[4].TokenPosition,
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

	// If verbose mode is set
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s user \"%s\" and saving to variable %s...\n",
			utils.ColouriseBlue("Asking"),
			utils.ColouriseGreen(prompt),
			utils.ColouriseYellow(variable_name),
		)
	}

	/* Create a reader to get the input from user. Create a buffer size of
	65,536 bytes which doesn't seem to be acknowledged by any operating
	system. See issue #1 on GitHub for more. Leave this as-is though as
	it does allow for some extra space for input on platforms such as
	Windows. This is also potentially an issue with stdin limitations on
	any one given platform.
	*/
	input_reader := bufio.NewReaderSize(os.Stdin, 65536)
	// Prompt as per the prompt provided by the script
	fmt.Print(prompt)
	/* Read in the line while looking for the new line character as the
	delimiter
	*/
	//user_input, user_input_error := input_reader.ReadString('\n')
	user_input, user_input_error := input_reader.ReadString('\n')
	// Convert the user_input_bytes to a string
	user_input = strings.TrimSuffix(user_input, "\n")

	if user_input_error != nil {
		Report(
			"There was an error getting the user input. Please report the "+
				"following error in yellow to the project's GitHub repository "+
				"and a copy of the script:\n\n"+
				utils.ColouriseYellow(user_input_error.Error()),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Get the final variable value here by checking to see if the value is
	a math expression
	*/
	final_variable_value := CalculateValue(loc, user_input)

	// Set the variable
	VARIABLES[variable_name] = final_variable_value

	// Return the final value of the variable
	return final_variable_value
}
