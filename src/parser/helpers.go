/*
The helpers module provides functions that help to clean up and check the lines

	of the script before the three main parsing functions -- Start(),
	Tokenise(), and Call() -- are run.
*/
package parser

import (
	"appetit/tools"
	"appetit/values"
	"fmt"
	"slices"
	"strings"
)

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
		/* Create a version of the line of code as a string and strip off any
		whitespace from the beginning and end.
		*/
		line_as_string := strings.TrimSpace(string(lines[line]))
		/* Get the length of the line just in case it's zero (ie. there's a
		blank line)
		*/
		length_of_line := len(line_as_string)

		// If the length of the line is not zero (ie. it's not empty)
		if length_of_line != 0 {
			// If the first line is not a comment
			if string(line_as_string[0]) != values.SYMBOL_COMMENT {
				// Append the comment-less line to the list of comment_free_lines
				comment_free_lines = append(comment_free_lines, lines[line])
			} else {
				/* Here, we're replacing the whole line with the comment with
				just the symbol. We don't want to remove it entirely as
				doing this throws off the line counting (which is helpful)
				for error reporting. Further, we want to remove the
				contents of the comment in case it would trigger a parsing
				error (eg. '"Hello' which is an incomplete string) so we
				replace it with something safe.
				*/
				comment_free_lines = append(
					comment_free_lines, values.SYMBOL_COMMENT)
			}
		} else if length_of_line == 0 {
			/* Otherwise, if the length of the line is zero (ie. it's an empty
			line), add it as a blank line as this needs to be counted as a
			meaningful line for error reporting.
			*/
			comment_free_lines = append(comment_free_lines, " ")
		}
	}
	// Return a slice of comment free lines
	return comment_free_lines
}

/*
This function checks that any minver statement is the first statement call. It
also checks for duplicate minver calls, thus checking to ensure that there is
only one. Returns a boolean to note whether the use of minver is valid (and
false otherwise) and a string holding a descriptive error message.
*/
func CheckValidMinverLocationCount(lines []string) (bool, string) {
	// Hold a list of the statement calls in the script
	var stmt_list []string
	// Loop over the lines and extract the statement calls
	for _, line := range lines {
		// If the lines is not a comment or blank line
		if line != values.SYMBOL_COMMENT && line != " " {
			/* Trim any blank space on either side of the line. What is most
			important here is the blank spaces at the beginning if someone
			indents a line of the script.
			*/
			trimmed_line := strings.TrimSpace(line)
			/* Get the statement name by splitting the line and extracting the
			first elements (ie. the statement name).
			*/
			stmt_name := strings.Split(trimmed_line, " ")[0]
			// Append the statement name to our list.
			stmt_list = append(stmt_list, stmt_name)
		}
	}

	// Determine if the statement list contains a minver call
	minver_present := slices.Contains(stmt_list, "minver")
	// If it does have a minver call
	if minver_present {
		// Hold the count of minver statements
		minver_count := 0
		// Loop over the stmt_list to count how many minver calls there are
		for _, stmt_name := range stmt_list {
			if stmt_name == "minver" {
				minver_count += 1
			}
		}

		/* If there's more than one minver statement, return false with an
		error message
		*/
		if minver_count > 1 {
			return false, fmt.Sprintf("There are multiple "+
				tools.ColouriseCyan("minver")+
				" calls in your script, specifically %v. Ensure that you "+
				"only have one and ensure that it is the first line "+
				"of your script.", minver_count)
		}
		/* If the first element (ie. the first line) is a minver call, return
		true
		*/
		if stmt_list[0] == "minver" {
			return true, ""
		} else if stmt_list[1] == "minver" && stmt_list[0][0:2] == "#!" {
			/* If the statement list has minver as the second element and the
			first one is a shebang line, we can also return true.
			*/
			return true, ""
		} else {
			// If we've gotten here, we can assume that there is no valid minver call.
			return false, "The " + tools.ColouriseCyan("minver") + " statement " +
				"needs to be the first line of the script. This helps to ensure " +
				"that the script is able to execute and doesn't fail part of " +
				"the way through. Move your " + tools.ColouriseCyan("minver") +
				" statement to the top of the script."
		}
	}
	return true, ""
}
