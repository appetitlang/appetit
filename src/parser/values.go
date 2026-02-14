/*
This holds a collection of values that are referenced throughout the
parser. This includes:
- Language name and version values (as consts)
- Symbols and non-statement keywords used in the language
- Mode settings (eg. are we running in dev mode?)
- Token and variable information
- Some other values that are necessary to track across the parser
*/
package parser

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

/*
	These values hold language constants. When updating the language to a new
	version, LANG_VERSION needs to be changed.
*/
// Whether we are running a developer build right now
const LANG_DEV bool = true

// Language name for easy reference throughout
const LANG_NAME string = "Appetit"

/*
The versioning system here is a simple integer based system where each version
increment represents a new version. No semver or calver here.
*/
const LANG_VERSION int = 1

/*
This is set as the build date but this is changed with the Makefile. The linker
flags in the Makefile change this as a function of
*/
var BuildDate string = "-development"

/*
Code name for the language. This is named for capital cities in no particular
order or pattern. Currently, this isn't in use.
*/
const LANG_CODENAME = "Canberra"

/*
	This section houses symbols and conjoining words in statements and for the
	language. Anytime these need to be checked or worked with, they should be
	pulled from here.
*/

// The action symbol.
const SYMBOL_ACTION string = "to"

// Comment symbol.
const SYMBOL_COMMENT string = "-"

// The valid assignment operator
const SYMBOL_OPERATOR_ASSIGNMENT string = "="

// The reserved variable prefix
const SYMBOL_RESERVED_VARIABLE_PREFIX = "b_"

// Variable substitution symbol
const SYMBOL_VARIABLE_SUBSTITUTION = "#"

/*
	This section houses the state of any modes that we might be running, set
	via flags passed to the parameter.
*/
// Whether we will allow the execute statements
var MODE_ALLOW_EXEC bool = false

// Whether we are in developer mode
var MODE_DEV bool = false

// Whether we are verbose with our output
var MODE_VERBOSE bool = false

/*
Do we have a shebang line? If so, set this to true. This is necessary for
the minver statement
*/
var SHEBANG_PRESENT bool = false

// Valid statement names
var STATEMENT_NAMES []string

/*
The TOKEN_TREE holds the tokens in a "tree" which is a glorified list of
tokens.
*/
var TOKEN_TREE []Token

/*
	This houses the variables. This is prepopulated with the reserved
	variables. You will notice that some of these reserved variables are empty.
	They are constructed in the BuildReservedVariables() function. The reason
	that this is using string formatting throughout is to allow an easy change
	to the RESERVED_VARIABLE_PREFIX if need be.
*/

var VARIABLES = map[string]string{
	fmt.Sprintf(
		"%sarch",
		SYMBOL_RESERVED_VARIABLE_PREFIX): runtime.GOARCH,
	fmt.Sprintf(
		"%scpu",
		SYMBOL_RESERVED_VARIABLE_PREFIX): strconv.Itoa(runtime.NumCPU()),
	fmt.Sprintf(
		"%sdate_dmy",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%sdate_ymd",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%shome",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%shostname",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%sipv4",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%sos",
		SYMBOL_RESERVED_VARIABLE_PREFIX): runtime.GOOS,
	fmt.Sprintf(
		"%stempdir",
		SYMBOL_RESERVED_VARIABLE_PREFIX): os.TempDir(),
	fmt.Sprintf(
		"%stime",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%szone",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%suser",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
	fmt.Sprintf(
		"%swd",
		SYMBOL_RESERVED_VARIABLE_PREFIX): "",
}
