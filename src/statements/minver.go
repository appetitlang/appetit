/*
The statements package controls the execution of actual statements.

This module deals with the minimum version check called by for the minver
statement.
*/
package statements

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"strconv"
)

/*
	Check the minimum version required to run the script. Parameters include
	the tokens. Returns nothing.
*/
func MinVer(tokens []values.Token) int {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, token_number_err := investigator.ValidNumberOfTokens(tokens, 2)
	/* Get the minimum version as a string. As we check whether this is an
		integer above, we should be good to go here.
	*/
	min_ver_string := tokens[2].TokenValue
	// Get the minver set by the user as an integer for comparison
	min_ver, int_conversion_err := strconv.Atoi(min_ver_string)
	/* If there is an error trying to do the conversion or if the min_ver is
		less than zero, report an error. This also captures negative integers
		in particular as the negative sign and the integer are tokenised as
		seperate tokens.
	*/
	if int_conversion_err != nil || min_ver <= 0 || token_number_err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("minver") + " statement needs to " +
			"include a valid non-zero positive integer. A valid " +
			tools.ColouriseCyan("minver") + " statement needs to follow the " +
			"form:\n\t" + tools.ColouriseCyan("minver") +
			tools.ColouriseYellow(" [version number]") + "\nAn example of a " +
			"working version check might be:\n\t" +
			tools.ColouriseCyan("minver") + tools.ColouriseYellow(" 3") +
			"\nMake sure that you have none of the following for the " +
			tools.ColouriseCyan("minver") + " statement value:\n\t" +
			"- Negative number\n\t- Float (ie. decimal number)\n\t" +
			"- String\n\t- No value\n",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Check if the minimum version is greater than or equal to the language
		version.
	*/
	if min_ver > values.LANG_VERSION {
		investigator.Report(
			"The script you're running here requires a newer version of " +
			"the interpreter. You are running version " +
			tools.ColouriseYellow(strconv.Itoa(values.LANG_VERSION)) + 
			" but the script requires at least version " +
			tools.ColouriseYellow(min_ver_string) + ". Check to see if " +
			"a newer version is available. ",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			":: Setting the minimum version of Appetit (%s) required to " +
			"run this script to %s\n",
			tools.ColouriseBlue("minver"),
			tools.ColouriseGreen(strconv.Itoa(min_ver)),
		)
	}

	// Return the minimum version required
	return min_ver
}