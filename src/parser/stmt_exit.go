/*
This deals with the exit statement.
*/
package parser

import (
	"appetit/utils"
	"fmt"
	"os"
	"strconv"
)

/*
Handle an exit statement call. This one is very basic and doesn't require
much of the end user other than the statement call itself.
*/
func Exit(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 1)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("exit")+" statement needs "+
				"to follow the form:\n\n\t"+utils.ColouriseCyan("exit")+
				"\n\nThere are no values that you can or need to pass which is "+
				"most likely the cause here.\n\n"+
				"Your line of code looks like the following:\n\n\t"+
				utils.ColouriseRed(full_loc)+"\n\n",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if MODE_VERBOSE {
		fmt.Println(":: Exiting...")
	}
	// Finally, exit
	os.Exit(0)
}
