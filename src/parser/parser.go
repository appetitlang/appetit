/*
The parser module deals with tokenising the script and delegating to the
statement modules.
*/
package parser

import (
	"appetit/investigator"
	"appetit/statements"
	"appetit/tools"
	"appetit/values"
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"text/scanner"
)

/*
	Start executing commands in a script by passing the lines to the tokeniser
	and the delegator. The only parameter is the lines of the script. No
	returns.
*/
func Start(lines []string, dev_mode bool) {
	non_comment_line_count := 1
	// Loop over the lines
	for line := range lines {
		// Create a string version of the line
		line_as_string := string(lines[line])
		
		// Hold the length of the line
		line_length := len(line_as_string)

		/* It's possible that the line has no length (ie. a blank line) so we
			need to skip over them.
		*/
		if line_length > 0 {
			/* Pass each line to the delegator. Here, we start by checking to
			see what the first character of the line is to see if it's a
			SYMBOL_COMMENT. If it is not (ie. it's a line that requires
			parsing), we send the line to the tokeniser and the delegator.
			While the RemoveComments() function "removed" the comments, it only
			did so as far as it stripped away all the characters instead of the
			comment symbol so there are still comment line in that need to be
			accounted for here.
			*/
			if string(line_as_string[0]) != values.SYMBOL_COMMENT {
				/* Tokenise the line, adding one to the line number to account
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
			/* Pass an empty string. This is needed; if we don't pass this
				here, blank lines are skipped which results in line counts not
				being accurate. This is held in nothing as the tokenised value
				is irrelevant so we can dispense with this.
			*/
			_ = Tokenise(" ", line+1, -1)
		}
	}
}

/*
	Tokenise each line of the code a create a slice. Parameters include
	line_of_script (the non-tokenised line of the script that we are working to
	tokenise) and loc, the lines of code that we are working with where:
		- element 0 is the actual line of code;
		- element 1 is the counter used to track non-comment lines of code,
			those lines of code that are deemed significant (ie. those where
			the line has a statement call).
	Returns a slice of strings that represents a tokenised line where the first
	element is the line number and each subsequent element is a token in the
	line.
*/
func Tokenise(line_of_script string, loc ...int) []values.Token {
	// Set up a text scanner to tokenise the script
	var ls scanner.Scanner
	// Initialise against a line of the script passed
	tokeniser := ls.Init(strings.NewReader(line_of_script))
	/* Set up an error handler. Thanks to https://github.com/AOSPworking/
		pkg_repo_tool/blob/82e14e49258f/parser/parser.go
	*/

	// Catch errors with the tokeniser and handle them
	tokeniser.Error = func(s *scanner.Scanner, msg string) {
		// Call the investigators
		investigator.ReportTokeniserErrors(msg, loc[0])
	}

	// Create a slice that houses tokens for each part of the line of code
	var tokenised_line []values.Token
	// Create the line of code token
	loc_token := values.Token{
		FullLineOfCode: line_of_script,
		LineNumber: loc[0],
		TokenPosition: strconv.Itoa(tokeniser.Position.Column),
		TokenValue: "",
		TokenType: reflect.TypeOf(tokeniser.TokenText()).String(),
		NonCommentLineNumber: loc[1],
	}
	// Add the line of code to the tokenised line
	tokenised_line = append(tokenised_line, loc_token)

	// Loop over the tokens
	for tok := tokeniser.Scan(); tok != scanner.EOF; tok = tokeniser.Scan() {
		// Create a token
		token := values.Token{
			FullLineOfCode: line_of_script,
			LineNumber: loc[0],
			TokenPosition: strconv.Itoa(tokeniser.Position.Column),
			TokenValue: tokeniser.TokenText(),
			TokenType: reflect.TypeOf(tokeniser.TokenText()).String(),
			NonCommentLineNumber: loc[1],
		}
		/* Append the token to the slice of tokens that will be passed to
			statement modules in the Call()
		*/
		tokenised_line = append(tokenised_line, token)
	}

	// Append the tokens to the TOKEN_TREE
	values.TOKEN_TREE = append(values.TOKEN_TREE, tokenised_line...)

	// Return the tokenised line
	return tokenised_line
}

/*
	Start delegating lines of tokens to the functions that will do the work of
	actually executing functionality. Parameters include tokens, a slice of
	strings that contains the tokens. Returns nothing.
*/
func Call(tokens []values.Token) {
	/* This will catch empty lines as the slice will have a line number but no
		other elements. So, if the length is 1, we can assume that the only
		element is the line number and nothing else.
	*/
	// If there is more than one token
	if len(tokens) > 1 {

		// Get the statement name
		stmt_name := tokens[1].TokenValue
		// Create a map of statmements and their associated function calls
		statement_map := map[string]func(){
			"ask": func() { statements.Ask(tokens) },
			"copydirectory": func() { statements.CopyPath(tokens) },
			"copyfile": func() { statements.CopyFile(tokens) },
			"deletedirectory": func() { statements.DeletePath(tokens)},
			"deletefile": func() { statements.DeleteFile(tokens) },
			"download": func() { statements.Download(tokens) },
			"execute": func() { statements.ExecuteCommand(tokens) },
			"exit": func() { statements.Exit(tokens) },
			"makedirectory": func() { statements.CreatePath(tokens) },
			"makefile": func() { statements.MakeFile(tokens) },
			"minver": func() { statements.MinVer(tokens) },
			"movedirectory": func() { statements.MovePath(tokens) },
			"movefile": func() { statements.MoveFile(tokens) },
			"pause": func() { statements.Pause(tokens) },
			"set": func() { statements.Set(tokens) },
			"write": func() { statements.Writeln(tokens, false) },
			"writeln": func() { statements.Writeln(tokens, true) },
			"zipdirectory": func() { statements.ZipFromPath(tokens) },
			"zipfile": func() { statements.ZipFromFile(tokens) },
		}

		/* This may be unnecessarily repetitive but it helps to consolidate the
			list of statement names that are considered valid to one place:
			here.
			
			Here, we consolidate the keys from the statement_map map and set
			the values.STATEMENT_NAMES to the list of keys. Thanks to
			https://stackoverflow.com/q/21362950.
		*/
		values.STATEMENT_NAMES = slices.Collect(maps.Keys(statement_map))

		// First up, we need to check to see if there is a shebang line
		if tokens[0].FullLineOfCode[0:2] == "#!" {
			/* Set the SHEBANG_PRESENT value to true so that
				statements.MinVer() calls can ignore that the minver statement
				isn't on the first line and is, instead, on the second line
			*/
			values.SHEBANG_PRESENT = true
			// Break out of the function call
			return
		} else {
			valid_stmt := investigator.CheckIsStatement(
				tokens[1].FullLineOfCode,
				tokens[1].TokenValue,
			)

			if !valid_stmt {
				// Get the line of the script
				loc := strconv.Itoa(tokens[0].LineNumber)
				// Get full line of code
				full_loc := tokens[0].FullLineOfCode
				// Get the statement name
				stmt_name := tokens[1].TokenValue
				/* This executes if the statement doesn't exist (ie. exists is
					false). Report back that it doesn't exist with a list of valid
					statements.
				*/
				investigator.Report(
					"The statement passed - " + tools.ColouriseYellow(stmt_name) +
					" - is not a valid statement. Valid statements include " +
					values.ListStatements() + ".",
					loc,
					tokens[0].TokenPosition,
					full_loc,
				)
			}
		}

		/* If the statement is in the statement map, execute the corresponding
			statement function. The exists here returns true if it does indeed
			exist.
		*/
		if call_stmt, exists := statement_map[stmt_name]; exists {
			// Call the corresponding statement from the statement_map
			call_stmt()
		}
	}
}