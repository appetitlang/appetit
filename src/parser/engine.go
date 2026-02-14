/*
The engine deals with tokenising the script and delegating to the
statement functions. This is home to only three functions:
  - Tokenise() - this will tokenise a line and return a line that has been
    tokenised.
  - Start() - this starts the process of executing the script or, where needed,
    tokenising empty lines to allow for clean execution.
  - Call() - this executes the appropriate statement functions.
*/
package parser

import (
	"appetit/utils"
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"text/scanner"
)

/*
Tokenise each line of the code a create a slice. Parameters include
line_of_script (the non-tokenised line of the script that we are working to
tokenise) and loc, the lines of code that we are working with where:
  - element 0 is the actual line of code;
  - element 1 is the counter used to track non-comment lines of code,
    those lines of code that are deemed significant (ie. those where
    the line has a statement call).

Returns a slice of Tokens that represents a tokenised line where the first
element is the line number and each subsequent element is a token in the
line.
*/
func Tokenise(
	line_of_script string,
	line_number int, non_comment_line_number int) []Token {
	// Set up a text scanner to tokenise the script
	var ls scanner.Scanner
	// Initialise against a line of the script passed
	tokeniser := ls.Init(strings.NewReader(line_of_script))
	/*
		Set up an error handler. Thanks to https://github.com/AOSPworking/
		pkg_repo_tool/blob/82e14e49258f/parser/parser.go.
	*/

	// Catch errors with the tokeniser and handle them
	tokeniser.Error = func(s *scanner.Scanner, msg string) {
		// Call the utilss
		ReportTokeniserErrors(msg, line_number)
	}

	// Create a slice that houses tokens for each part of the line of code
	var tokenised_line []Token
	// Create the line of code token
	loc_token := Token{
		FullLineOfCode:       line_of_script,
		LineNumber:           line_number,
		TokenPosition:        strconv.Itoa(tokeniser.Position.Column),
		TokenValue:           "",
		TokenType:            reflect.TypeOf(tokeniser.TokenText()).String(),
		NonCommentLineNumber: non_comment_line_number,
	}
	// Add the line of code to the tokenised line
	tokenised_line = append(tokenised_line, loc_token)

	// Loop over the tokens
	for tok := tokeniser.Scan(); tok != scanner.EOF; tok = tokeniser.Scan() {
		// Create a token
		token := Token{
			FullLineOfCode:       line_of_script,
			LineNumber:           line_number,
			TokenPosition:        strconv.Itoa(tokeniser.Position.Column),
			TokenValue:           tokeniser.TokenText(),
			TokenType:            reflect.TypeOf(tokeniser.TokenText()).String(),
			NonCommentLineNumber: non_comment_line_number,
		}
		/*
			Append the token to the slice of tokens that will be passed to
			statement functions in the Call().
		*/
		tokenised_line = append(tokenised_line, token)
	}

	// Append the tokens to the TOKEN_TREE
	TOKEN_TREE = append(TOKEN_TREE, tokenised_line...)

	// Return the tokenised line
	return tokenised_line
}

/*
Start executing commands in a script by passing the lines to the tokeniser
and the Call() function. The only parameter is the lines of the script. No
returns.
*/
func Start(lines []string, dev_mode bool) {
	/*
		Before we start parsing, set any reserved variables that require
		"computation". Do a quick check to make sure that the minver statement
		call, if present, is the first language specific call.
	*/
	valid_minver, message := CheckValidMinverLocationAndCount(lines)
	// If it's not appropriately located in the script, error out
	if !valid_minver {
		ReportSimple(message)
	}

	// Counter for the non-comment lines
	non_comment_line_count := 1
	// Loop over the lines
	for line := range lines {
		// Create a string version of the line
		line_as_string := string(lines[line])

		// Hold the length of the line
		line_length := len(line_as_string)

		/*
			It's possible that the line has no length (ie. a blank line) so we
			need to skip over them.
		*/
		if line_length > 0 {
			/*
				Here, we start by checking to see what the first character of
				the line is to see if it's a SYMBOL_COMMENT. If it is not (ie.
				it's a line that requires parsing), we send the line to the
				tokeniser and the Call() function. While the RemoveComments()
				function "removed" the comments, it only did so as far as it
				stripped away all the characters instead of the comment symbol
				so there are still comment line in that need to be accounted
				for here.
			*/
			if string(line_as_string[0]) != SYMBOL_COMMENT {
				/*
					Tokenise the line, adding one to the line number to account
					for the starting from zero.
				*/
				tokenised_line := Tokenise(
					lines[line],
					line+1, non_comment_line_count)
				// Increment the non_comment_line_count
				non_comment_line_count += 1
				// If dev_mode is enabled
				if dev_mode {
					// Print the tokens
					PrintTokenInfo(tokenised_line)
					// If dev mode is not enabled, delegate execution
				} else {
					// Delegate to the statements package to start execution
					Call(tokenised_line)
				}
			}
		} else if line_length == 0 {
			/*
				Pass an empty string. This is needed; if we don't pass this
				here, blank lines are skipped which results in line counts not
				being accurate. This is held in nothing as the tokenised value
				is irrelevant so we can dispense with this.
			*/
			_ = Tokenise(" ", line+1, -1)
		}
	}
}

