/*
The helpers provides functions that serves as helpful functions to be used
across the parser including statements. The idea here is that the
parser "engine" needs to be home only to the three essential functions - Call(),
Start(), and Tokenise(). In light of that, this provides support
functions for the parser "engine" that are sometimes used elsewhere (eg. the
PrepScript function is used by the run statement).
*/
package parser

import (
	"appetit/utils"
	"go/token"
	"go/types"
	"os"
	"slices"
	"strings"
)

/*
Calculate the value variable to see if it's a math expression. Parameters
include loc, the line of the script (if an error needs to be reported) and
the value to calculate. Returns the value of the calculated string.
*/
func CalculateValue(loc string, value string) string {
	// Thanks to https://stackoverflow.com/a/65484336 for the below
	// Get a file set of tokens
	expression_tokens := token.NewFileSet()
	// Get the tokens and evaluate them
	tv, err := types.Eval(expression_tokens, nil, token.NoPos, value)
	// If there's an error trying to evaluate the tokens, return the value
	if err != nil {
		return value
	}
	// Return the calculated value as a string.
	return tv.Value.ExactString()
}

/*
Prepare a script for execution. Here, open it up and strip the comments. This
returns a slice of the lines of the script with comments replaced with just
the comment symbol.
*/
func PrepScript(file_name string) []string {
	// Create a slice for the lines of code
	var lines []string

	// Read the file
	script, err := os.ReadFile(file_name)
	// If the file couldn't be opened
	if err != nil {
		// Report the error
		Report(
			"Unknown file: "+utils.ColouriseMagenta(file_name)+".",
			"n/a",
			"n/a",
			"n/a",
		)
		// Exit the script
		os.Exit(0)
	}
	// Split the lines of the script by lines and into strings
	lines = strings.Split(string(script), "\n")
	// Remove the comments from the script
	lines = RemoveComments(lines)
	/*
		Return the lines which, by this point, will be a slice of lines where
		the	comments have been replaced with "blank" comments (ie. "-" only)
	*/
	return lines
}

/*
This function removes any comments from the slice that houses the lines of
code that are eventually passed to the tokeniser. Parameters include lines,
the lines of the script inclusive of any comments. Returns a slice of
strings that represents the script absent any lines of comments.
*/
func RemoveComments(lines []string) []string {
	// Hold the comment free lines
	var comment_free_lines []string
	// Loop over the lines
	for line := range lines {
		/*
			Create a version of the line of code as a string and strip off any
			whitespace from the beginning and end.
		*/
		line_as_string := strings.TrimSpace(string(lines[line]))
		/* Get the length of the line just in case it's zero (ie. there's a
		blank line)
		*/
		length_of_line := len(line_as_string)

		/*
			If the length of the line is not zero (ie. it's not empty), we know
			that we are working with a line of code or a comment so we need to
			check it.
		*/
		if length_of_line != 0 {
			/*
				If the first character in the line is not a comment symbol, we
				can continue on the assumption that the line of code is a
				comment.
			*/
			if string(line_as_string[0]) != SYMBOL_COMMENT {
				/*
					Append the line to the list of comment_free_lines on the
					assumption that it is not a comment and thus needs to be
					left untouched.
				*/
				comment_free_lines = append(comment_free_lines, lines[line])
			} else {
				/*
					On the assumption. that the line is a comment because the
					first character is a comment symbol, we add a line to the
					comment_free_lines that is just the comment symbol. This
					enables us to keep track of lines accurately (ie. a comment
					line is still a line that needs to be counted) just as we
					want to strip away the comment in case it has a syntax
					error (commented lines, otherwise, would get parsed and if
					they had a syntax error, they would trigger an error).
				*/
				comment_free_lines = append(
					comment_free_lines,
					SYMBOL_COMMENT,
				)
			}
		} else if length_of_line == 0 {
			/*
				Otherwise, if the length of the line is zero (ie. it's an empty
				line), add it as a blank line as this needs to be counted as a
				meaningful line for error reporting.
			*/
			comment_free_lines = append(comment_free_lines, " ")
		}
	}
	/*
		Return a slice of comment free lines. This will include lines that are
		empty comments (ie. "-") and blank lines.
	*/
	return comment_free_lines
}

/*
Create a string list of the statement names that can be easily printed if
need be. No parameters. Returns a string representation of the list of
statements.
*/
func ListStatements() string {
	// Hold the list of statements
	var statement_names []string

	// For each statement in STATEMENT_NAMES
	for stmt := range STATEMENT_NAMES {
		// Append the statement name to the list of statement_names
		statement_names = append(
			statement_names, utils.ColouriseMagenta(STATEMENT_NAMES[stmt]),
		)
	}
	// Sort the list of statement names
	slices.Sort(statement_names)
	// Return a joined version of this list
	return strings.Join(statement_names, ", ")
}
