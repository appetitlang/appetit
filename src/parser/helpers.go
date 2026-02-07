/*
The helpers module provides functions that serves as helpful functions to be
used across the parser modules including statements. The idea here is that the
parser "engine" needs to be home only to the three essential functions - Call(),
Start(), and Tokenise(). In light of that, this module provides support
functions for the parser "engine" that are sometimes used elsewhere (eg. the
PrepScript function is used by the stmt_run module).
*/
package parser

import (
	"appetit/utils"
	"fmt"
	"net"
	"os"
	"os/user"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
This function takes in a float64 and returns a comma seperated version of it as
a string. This is helpful for producing a human readable version of the number
in places such as the download statement. Takes in a float64 and returns a
string.
*/
func CommaSeperator(number float64) string {
	// Get the integer representation of the number
	number_int := int64(number)
	// Create a string version of the number
	string_number := strconv.Itoa(int(number_int))
	// Split the number into seperate characters
	chars := []rune(string_number)

	// Hold the final number
	final_number := ""

	/*
		Count how many digits we've looked at so that we can track where the
		comma needs to go. We're starting here at one as we will be starting
		with our first digit when this is accessed for the first time.
	*/
	comma_counter := 1
	/*
		Loop over the characters starting at the end and working from the end
		back to the beginning.
	*/
	for char_count := len(chars) - 1; char_count >= 0; char_count-- {
		/*
			If the counter is less than 3, we just append the digit to the
			beginning of the final_number.
		*/
		if comma_counter < 3 {
			// Add the digit to the beginning of the final_number
			final_number = string(chars[char_count]) + final_number
			// Increment the counter
			comma_counter++
			/*
				If we've hit a comma_counter value of 3, we need to add in the
				comma and then reset the comma_counter to 1.
			*/
		} else {
			// Add a comma and the digit to the beginning of the final_number
			final_number = "," + string(chars[char_count]) + final_number
			// Reset the counter
			comma_counter = 1
		}
	}
	/*
		If the number of digits is a multiple of 3, a leading comma will be
		there so let's remove it.
	*/
	final_number = strings.TrimLeft(final_number, ",")
	// Return the final comma seperated number.
	return final_number
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
Print out a token to the console, properly formatted. Takes one parameter:
the tokens on a line to be printed. Note that this does not, for obvious
reasons, print out a tokenised version of a blank line.
*/
func PrintTokenInfo(tokens []Token) {

	// Hold the "dot point" symbol from token properties
	dot_point := utils.ColouriseMagenta("  ::")

	// Hold the current line
	cur_line := 0

	// Token counter
	token_count := 0

	// Counter for the token on the current line
	cur_line_token_number := 1

	// Loop over the tokens
	for token := range tokens {
		// If the token value isn't nothing (ie. it's something meaningful)
		if tokens[token].TokenValue != "" {
			// Increment the token counter
			token_count += 1
			/*
				If the current line doesn't equal cur_line (ie. we've
				encountered a new line of tokens).
			*/
			if tokens[token].LineNumber != cur_line {
				// Reset the current line token counter
				cur_line_token_number = 1
				// Print the token line number header
				fmt.Printf(
					utils.ColouriseGreen("\n\nLINE %d\n"),
					tokens[token].LineNumber,
				)
				fmt.Println(tokens[token].FullLineOfCode + "\n")
				// Update the cur_line
				cur_line = tokens[token].LineNumber
				// Otherwise
			} else {
				// Print a new blank line
				fmt.Println()
			}
			// Print out the current line token counter
			fmt.Printf(
				utils.ColouriseRed("  [Token %d]\n"),
				cur_line_token_number,
			)
			// Print out the full line of code
			fmt.Printf(
				"%s Full Line of Code: %s\n",
				dot_point,
				tokens[token].FullLineOfCode,
			)
			// Print out the token column/position (ie. where the token starts)
			fmt.Printf(
				"%s Position: %s\n",
				dot_point,
				tokens[token].TokenPosition,
			)
			// Print out the token type
			fmt.Printf(
				"%s Type: %s\n",
				dot_point,
				tokens[token].TokenType,
			)
			// Print out the token value
			fmt.Printf(
				"%s Value: %s\n",
				dot_point,
				tokens[token].TokenValue,
			)
			/*
				Print out the line number. While this may seem redundant, it's
				helpful to know that the line number is being added to a token
				properly.
			*/
			fmt.Printf(
				"%s Line Number: %d\n",
				dot_point,
				tokens[token].LineNumber,
			)
			/*
				Print out the non-comment line number, the line counter that
				counts the non-comment related lines.
			*/
			fmt.Printf(
				"%s Non Comment Line Number: %d\n",
				dot_point,
				tokens[token].NonCommentLineNumber,
			)

			// Increment the current line token counter
			cur_line_token_number += 1
		}
	}
}

/*
Create any values for built in reserved variables that require building.
This addresses the empty ones in the VARIABLES map. No parameters and no
returns.
*/
func BuildReservedVariables() {

	cur_user, cur_user_error := user.Current()
	if cur_user_error != nil {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"user"] = ""
	} else {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"user"] = cur_user.Username
	}

	// Create a date
	date := time.Now()

	// Get the date in dd-mm-yyyy format
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_dmy"] = fmt.Sprintf(
		"%d-%d-%d", date.Day(), date.Month(), date.Year(),
	)

	// Get the date in yyyy-mm-dd format
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_ymd"] = fmt.Sprintf(
		"%d-%d-%d", date.Year(), date.Month(), date.Day(),
	)

	// Get the current time
	time_now := time.Now()

	// Get the time in hh-mm-ss in 24 hour format
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"time"] = fmt.Sprintf(
		"%d-%d-%d", time_now.Hour(), time_now.Minute(), time_now.Second(),
	)

	// Get the time zone, ignoring the offset as this isn't needed
	time_zone, _ := time_now.Zone()

	// Get the timezone
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"zone"] = time_zone

	// Get the hostname
	host, err := os.Hostname()
	// If there's no error, set the b_host to the hostname.
	if err == nil {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"hostname"] = host
	}
	// Get the user home directory
	home, err := os.UserHomeDir()
	// If there's no error, set the b_home to the home directory.
	if err == nil {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"home"] = home
	}

	// Get the working directory
	wd, err := syscall.Getwd()
	// If there's no error, set the b_wd to the working directory.
	if err == nil {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"wd"] = wd
	}

	/* House the ip address. This now improves on the old format which depended
	both on an external network connection and the stability of Google's
	public DNS service which is not something I want to depend on.
	*/
	// Get the interfaces
	ipaddrs, ipaddrs_err := net.InterfaceAddrs()
	// If we can't, abandon ship and save n/a to the ipv4 reserved variable
	if ipaddrs_err != nil {
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"ipv4"] = "n/a"
	}

	// Iterate over the addresses
	for _, ipv4_addresses := range ipaddrs {
		// If the address is an ip address and is not a loopback address...
		if ip, ip_ok := ipv4_addresses.(*net.IPNet); ip_ok && !ip.IP.IsLoopback() {
			// If converting it to an IPv4 address doesn't yield an error...
			if ip.IP.To4() != nil {
				// Set the IPv4 address reserved variable
				VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"ipv4"] = ip.IP.String()
			}
		}
	}
}

/*
Create a string list of the reserved variables that can be easily printed
if need be. No parameters. Returns a string representation of the list of
reserved variables.
*/
func ListReservedVariables() string {
	// Hold the list of reserved variables
	var reserved_var []string
	// For each variable in the VARIABLES map
	for vars := range VARIABLES {
		/* If the prefix -- signified by the string from 0 to the length of
		the RESERVED_VARIABLE_PREFIX -- is the RESERVED_VARIABLE_PREFIX
		*/
		if vars[0:len(SYMBOL_RESERVED_VARIABLE_PREFIX)] ==
			SYMBOL_RESERVED_VARIABLE_PREFIX {
			/* Append the reserved variable to the list of them with a
			colourised version
			*/
			reserved_var = append(reserved_var, utils.ColouriseMagenta(vars))
		}
	}
	// Sort the list of reserved variables
	slices.Sort(reserved_var)

	var var_names string
	for _, var_name := range reserved_var {
		var_names += "\n\t- " + var_name
	}
	return var_names
	// Return a joined version of this list
	// return strings.Join(reserved_var, ", ")
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