/*
Start delegating lines of tokens to the functions that will do the work of
actually executing functionality. Parameters include tokens, a slice of
strings that contains the tokens. Returns nothing.
*/
func Call(tokens []Token) {
	/*
		Build the list of reserved variables so that each statement call has
		access to an up to date set of variables.
	*/
	BuildReservedVariables()
	/*
		This will catch empty lines as the slice will have a line number but no
		other elements. So, if the number of tokens is greater than 1, we can
		assume that the line is something that may require a statement call.
	*/
	//
	if len(tokens) > 1 {

		// Get the statement name
		stmt_name := tokens[1].TokenValue
		// Create a map of statmements and their associated function calls
		statement_map := map[string]func(){
			"ask":             func() { Ask(tokens) },
			"copydirectory":   func() { CopyPath(tokens) },
			"copyfile":        func() { CopyFile(tokens) },
			"deletedirectory": func() { DeletePath(tokens) },
			"deletefile":      func() { DeleteFile(tokens) },
			"download":        func() { Download(tokens) },
			"execute":         func() { ExecuteCommand(tokens) },
			"exit":            func() { Exit(tokens) },
			"makedirectory":   func() { CreatePath(tokens) },
			"makefile":        func() { MakeFile(tokens) },
			"minver":          func() { MinVer(tokens) },
			"movedirectory":   func() { MovePath(tokens) },
			"movefile":        func() { MoveFile(tokens) },
			"pause":           func() { Pause(tokens) },
			"run":             func() { Run(tokens) },
			"set":             func() { Set(tokens) },
			"write":           func() { Writeln(tokens, false) },
			"writeln":         func() { Writeln(tokens, true) },
			"zipdirectory":    func() { ZipFromPath(tokens) },
			"zipfile":         func() { ZipFromFile(tokens) },
		}

		/*
			Here, we consolidate the keys from the statement_map map and set
			the STATEMENT_NAMES to the list of keys. We only do this if
			the list is empty, helping to cut down how many times that this may
			need to happen. Thanks to https://stackoverflow.com/q/21362950.
		*/
		if len(STATEMENT_NAMES) == 0 {
			STATEMENT_NAMES = slices.Collect(maps.Keys(statement_map))
		}

		/*
			First up, we need to check to see if there is a shebang line. This
			is particularly helpful for *nix users who might need or want to
			run this without invoking the interpreter directly. We also check
			that the shebang line points to a valid file.
		*/
		valid_shebang, _ := CheckShebang(
			tokens[0].FullLineOfCode)

		/*if !valid_interpreter {
			Warning(
				"The interpreter passed as a shebang does not exist.",
				strconv.Itoa(tokens[0].LineNumber),
			)
		}*/

		if valid_shebang {
			/*
				Set the SHEBANG_PRESENT value to true so that
				MinVer() calls can ignore that the minver statement isn't on
				the first line and is, instead, on the second line.
			*/
			SHEBANG_PRESENT = true
			// Break out of the function call
			return
		} else {
			//valid_stmt := CheckIsStatement(tokens[1].FullLineOfCode)
			valid_stmt := CheckIsStatement(tokens[1].TokenValue)

			if !valid_stmt {
				// Get the line of the script
				loc := strconv.Itoa(tokens[0].LineNumber)
				// Get full line of code
				full_loc := tokens[0].FullLineOfCode
				// Get the statement name
				stmt_name := tokens[1].TokenValue
				/*
					This executes if the statement doesn't exist (ie. exists is
					false). Report back that it doesn't exist with a list of
					valid statements.
				*/
				Report(
					"The statement passed - "+utils.ColouriseYellow(stmt_name)+
						" - is not a valid statement. Valid statements "+
						"include "+ListStatements()+".",
					loc,
					tokens[0].TokenPosition,
					full_loc,
				)
			}
		}

		/*
			If the statement is in the statement map, execute the corresponding
			statement function. The exists here returns true if it does indeed
			exist.
		*/
		if call_stmt, exists := statement_map[stmt_name]; exists {
			// Call the corresponding statement from the statement_map
			call_stmt()
		}
	}
}
