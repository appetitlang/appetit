/*
This module deals with the pause statement.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"strconv"
	"time"
)

/*
Pause the execution of the script. Parameters include the tokens. Returns
nothing.
*/
func Pause(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("pause")+" statement needs to "+
				"follow the form "+tools.ColouriseCyan("pause")+" "+
				tools.ColouriseYellow("[version number]")+". A common "+
				"issue here is excluding a version number or passing one as a "+
				"string (eg. "+tools.ColouriseGreen("\"3\"")+"). An "+
				"example of a working version check might be "+
				tools.ColouriseCyan("pause")+tools.ColouriseYellow(" 3"),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Convert the loc to a simple number to keep things simple. We can discard
	the line number as this will always be an integer.
	*/
	pause_as_string := tokens[2].TokenValue

	// Create an integer version of the pause length
	pause_int, err := strconv.Atoi(pause_as_string)
	/* If there is an error trying to do the conversion or if the pause_int is
	less than zero, report an error.
	*/
	if err != nil || pause_int < 0 {
		investigator.Report(
			"The version number "+tools.ColouriseYellow(pause_as_string)+
				" is not a valid version. You need to use a positive non-zero "+
				"integer.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set...
	if values.MODE_VERBOSE {
		fmt.Printf(":: Pausing for %s seconds...", pause_as_string)
	}
	// Pause execution by sleeping for the required number of seconds
	time.Sleep(time.Duration(pause_int) * time.Second)
}
