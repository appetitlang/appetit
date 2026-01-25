/*
The reserved module holds values that are needed across the language. It also
offers support for checking whether values and operations are conformant with
the language.

This module holds values that are specific to the language syntax itself.
*/
package values

import (
	"appetit/tools"
	"fmt"
	"net"
	"os"
	"os/user"
	"slices"
	"strings"
	"syscall"
	"time"
)

// Whether we are running a developer build right now
const LANG_DEV bool = true

// Language name for easy reference throughout
const LANG_NAME string = "Appetit"

/*
Version. The versioning system here is a simple integer based system where
each version increment represents a new version. No semver here.
*/
const LANG_VERSION int = 1

// Whether we will allow the execute statements
var ALLOW_EXEC bool = false

// Whether we are verbose with our output
var MODE_VERBOSE bool = false

// Whether we are in developer mode
var MODE_DEV bool = false

// The valid assignment operator
var OPERATOR_ASSIGNMENT string = "="

// The reserved variable prefix
var RESERVED_VARIABLE_PREFIX = "b_"

// Valid statement names
/*var STATEMENT_NAMES = []string{
	"ask",
	"copydirectory",
	"copyfile",
	"deletedirectory",
	"deletefile",
	"download",
	"execute",
	"makedirectory",
	"makefile",
	"minver",
	"movedirectory",
	"movefile",
	"pause",
	"set",
	"write",
	"writeln",
	"zipdirectory",
	"zipfile",
}*/
var STATEMENT_NAMES []string

/*
	 The action symbol. This is used in placements where we don't need an
		assignment operator but need something to split something where something
		is happening to something.
*/
var SYMBOL_ACTION string = "to"

// Comment symbol.
var SYMBOL_COMMENT string = "-"

// Variable substitution symbol
var SYMBOL_VARIABLE_SUBSTITUTION = "#"

/*
Create any values for built in reserved variables that require building.
This addresses the empty ones in the VARIABLES map. No parameters and no
returns.
*/
func BuildReservedVariables() {

	cur_user, cur_user_error := user.Current()
	if cur_user_error != nil {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"user"] = ""
	} else {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"user"] = cur_user.Username
	}

	// Create a date
	date := time.Now()

	// Get the date in dd-mm-yyyy format
	VARIABLES[RESERVED_VARIABLE_PREFIX+"date_dmy"] = fmt.Sprintf(
		"%d-%d-%d", date.Day(), date.Month(), date.Year(),
	)

	// Get the date in yyyy-mm-dd format
	VARIABLES[RESERVED_VARIABLE_PREFIX+"date_ymd"] = fmt.Sprintf(
		"%d-%d-%d", date.Year(), date.Month(), date.Day(),
	)

	// Get the current time
	time_now := time.Now()

	// Get the time in hh-mm-ss in 24 hour format
	VARIABLES[RESERVED_VARIABLE_PREFIX+"time"] = fmt.Sprintf(
		"%d-%d-%d", time_now.Hour(), time_now.Minute(), time_now.Second(),
	)

	// Get the time zone, ignoring the offset as this isn't needed
	time_zone, _ := time_now.Zone()

	// Get the timezone
	VARIABLES[RESERVED_VARIABLE_PREFIX+"zone"] = time_zone

	// Get the hostname
	host, err := os.Hostname()
	// If there's no error, set the b_host to the hostname.
	if err == nil {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"hostname"] = host
	}
	// Get the user home directory
	home, err := os.UserHomeDir()
	// If there's no error, set the b_home to the home directory.
	if err == nil {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"home"] = home
	}

	// Get the working directory
	wd, err := syscall.Getwd()
	// If there's no error, set the b_wd to the working directory.
	if err == nil {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"wd"] = wd
	}

	/* House the ip address. This now improves on the old format which depended
	both on an external network connection and the stability of Google's
	public DNS service which is not something I want to depend on.
	*/
	// Get the interfaces
	ipaddrs, ipaddrs_err := net.InterfaceAddrs()
	// If we can't, abandon ship and save n/a to the ipv4 reserved variable
	if ipaddrs_err != nil {
		VARIABLES[RESERVED_VARIABLE_PREFIX+"ipv4"] = "n/a"
	}

	// Iterate over the addresses
	for _, ipv4_addresses := range ipaddrs {
		// If the address is an ip address and is not a loopback address...
		if ip, ip_ok := ipv4_addresses.(*net.IPNet); ip_ok && !ip.IP.IsLoopback() {
			// If converting it to an IPv4 address doesn't yield an error...
			if ip.IP.To4() != nil {
				// Set the IPv4 address reserved variable
				VARIABLES[RESERVED_VARIABLE_PREFIX+"ipv4"] = ip.IP.String()
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
		if vars[0:len(RESERVED_VARIABLE_PREFIX)] == RESERVED_VARIABLE_PREFIX {
			/* Append the reserved variable to the list of them with a
			colourised version
			*/
			reserved_var = append(reserved_var, tools.ColouriseMagenta(vars))
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
			statement_names, tools.ColouriseMagenta(STATEMENT_NAMES[stmt]),
		)
	}
	// Sort the list of statement names
	slices.Sort(statement_names)
	// Return a joined version of this list
	return strings.Join(statement_names, ", ")
}
