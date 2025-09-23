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
	Tokenise each line of the code a create a slice. Parameters include loc
	(the line of code that we are working with) and line_of_script (the
	non-tokenised line of the script that we are working to tokenise).
	Returns a slice of strings that represents a tokenised line where the first
	element is the line number and each subsequent element is a token in the
	line.
*/
func Tokeniser(loc int, line_of_script string) []values.Token {
	// Set up a text scanner to tokenise the script
	var ls scanner.Scanner
	// Initialise against a line of the script passed
	tokeniser := ls.Init(strings.NewReader(line_of_script))
	/* Set up an error handler. Thanks to https://github.com/AOSPworking/
		pkg_repo_tool/blob/82e14e49258f/parser/parser.go
	*/

	// Catch errors with the tokeniser and handle them
	tokeniser.Error = func(s *scanner.Scanner, msg string) {
		/*
		Set up an elaborate switch/case to capture any anticipated errors
		Parameters include the scanner and the message that gets reported.
		Returns nothing.
	*/
		switch msg {
		// Catch an unterminated literal
		case "literal not terminated":
			investigator.Report(
				"Your line of code has an incomplete string. Did you forget " +
				"an opening or closing quotation mark?\n\nSomething like " +
				"the following line of code will trigger this error:\n" +
				tools.ColouriseCyan("writeln ") +
				tools.ColouriseGreen("\"Hello world") +
				tools.ColouriseRed("_") + " <- (notice the lack of a " +
				"closing quotation mark here).",
				strconv.Itoa(loc),
				"n/a",
				"n/a",
			)
		// Catch an invalid char literal
		case "invalid char literal":
			investigator.Report(
				"Your line of code use single quotation marks instead of " +
				"the required double quotation marks. See the example:\n\n" +
				tools.ColouriseCyan("writeln ") +
				tools.ColouriseGreen("'Hello world'") + " <- (notice the" +
				" lack of double quotation marks here).",
				strconv.Itoa(loc),
				"n/a",
				"n/a",
			)
		default:
			/* Report everything else in their "Go form." It is hoped that, some
				day, this will not need to exist
			*/
			investigator.Report(
				msg,
				strconv.Itoa(loc),
				"n/a",
				"n/a",
			)
		}
	}

	// Create a slice that houses tokens for each part of the line of code
	var tokenised_line []values.Token
	// Create the line of code token
	loc_token := values.Token{
		FullLineOfCode: line_of_script,
		LineNumber: loc,
		TokenPosition: strconv.Itoa(tokeniser.Position.Column),
		TokenValue: "",
		TokenType: reflect.TypeOf(tokeniser.TokenText()).String(),
	}
	// Add the line of code to the tokenised line
	tokenised_line = append(tokenised_line, loc_token)

	// Loop over the tokens
	for tok := tokeniser.Scan(); tok != scanner.EOF; tok = tokeniser.Scan() {
		// Create a token
		token := values.Token{
			FullLineOfCode: line_of_script,
			LineNumber: loc,
			TokenPosition: strconv.Itoa(tokeniser.Position.Column),
			TokenValue: tokeniser.TokenText(),
			TokenType: reflect.TypeOf(tokeniser.TokenText()).String(),
		}
		/* Append the token to the slice of tokens that will be passed to
			statement modules in the Delegator()
		*/
		tokenised_line = append(tokenised_line, token)
	}

	// Return the tokenised line
	return tokenised_line
}

/*
	Start delegating lines of tokens to the functions that will do the work of
	actually executing functionality. Parameters include tokens, a slice of
	strings that contains the tokens. Returns nothing.
*/
func Delegator(tokens []values.Token) {
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