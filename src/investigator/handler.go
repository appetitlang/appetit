/*
The handler module provides error handling and reporting for language specific
issues.
*/
package investigator

import (
	"appetit/tools"
	"fmt"
	"os"
	"strings"
)

/*
	Report an error. Parameters include error_message, the error message
	itself, the line_number, the line number that triggered the error, the
	token position, the place where the error occured on the line, and the full
	line of code as a string. Returns nothing.
*/
func Report(
	error_message string,
	line_number string, token_pos string, full_loc string) {
	/*
		For context, an error looks similar to the following:

		[Error on line X] or [Error]
		This is the error
		[End of Error]
	*/

	// This holds the header
	error_msg := ""

	// If the line_number is N/A (ie. the error isn't specific to a line)
	if line_number == "n/a" {
		// Create a general error
		error_msg = tools.ColouriseRed("\n[Error]\n")
	} else if token_pos == "n/a" {
		error_msg = tools.ColouriseRed(
			"\n\n[Error on line " + line_number + "]\n",
		)
	} else {
		// Otherwise, create a header with the line number
		error_msg = tools.ColouriseRed(
			"\n[Error on line " + line_number + ", position " + token_pos +
			"]\n",
		)
		// If there is a line of code passed to 
		if full_loc != "n/a" {
			error_msg += tools.ColouriseYellow("\n[Line of Code]\n")
			error_msg += fmt.Sprintf("%s\n\n", full_loc)
		}
	}
	// Append the user error
	error_msg += fmt.Sprintf(
		"%s%s\n",
		tools.ColouriseYellow("\n[Message]\n"),
		error_message,
	)
	// Put it all together
	fmt.Printf("%s\n", error_msg)
	// Abandon ship
	os.Exit(0)
}

/*
	Report an error with the first word capitalised as need be. This is
	particularly helpful in those moments where we are dealing with errors that
	conform to Go standards (ie. first letter is lower case and there is no
	trailing punctuation). Parameters include error_message, the error message
	itself and line_number, the line number that triggered the error. Returns
	nothing.
*/
func ReportWithFixes(
	error_message string, line_number string, token_pos string,
	full_loc string) {
	error_message = strings.ToTitle(string(error_message[0])) +
					error_message[1:] + "."
	Report(error_message, line_number, token_pos, full_loc)
}