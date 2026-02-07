/*
The checks and fixes modules houses parser functions that do two things:
  - Run checks on statement calls and tokens to make sure that they are
    conformant with the language
  - Run fixes on statement values to make sure that they are internally
    represented appropriately and outwardly visible as you would expect
*/
package parser

import (
	"appetit/utils"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
Check that an action is signaled by an appropriate
values.SYMBOL_ACTION. Parameters include loc, the line of the script
(if an error needs to be reported and the value to check. Returns
error if the action is incorrect.
*/
func CheckAction(loc string, action_value string) error {
	// If the action_value is the action symbol...
	if action_value != SYMBOL_ACTION {
		// Return an error
		return fmt.Errorf(
			"an action was made using an invalid action statement (%s), "+
				"please ensure that you use %s",
			utils.ColouriseMagenta(action_value),
			utils.ColouriseMagenta(SYMBOL_ACTION),
		)
	}
	// If we've gotten here, then we have no error
	return nil
}

/*
This function checks that any minver statement is the first statement call. It
also checks for duplicate minver calls, thus checking to ensure that there is
only one. Returns a boolean to note whether the use of minver is valid (and
false otherwise) and a string holding a descriptive error message.
*/
func CheckValidMinverLocationAndCount(lines []string) (bool, string) {
	// Hold a list of the statement calls in the script
	var stmt_list []string
	// Loop over the lines and extract the statement calls
	for _, line := range lines {
		// If the lines is not a comment or blank line
		if line != SYMBOL_COMMENT && line != " " {
			/*
				Trim any blank space on either side of the line. What is most
				important here is the blank spaces at the beginning if someone
				indents a line of the script.
			*/
			trimmed_line := strings.TrimSpace(line)
			/*
				Get the statement name by splitting the line and extracting the
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

		/*
			If there's more than one minver statement, return false with an
			error message.
		*/
		if minver_count > 1 {
			return false, fmt.Sprintf("There are multiple "+
				utils.ColouriseCyan("minver")+
				" calls in your script, specifically %v. Ensure that you "+
				"only have one and ensure that it is the first line "+
				"of your script.", minver_count)
		}
		/*
			If the first element (ie. the first line) is a minver call, return
			true.
		*/
		if stmt_list[0] == "minver" {
			return true, ""
		} else if stmt_list[1] == "minver" && stmt_list[0][0:2] == "#!" {
			/*
				If the statement list has minver as the second element and the
				first one is a shebang line, we can also return true.
			*/
			return true, ""
		} else {
			/*
				If we've gotten here, we can assume that there is no valid
				minver call.
			*/
			return false, "The " + utils.ColouriseCyan("minver") +
				" statement needs to be the first line of the script. This " +
				"helps to ensure that the script is able to execute and " +
				"doesn't fail part of the way through. Move your " +
				utils.ColouriseCyan("minver") + " statement to the top of " +
				"the script."
		}
	}
	return true, ""
}

/*
Check if a line of code is a shebang line that points to a valid interpreter.
This function therefore does two things. First, it checks if the first two
characters are a shebang and second, it checks if the path after the shebang
is a valid path. Returns a boolean to signify if we are good to go (true) or
not (false) and a second boolean if we have a valid (true) or invalid (false)
path.
*/
func CheckShebang(line string) (bool, bool) {
	first_two_chars := line[0:2]
	_, path_error := os.Stat(line[3:])

	// If there is no valid shebang character but a valid interpreter path
	if first_two_chars != "#!" && path_error == nil {
		// Return false for the shebang, true for the path
		return false, true
	} else if first_two_chars != "#!" && path_error != nil {
		// Return false for both the shebang and path
		return false, false
	} else if first_two_chars == "#!" && path_error == nil {
		// Return true for both the shebang and path
		return true, true
	} else {
		// Return true for the shebang, false for the path
		return true, false
	}
}

/*
Check that something is a valid statement. Parameters include loc, the line
of the script (if an error needs to be reported) and the value to check.
Returns bool, true if value_name is a valid statement name.
TODO: Remove loc parameter as it's not used
*/
func CheckIsStatement(loc string, value_name string) bool {
	// See if the value passed to this is a valid statement name
	return slices.Contains(STATEMENT_NAMES, value_name)
}

/*
Check that the prefix of the variable isn't the reserved one. Parameters
include loc, the line of the script (if an error needs to be reported and
the value to check. Returns bool, true if the variable does not start
with the reserved variable prefix.
*/
func CheckVariablePrefix(
	loc string, prefix string, variable_name string) error {
	// If the prefix is not the reserved variable prefix...
	if prefix != SYMBOL_RESERVED_VARIABLE_PREFIX {
		// Return nil as we're good to go
		return nil
	}
	// Report back an error while noting the reservation of the variable prefix
	return fmt.Errorf(
		"you've named a variable %s which starts with %s (this is not "+
			"allowed). If you were trying to use a reserved variable, "+
			"consult the following list: %s",
		utils.ColouriseYellow(variable_name),
		utils.ColouriseYellow(SYMBOL_RESERVED_VARIABLE_PREFIX),
		ListReservedVariables(),
	)
}

/*
Check that an assignment operator is a valid assignment operator.
Parameters include loc, the line of the script (if an error needs to be
reported and the value to check. Returns bool, true if the assignment
operator is valid.
*/
func CheckValidAssignment(loc string, value_name string) error {
	// If the value_name is not the ASSIGNMENT_OPERATOR...
	if value_name != SYMBOL_OPERATOR_ASSIGNMENT {
		// Report back an error while providing the operator to use
		return fmt.Errorf(
			"an assignment was made using an invalid operator (%s), please "+
				"ensure that you use %s"+value_name,
			utils.ColouriseMagenta(SYMBOL_OPERATOR_ASSIGNMENT),
		)
	}
	return nil
}

/*
Report whether there is an appropriate number of tokens. Parameters include
the tokens and the required_number which is the required number of tokens
to have in the line. Returns a bool, true if there is a valid number of
tokens and false otherwise. Additionally, an error is reported to make the
gopher happy.
*/
func CheckValidNumberOfTokens(
	tokens []Token, required_number int) (bool, error) {
	/* Get the token count and subtracting one to account for the fact that the
	line number is included.
	*/
	token_count := len(tokens) - 1

	// If the token_count does not equal the required number of tokens
	if token_count != required_number {
		// Return false an an error message
		return false, fmt.Errorf("invalid number of tokens")
	}

	// If we got here, we can
	return true, nil
}

/*
Check to ensure that a file exists. Parameters include the file name
itself. Returns a boolean, true if the file exists, false if it does not.
*/
func CheckFileExists(file_name string) bool {
	// Check whether the file exists
	_, err := os.Stat(file_name)
	/* Return whether it does or does not exist. Here, we're returning true if
	there is not an error (and thus the file is presumed to exist) and false if
	there is an error (and thus the file is presumed not to exist).
	*/
	return !errors.Is(err, os.ErrNotExist)
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
		return string(input[1 : len(input)-1])
		/* If the first and last character do not need to be stripped (eg.
		printing an integer), simply return the input
		*/
	} else {
		return input
	}
}
