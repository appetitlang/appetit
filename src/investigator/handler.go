/*
The handler module provides error handling and reporting for language specific
issues.
*/
package investigator

import (
	"appetit/tools"
	"fmt"
	"os"
	"strconv"
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


	fmt.Println(tools.ColouriseRed("\n[Error]"))
	fmt.Println(tools.ColouriseMagenta("Line of Code: ") + full_loc)
	fmt.Println(tools.ColouriseMagenta("Line Number:  ") + line_number)
	fmt.Println(tools.ColouriseMagenta("Position:     ") + token_pos)
	fmt.Println(tools.ColouriseRed("\n[Message]"))
	fmt.Print(error_message)
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

/*
	This handles errors reported by the tokeniser that are not language
	specific but would nonetheless cause issues for the script. For instance,
	an unmatched pair of quotation marks isn't a language error (per se) but
	would nonetheless cause an issue. Parameters here include the message
	reported back by the tokeniser.
*/
func ReportTokeniserErrors(message string, loc int) {
	/*
		Set up an elaborate switch/case to capture any anticipated errors
		Parameters include the scanner and the message that gets reported.
		Returns nothing.

		TODO: more carefully account for some of the errors reported back by
		scanner.Init(). See here: https://cs.opensource.google/go/go/+/refs/
		// tags/go1.25.1:src/text/scanner/scanner.go;l=181. Find all the
		s.error() calls.
	*/
	switch message {
	// Catch an unterminated literal
	case "literal not terminated":
		Report(
			"Your line of code has an incomplete string. Did you forget " +
			"an opening or closing quotation mark?\n\nSomething like " +
			"the following line of code will trigger this error:\n" +
			tools.ColouriseCyan("writeln ") +
			tools.ColouriseGreen("\"Hello world") +
			tools.ColouriseRed("_") + " <- (notice the lack of a " +
			"closing quotation mark here).",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	// Catch an invalid char literal
	case "invalid char literal":
		Report(
			"Your line of code use single quotation marks instead of " +
			"the required double quotation marks. See the example:\n\n" +
			tools.ColouriseCyan("writeln ") +
			tools.ColouriseGreen("'Hello world'") + " <- (notice the" +
			" lack of double quotation marks here).",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	case "comment not terminated":
		Report(
			"You've included a Go style comment as a statement call which " +
			"is not valid. Comments are single line and take the following " +
			"form:\n\n" + tools.ColouriseGrey(
			" - This is a comment."),
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	case "invalid char escape":
		Report(
			"You've included an invalid character escape. You need to use " +
			"one of the following: " + tools.ColouriseMagenta("\\n") +
			" (for new line), " + tools.ColouriseMagenta("\\t") + " (for " +
			"tab indentation).",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	default:
		/* Report everything else in their "Go form." It is hoped that, some
			day, this will not need to exist.
		*/
		Report(
			message + ". Please report this error with the erroneous line " +
			"of code as this isn't accounted for in the error checking.",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	}
}