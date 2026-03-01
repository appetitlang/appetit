/*
These functions provide error handling and reporting for language specific
issues.
*/
package parser

import (
	"appetit/utils"
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

	// Get the token position and convert it to an integer
	position, _ := strconv.Atoi(token_pos)
	/*
		Set the header for the line of code header so that we can also get its
		length
	*/
	loc_title := "Line of Code: "
	/*
		Get the position but subtract one as we want to insert the error arrow
		at the right place
	*/
	position += len(loc_title)
	// Set up the error arrow
	error_pos_symbol := utils.ColouriseRed("^") // ⇈
	// Print the error information
	fmt.Println(utils.ColouriseRed("\n[ERROR]\n\n[Location]"))
	fmt.Println(utils.ColouriseMagenta(" Line Number: ") + line_number)
	fmt.Println(utils.ColouriseMagenta("    Position: ") + token_pos)
	fmt.Println(utils.ColouriseMagenta(loc_title) + full_loc)
	/*
		Print out the arrow to note where the error starts by repeating some
		blank spaces at the beginning that is the length of the line of code
		header.
	*/
	fmt.Printf("%s%s\n",
		strings.Repeat(" ", position-1),
		error_pos_symbol,
	)
	fmt.Println(utils.ColouriseRed("\n[Description]"))
	fmt.Printf("%s\n\n", error_message)
	// Abandon ship
	os.Exit(0)
}

/*
Report an error. Parameters include error_message, the error message
itself. Unlike Report(), this is designed for errors that aren't line or
syntax specific. This includes something like
parser.CheckValidMinverLocationCount() as an example or passing the
interpreter no script. Returns nothing.
*/
func ReportSimple(error_message string) {
	fmt.Println(utils.ColouriseRed("\n[ERROR]"))
	fmt.Println(error_message + "\n")
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
func ReportWithFixes(error_message string, line_number string,
	token_pos string, full_loc string) {
	// Capitalise the error message
	error_message = strings.ToTitle(string(error_message[0])) +
		error_message[1:] + "."
	// Report the error
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

		If scanner.Init() throws an error, this is the place to fix it.
		See here: https://cs.opensource.google/go/go/+/refs/tags/go1.25.1:src/
		text/scanner/scanner.go;l=181. Find all the s.error() calls. This is an
		ongoing to-do that is likely only to be resolved by way of more testing
		and usage.
	*/
	switch message {
	// Catch an unterminated literal
	case "literal not terminated":
		ReportSimple(
			"Line " + strconv.Itoa(loc) + " has an incomplete string. Did " +
				"you forget an opening or closing quotation mark? Something " +
				"like the following line of code will trigger this error:" +
				"\n\n\t" + utils.ColouriseCyan("writeln ") +
				utils.ColouriseGreen("\"Hello world") +
				utils.ColouriseRed("_"),
		)
	// Catch an invalid char literal
	case "invalid char literal":
		Report(
			"Your line of code use single quotation marks instead of "+
				"the required double quotation marks. See the example:\n\n"+
				utils.ColouriseCyan("writeln ")+
				utils.ColouriseGreen("'Hello world'")+" <- (notice the"+
				" lack of double quotation marks here).",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	case "comment not terminated":
		Report(
			"You've included a Go style comment as a statement call which "+
				"is not valid. Comments are single line and take the "+
				"following "+"form:\n\n"+utils.ColouriseGrey(
				" - This is a comment."),
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	case "invalid char escape":
		Report(
			"You've included an invalid character escape. You need to use "+
				"one of the following: "+utils.ColouriseMagenta("\\n")+
				" (for new line), "+utils.ColouriseMagenta("\\t")+" (for "+
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
			message+". Please report this error with the erroneous line "+
				"of code as this isn't accounted for in the error checking."+
				"It may be some time before you see this error message fixed"+
				" in a release.",
			strconv.Itoa(loc),
			"n/a",
			"n/a",
		)
	}
}

func Warning(warning string, line_number string) {
	fmt.Println(utils.ColouriseYellow("\n[WARNING]"))
	fmt.Println(utils.ColouriseMagenta(" Line Number: ") + line_number)
	fmt.Println(warning)
}
