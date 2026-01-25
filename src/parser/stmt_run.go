/*
This module deals with the run statement.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Run(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("run")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("run")+
				tools.ColouriseGreen(" \"[script]\"")+". An example of a "+
				"working version check might be "+tools.ColouriseCyan("run")+
				tools.ColouriseGreen("\"other_script.apt\"")+"\n\n"+
				"Line of Code: "+tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Set the path name for the script to be run, fixing any issues with the
	string.
	*/
	script_name := tools.FixStringCombined(tokens[2].TokenValue)
	// Replace any variables in the output string
	script_name = VariableTemplater(script_name)

	_, exist_error := os.Stat(script_name)
	if errors.Is(exist_error, os.ErrNotExist) {
		investigator.Report(
			"The script - "+tools.ColouriseYellow(script_name)+" - does "+
				"not exist and/or can't be accessed. Double check to verify "+
				"that the script exists.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	contents := PrepScript(script_name)

	if values.MODE_DEV {
		// Start printing out the tokens
		fmt.Println(tools.ColouriseYellow("\nTokens"))
		Start(contents, true)
	} else {
		Start(contents, false)
	}
}
