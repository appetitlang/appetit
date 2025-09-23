/*
The fixer module deals with fixing strings.
*/
package tools

import (
	"os"
	"strings"
)

/*
	Remove the leading and trailing quotation mark by being sensible. Go's
	strings.Trim* functions remove all so if someone has something like
	"Hello World\"", the strings.Trim* functions would remove both quotation
	marks at the end. This function sensibly checks for only one quotation mark
	on either end. Parameters include the input which is the string to fix the
	quotation marks on. Returns a fixed string, that is, one with leading and
	trailing quotation marks handled.
*/
func FixStringQuotations(input string) string {
	// Get the first character
	first_char := string(input[0])
	// Get the last character
	last_char := string(input[len(input)-1])

	/* First check - see if the first and last character are quotation marks
		and if so, replace the leading and trailing quotation marks
	*/
	if first_char == "\"" && last_char == "\"" {
		/* Return everything between the second character and the second last
			one.
		*/
		return string(input[1:len(input)-1])
	/* If the first and last character do not need to be stripped (eg.
		printing an integer), simply return the input
	*/
	} else {
		return input
	}
}

/*
	Fix escape characters in strings. Parameters include the input which is the
	string to fix the escapes on. Returns a fixed string, that is, one with
	escapes handled.
*/
func FixStringEscapes(input string) string {
	// Fix any escapes of quotation marks with the quotation marks themselves
	fixed_string := strings.ReplaceAll(input, "\\\"", "\"")
	// Replace new lines with actual new lines
	fixed_string = strings.ReplaceAll(fixed_string, "\\n", "\n")
	// Replace carriage returns with actual carriage returns
	fixed_string = strings.ReplaceAll(fixed_string, "\\r", "\r")
	// Return the fixed_string
	return fixed_string
}

/*
	Fix paths if they are missing a trailing path seperator. This is handy as
	sometimes paths are required and we want to ensure that trailing path
	seperators are present.
*/
func FixPathSeperators(input string) string {
	// If the last character is not a path seperator...
	if input[len(input)-1] != os.PathSeparator {
		// Add the path seperator and return that
		return input + string(os.PathSeparator)
	}
	/* If we've gotten here, the last character is a path seperator so we can
		return the string as is
	*/
	return input
}

/*
	This is a helper function that consolidates some of the fixer functions
	above. This is done as the functions called here are often called together
	so this can cut down on repetition. Takes in a string, fixes it, and
	returns it.
*/
func FixStringCombined(input string) string {
	// Fix the quotations so that they are properly escaped
	result := FixStringQuotations(input)
	// Fix any escapes that need to be accounted for
	result = FixStringEscapes(result)
	// Return the fixed string
	return result
}