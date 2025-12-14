/*
The parser module deals with tokenising the script and delegating to the
statement modules. This module deals specifically with outputting developer
information as part of parsing.
*/
package parser

import (
	"appetit/tools"
	"appetit/values"
	"fmt"
)

/*
	Print out a token to the console, properly formatted. Takes one parameter:
	the tokens on a line to be printed. Note that this does not, for obvious
	reasons, print out a tokenised version of a blank line.
*/
func PrintTokenInfo(tokens []values.Token) {

	// Hold the "dot point" symbol from token properties
	dot_point := tools.ColouriseMagenta("  ::")

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
			/* If the current line doesn't equal cur_line (ie. we've
				encountered a new line of tokens)
			*/
			if tokens[token].LineNumber != cur_line {
				// Reset the current line token counter
				cur_line_token_number = 1
				// Print the token line number header
				fmt.Printf(
					tools.ColouriseGreen("\n\nLINE %d\n"),
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
				tools.ColouriseRed("  [Token %d]\n"),
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
			/* Print out the line number. While this may seem redundant, it's
				helpful to know that the line number is being added to a token
				properly.
			*/
			fmt.Printf(
				"%s Line Number: %d\n",
				dot_point,
				tokens[token].LineNumber,
			)
			/* Print out the non-comment line number, the line counter that
				counts the non-comment related lines
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