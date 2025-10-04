/*
The runtime module holds values that are created through the execution of a
script. You can think of this as the module that is home to values that might
be created that are specific to the execution of the script.

This module holds values that are specific to the language syntax itself.
*/
package values

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

/*
	The Token type houses information about a particular token and serves as an
	object that houses the relevant information. The structure of the token is
	as follows:

	FullLineOfCode [string]: the full line of code that the token is embedded
	in which is helpful for error reporting

	LineNumber [int]: the line number

	TokenPosition [string]: the starting position (column) of the token in
	question which is helpful for error reporting

	TokenValue [string]: the actual string value of the token

	TokenType [string]: the type of the token as a string

	NonCommentLineNumber [int]: the line counter for lines that aren't
	comments, that is, lines with statement calls
*/
type Token struct {
	FullLineOfCode string
	LineNumber int
	TokenPosition string
	TokenValue string
	TokenType string
	NonCommentLineNumber int
}

/*
	A helper function to make Token printing easier. It takes no parameters nor
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
	Do we have a shebang line? If so, set this to true. This is necessary for
	the minver statement 
*/
var SHEBANG_PRESENT bool = false

/*
	The TOKEN_TREE holds the tokens in a "tree" which is a glorified list of
	tokens.
*/
var TOKEN_TREE []byte

/*
	This houses the variables. This is prepopulated with the reserved
	variables. You will notice that some of these reserved variables are empty.
	They are constructed in the BuildReservedVariables() function. The reason
	that this is using string formatting throughout is to allow an easy change
	to the RESERVED_VARIABLE_PREFIX if need be.
*/

var VARIABLES = map[string]string{
	fmt.Sprintf("%sarch", RESERVED_VARIABLE_PREFIX): runtime.GOARCH,
	fmt.Sprintf("%scpu", RESERVED_VARIABLE_PREFIX): strconv.Itoa(runtime.NumCPU()),
	fmt.Sprintf("%sdate_dmy", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%sdate_ymd", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%shome", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%shostname", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%sipv4", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%sos", RESERVED_VARIABLE_PREFIX): runtime.GOOS,
	fmt.Sprintf("%stempdir", RESERVED_VARIABLE_PREFIX): os.TempDir(),
	fmt.Sprintf("%suser", RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf("%swd", RESERVED_VARIABLE_PREFIX): "",
}