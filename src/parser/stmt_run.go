/*
This module deals with the run statement.
*/
package parser

import (
	"appetit/utils"
	"fmt"
	"strconv"
)

func Run(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("run")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("run")+
				utils.ColouriseGreen(" \"[script]\"")+". An example of a "+
				"working version check might be "+utils.ColouriseCyan("run")+
				utils.ColouriseGreen("\"other_script.apt\"")+"\n\n"+
				"Line of Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Set the path name for the script to be run, fixing any issues with the
	string.
	*/
	script_name := FixStringCombined(tokens[2].TokenValue)
	// Replace any variables in the output string
	script_name = VariableTemplater(script_name)

	file_exists := CheckFileExists(script_name)

	if !file_exists {
		Report(
			"The script - "+utils.ColouriseYellow(script_name)+" - does "+
				"not exist and/or can't be accessed. Double check to verify "+
				"that the script exists.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	contents := PrepScript(script_name)

	if MODE_DEV {
		// Start printing out the tokens
		fmt.Println(utils.ColouriseYellow("\nTokens"))
		Start(contents, true)
	} else {
		Start(contents, false)
	}
}
