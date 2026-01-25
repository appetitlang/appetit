/*
This module deals with the write and writeln statements.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"strconv"
)

/*
Write output to its own line or on one line. This handles both the write
and writeln statement. Parameters include the tokens, and newline as a bool
for whether output needs to add a new line (writeln) or leave the line
without a newline character at the end. Returns nothing.
*/
func Writeln(tokens []values.Token, newline bool) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("write/writeln")+" statement "+
				"needs to follow the form "+
				tools.ColouriseCyan("write/writeln")+" "+
				tools.ColouriseYellow("[content to be written]")+". A "+
				"common error here is trying to concatenate multiple values "+
				"into one statement call here. An example of a working version "+
				"might be "+tools.ColouriseCyan("write/writeln")+
				tools.ColouriseGreen("\"Hello World\"")+"\n\nLine of Code: "+
				tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the string to be printed
	trimmed_output := tools.FixStringCombined(tokens[2].TokenValue)
	// Replace any variables in the output string
	trimmed_output = VariableTemplater(trimmed_output)

	/* If newline is true, we are parsing a writeln, otherwise, we are parsing
	a write
	*/
	if newline {
		// Print out the output with a newline as we are parsing a writeln
		fmt.Printf("%s\n", trimmed_output)
		return fmt.Sprintf("%s\n", trimmed_output)
	} else {
		// Print out the output with a newline as we are parsing a write
		fmt.Print(trimmed_output)
		return trimmed_output
	}
}
