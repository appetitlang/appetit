/*
This file contains all token related code including the struct, struct methods
for the token, token related functions, and test token slices.
*/
package parser

import (
	"appetit/utils"
	"encoding/json"
	"fmt"
)

/*
The Token type houses information about a particular token and serves as an
object that houses the relevant information. The structure of the token is
as follows:
  - FullLineOfCode [string]: the full line of code that the token is embedded
    in which is helpful for error reporting
  - LineNumber [int]: the line number
  - TokenPosition [string]: the starting position (column) of the token in
    question which is helpful for error reporting
  - TokenValue [string]: the actual string value of the token
  - TokenType [string]: the type of the token as a string
  - NonCommentLineNumber [int]: the line counter for lines that aren't
    comments, that is, lines with statement calls
*/
type Token struct {
	FullLineOfCode       string
	LineNumber           int
	NonCommentLineNumber int
	TokenPosition        string
	TokenValue           string
	TokenType            string
}

/*
A helper method to make Token printing easier. It takes no parameters nor
does it return anything.
*/
func (token *Token) PrintToken() {
	// Marshal the toke into JSON and indent it with a tab
	indented_json, err := json.MarshalIndent(token, "", "\t")
	// If there's an error, report it
	if err != nil {
		fmt.Println("Error marshalling the token to an indented JSON: ", err)
	}
	// Print the formatted token
	fmt.Printf("Token: %s", string(indented_json))
}

/*
Similar to the above but focuses on printing out a token with some styling
applied.
*/
func (token *Token) PrintPrettyToken() {

	// Hold the "dot point" symbol from token properties
	dot_point := utils.ColouriseMagenta("  ::")

	// Print out the full line of code
	fmt.Printf(
		"%s Full Line of Code: %s\n",
		dot_point,
		token.FullLineOfCode,
	)
	/*
		Print out the line number. While this may seem redundant, it's
		helpful to know that the line number is being added to a token
		properly.
	*/
	fmt.Printf(
		"%s Line Number: %d\n",
		dot_point,
		token.LineNumber,
	)
	/*
		Print out the non-comment line number, the line counter that
		counts the non-comment related lines.
	*/
	fmt.Printf(
		"%s Non Comment Line Number: %d\n",
		dot_point,
		token.NonCommentLineNumber,
	)
	// Print out the token column/position (ie. where the token starts)
	fmt.Printf(
		"%s Position: %s\n",
		dot_point,
		token.TokenPosition,
	)
	// Print out the token type
	fmt.Printf(
		"%s Type: %s\n",
		dot_point,
		token.TokenType,
	)
	// Print out the token value
	fmt.Printf(
		"%s Value: %s\n",
		dot_point,
		token.TokenValue,
	)
}

/*
Print out a token to the console, properly formatted, for viewing on the
console. Takes one parameter: the tokens on a line to be printed. Note that
this does not, for obvious reasons, print out a tokenised version of a blank
line.
*/
func PrintTokenInfo(tokens []Token) {

	/*
		First up, we need to keep track of both what line we're on (cur_line)
		and what token we're working with on that line (cur_token_number).
	*/
	// Hold the current line
	cur_line := 0

	// Counter for the token on the current line
	cur_token_number := 1

	// Loop over the tokens
	for token := range tokens {
		// If the token value isn't nothing (ie. it's something meaningful)
		if tokens[token].TokenValue != "" {
			/*
				Once we've established that the current line has something
				meaningful, we need to check if this is a token on the current
				line of interest or a token on a line different from the one
				that we last looked at. If the current line doesn't equal
				cur_line (ie. we've encountered a new line of tokens).
			*/
			if tokens[token].LineNumber != cur_line {
				/*
					Reset the current token counter so that we accurately
					number the tokens in the output.
				*/
				cur_token_number = 1
				/*
					Since we're working on a new line of code here, we print
					out a header with the line number and the full line of code
					for reference.
				*/
				// Print the token line number header
				fmt.Printf(
					utils.ColouriseGreen("\n\n[Line %d] "),
					tokens[token].LineNumber,
				)
				fmt.Printf(
					"%s\n\n",
					tokens[token].FullLineOfCode,
				)

				/*
					Update the cur_line with the line number of the current
					token so that we can continue to track what line the last
					token was from (in this case, the current token).
				*/
				cur_line = tokens[token].LineNumber
				// Otherwise...
			} else {
				// Print a new blank line as a spacer for the next token
				fmt.Println()
			}
			// Print out the current token counter value
			fmt.Printf(
				utils.ColouriseRed("  [Token %d]\n"),
				cur_token_number,
			)
			// Print out a stylised token using a Token struct method
			tokens[token].PrintPrettyToken()

			// Increment the current line token counter
			cur_token_number += 1
		}
	}
}

/*
	The following Token{} slices are useful for testing purposes that model
	ideal token arrangements for various statement calls.
*/

