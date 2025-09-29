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
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("minver") + " statement needs to " +
			"follow the form " + tools.ColouriseCyan("minver") + " " +
			tools.ColouriseYellow("[version number]") + ". A common " +
			"issue here is excluding a version number or passing one as a " +
			"string (eg. " + tools.ColouriseGreen("\"3\"") + "). An " +
			"example of a working version check might be " +
			tools.ColouriseCyan("minver") + tools.ColouriseYellow(" 3") +
			" ( notice that there is an integer that is not in quotation " +
			"marks).",
			loc,
			tokens[1].TokenPosition,
			full_loc,
		)
	}

	/* Convert the loc to a simple number to keep things simple. We can discard
		the line number as this will always be an integer.
	*/
	loc_int, _ := strconv.Atoi(loc)
	/* Get the minimum version as a string. As we check whether this is an
		integer above, we should be good to go here.
	*/
	min_ver_string := tokens[2].TokenValue
	// Get the minver set by the user as an integer for comparison
	min_ver, err := strconv.Atoi(min_ver_string)
	/* If there is an error trying to do the conversion or if the min_ver is
		less than zero, report an error.
	*/
	if err != nil || min_ver < 0 {
		investigator.Report(
			"The version number " + tools.ColouriseYellow(min_ver_string) +
			" is not a valid version. You need to use a positive non-zero " +
			"integer.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Make sure that this is the first line. Note: this allows for shebang
		lines on *nix systems.
	*/
	if loc_int != 1 {
		// Here is the check for shebang systems.
		if loc_int == 2 && !values.SHEBANG_PRESENT {
			investigator.Report(
				"The " + tools.ColouriseCyan("minver") + " statement needs to " +
				"be the first line of the script. This helps to ensure that " +
				"the script is able to execute and doesn't fail part of the " +
				"way through. Move your " + tools.ColouriseCyan("minver") +
				" statement to the top of the script.",
				loc,
				"n/a",
				full_loc,
			)
		}
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
			"a newer version is available. " +
			"\n\nLine of Code: " + tools.ColouriseMagenta(full_loc),
			"n/a",
			"n/a",
			full_loc,
		)
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			":: Setting the %s required to run this script to %s\n",
			tools.ColouriseBlue("minver"),
			tools.ColouriseGreen(strconv.Itoa(min_ver)),
		)
	}

	// Return the minimum version required
	return min_ver
}