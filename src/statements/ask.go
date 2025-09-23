package statements

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
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
func Ask(tokens []values.Token) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("ask") + " statement needs " +
			"to follow the form " + tools.ColouriseCyan("ask") + " " +
			tools.ColouriseGreen("\"[question/prompt]\"") + " to " +
			tools.ColouriseYellow("\"[variable name]\"") + ". An example of " +
			"a working version check might be " + tools.ColouriseCyan("ask") +
			tools.ColouriseGreen("\"What is your name?\"") + " to " +
			tools.ColouriseGreen("\"name\"") + "\n\n" +
			"Line of Code: " + tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Fix the prompt to ensure that quotation marks and escapes are handled
		properly.
	*/
	prompt := tools.FixStringCombined(tokens[2].TokenValue)
	// Get a templated value for the prompt
	prompt = VariableTemplater(prompt)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := investigator.CheckAction(loc, action)
	if action_error != nil {
		investigator.Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}
	
	/* Fix the variable name to ensure that quotation marks and escapes are
		handled properly.
	*/
	variable_name := tools.FixStringCombined(tokens[4].TokenValue)
	
	/* Get the prefix of the variable so that we can check that it isn't
		reserved
	*/
	// Hold the (possible) prefix for checking
	var variable_prefix string
	// If the length of the variable is less than the RESERVED_VARIABLE_PREFIX
	if len(variable_name) < len(values.RESERVED_VARIABLE_PREFIX) {
		// Just set the prefix to the variable
		variable_prefix = string(variable_name)
	} else {
		// Otherwise, create a prefix to check against
		variable_prefix = string(
			variable_name[0:len(values.RESERVED_VARIABLE_PREFIX)],
		)
	}

	// Check the variable prefix
	var_prefix_error := investigator.CheckVariablePrefix(
		loc, variable_prefix, variable_name)
	if var_prefix_error != nil {
		investigator.ReportWithFixes(
			var_prefix_error.Error(),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Check that the variable name is not one of the statement names
	statement := investigator.CheckIsStatement(loc, variable_name)
	// If it is a statement
	if statement {
		investigator.ReportWithFixes(
			"The variable - " + tools.ColouriseYellow(variable_name) + " - " +
			"is not a valid variable name as it conflicts with a statement " +
			"name.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s user \"%s\" and saving to variable %s...\n",
			tools.ColouriseBlue("Asking"),
			tools.ColouriseGreen(prompt),
			tools.ColouriseYellow(variable_name),
		)
	}

	input_reader := bufio.NewReader(os.Stdin)
	// Prompt as per the prompt provided by the script
	fmt.Print(prompt)
	user_input, user_input_error := input_reader.ReadString('\n')
	user_input = strings.TrimSuffix(user_input, "\n")

	if user_input_error != nil {
		investigator.Report(
			"There was an error getting the user input. " +
			user_input_error.Error(),
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
	values.VARIABLES[variable_name] = final_variable_value
	// Return the final value of the variable
	return final_variable_value
}