// A simple ask statement call
var TEST_ASK = []Token{
	{
		FullLineOfCode:       "ask \"Greeting: \" to greeting",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "ask \"Greeting: \" to greeting",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "ask",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "ask \"Greeting: \" to greeting",
		LineNumber:           1,
		TokenPosition:        "5",
		TokenValue:           "\"Greeting: \"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "ask \"Greeting: \" to greeting",
		LineNumber:           1,
		TokenPosition:        "18",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "ask \"Greeting: \" to greeting",
		LineNumber:           1,
		TokenPosition:        "21",
		TokenValue:           "greeting",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple copydirectory statement call
var TEST_COPYDIR = []Token{
	{
		FullLineOfCode:       "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "copydirectory",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "15",
		TokenValue:           "\"/home/user/test\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "33",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "36",
		TokenValue:           "\"/home/user/test2\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple copyfile statement call
var TEST_COPYFILE = []Token{
	{
		FullLineOfCode:       "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "copyfile",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
		LineNumber:           1,
		TokenPosition:        "10",
		TokenValue:           "\"/home/user/test.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
		LineNumber:           1,
		TokenPosition:        "32",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
		LineNumber:           1,
		TokenPosition:        "35",
		TokenValue:           "\"/home/user/test2.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple deletedirectory statement call
var TEST_DELETEDIR = []Token{
	{
		FullLineOfCode:       "deletedirectory \"/home/user/test/\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "deletedirectory \"/home/user/test/\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "deletedirectory",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "deletedirectory \"/home/user/test/\"",
		LineNumber:           1,
		TokenPosition:        "17",
		TokenValue:           "\"/home/user/test/\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple deletefile statement call
var TEST_DELETEFILE = []Token{
	{
		FullLineOfCode:       "deletefile \"/home/user/test.txt\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "deletefile \"/home/user/test.txt\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "deletefile",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "deletefile \"/home/user/test.txt\"",
		LineNumber:           1,
		TokenPosition:        "12",
		TokenValue:           "\"/home/user/test.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple downlaod statement call
var TEST_DOWNLOADFILE = []Token{
	{
		FullLineOfCode:       "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "download",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
		LineNumber:           1,
		TokenPosition:        "10",
		TokenValue:           "\"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
		LineNumber:           1,
		TokenPosition:        "196",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
		LineNumber:           1,
		TokenPosition:        "199",
		TokenValue:           "\"#b_home/Desktop/del.jpg\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

var TEST_EXECUTE = []Token{
	{
		FullLineOfCode:       "execute \"ls -l\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "execute \"ls -l\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "execute",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "execute \"ls -l\"",
		LineNumber:           1,
		TokenPosition:        "9",
		TokenValue:           "\"ls -l\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple exit statement call
var TEST_EXIT = []Token{
	{
		FullLineOfCode:       "exit",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "exit",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "exit",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple makedirectory call
var TEST_MAKEDIR = []Token{
	{
		FullLineOfCode:       "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "makedirectory",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber:           1,
		TokenPosition:        "15",
		TokenValue:           "\"#b_home/Downloads/testdir2\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple makefile call
var TEST_MAKEFILE = []Token{
	{
		FullLineOfCode:       "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "makefile",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber:           1,
		TokenPosition:        "10",
		TokenValue:           "\"#b_home/Downloads/testdir2.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple set of tokens for a minver call
var TEST_MINVER = []Token{
	{
		FullLineOfCode:       "minver 1",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "minver 1",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "minver",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "minver 1",
		LineNumber:           1,
		TokenPosition:        "8",
		TokenValue:           "1",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple movedir statement call
var TEST_MOVEDIR = []Token{
	{
		FullLineOfCode:       "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "movedirectory",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "15",
		TokenValue:           "\"/home/user/test\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "33",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
		LineNumber:           1,
		TokenPosition:        "36",
		TokenValue:           "\"/home/user/test2\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple movefile call
var TEST_MOVEFILE = []Token{
	{
		FullLineOfCode:       "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "movefile",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
		LineNumber:           1,
		TokenPosition:        "10",
		TokenValue:           "\"/home/user/test.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
		LineNumber:           1,
		TokenPosition:        "32",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
		LineNumber:           1,
		TokenPosition:        "35",
		TokenValue:           "\"/home/user/test\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

var TEST_PAUSE = []Token{
	{
		FullLineOfCode:       "pause 3",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "pause 3",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "pause",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "pause 3",
		LineNumber:           1,
		TokenPosition:        "7",
		TokenValue:           "3",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple run call
var TEST_RUN = []Token{
	{
		FullLineOfCode:       "run \"../samples/write.apt\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "run \"../samples/write.apt\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "run",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "run \"../samples/write.apt\"",
		LineNumber:           1,
		TokenPosition:        "5",
		TokenValue:           "\"../samples/write.apt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple set call
var TEST_SET = []Token{
	{
		FullLineOfCode:       "set name = \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "set name = \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "set",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "set name = \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "5",
		TokenValue:           "name",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "set name = \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "10",
		TokenValue:           "=",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "set name = \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "12",
		TokenValue:           "\"Hello World!\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple write call
var TEST_WRITE = []Token{
	{
		FullLineOfCode:       "write \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "write \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "write",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "write \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "7",
		TokenValue:           "\"Hello World!\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple writeln call
var TEST_WRITELN = []Token{
	{
		FullLineOfCode:       "writeln \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "writeln \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "writeln",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "writeln \"Hello World!\"",
		LineNumber:           1,
		TokenPosition:        "9",
		TokenValue:           "\"Hello World!\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple zipdirectory call
var TEST_ZIPDIR = []Token{
	{
		FullLineOfCode:       "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "zipdirectory",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "14",
		TokenValue:           "\"/home/user/test\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "32",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "35",
		TokenValue:           "\"/home/user/test2.zip\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}

// A simple zipdirectory call
var TEST_ZIPFILE = []Token{
	{
		FullLineOfCode:       "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "0",
		TokenValue:           "",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "1",
		TokenValue:           "zipfile",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "9",
		TokenValue:           "\"/home/user/test.txt\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "31",
		TokenValue:           "to",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode:       "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber:           1,
		TokenPosition:        "34",
		TokenValue:           "\"/home/user/test2.zip\"",
		TokenType:            "string",
		NonCommentLineNumber: 1,
	},
}
