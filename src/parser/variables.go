/*
This deals with variable related functions. This includes working with values
in statements that require 'calculation' inclusive of variables.
*/
package parser

import (
	"appetit/utils"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"time"
)

/*
Check to see if a variable is set. This is a helper function for the
BuildReservedVariables() function. Takes one parameter - the variable to check
and returns two bools: whether the variable exists and whether the variable has
a value. This pair of booleans can only return three combinations:
 1. true, false - the variable exists but has no assigned value (this would
    include placeholders for reserved variables).
 2. true, true - the variable exists and has an assigned value.
 3. false, false - the variable doesn't exist and therefore, doesn't have an
    assigned value

There is no "false, true" pair here - you can't have a variable that doesn't
exist have a value.
*/
func CheckVariableExistence(var_name string) (bool, bool) {
	// This conditional will be met if the VARIABLE exists
	if value, ok := VARIABLES[var_name]; ok {
		// If the value is nothing, we have a variable but no value
		if value == "" {
			return true, false
			// If the variable is present and has a value, return true for both
		} else {
			return true, true
		}
		// If we've gotten here, we can assume that the variable does not exist
	} else {
		return false, false
	}
}

/*
Create any values for built in reserved variables that require building.
This addresses the empty ones in the VARIABLES map and updates those that
require specific values for each statement call. This does not update each
variable's value each call, however, as checks are made to see if those values
that won't change during runtime already have values. No parameters and no
returns.
*/
func BuildReservedVariables() {
	// Check to see if the user reserved variable has a value
	_, cur_user_value := CheckVariableExistence(
		SYMBOL_RESERVED_VARIABLE_PREFIX + "user")
	/*
		If it doesn't, add one. If it doesn't, this condition won't be met so
		the code will continue along.
	*/
	if !cur_user_value {
		// Get the current user
		cur_user, cur_user_error := user.Current()
		// Assuming that there was no issue getting the current user, assign it
		if cur_user_error != nil {
			VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"user"] = ""
		} else {
			VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"user"] = cur_user.Username
		}
	}

	// Create a date
	date_time := time.Now()
	// Get the day
	date_day := fmt.Sprintf("%02d", date_time.Day())
	// Get the month
	date_month := fmt.Sprintf("%02d", date_time.Month())
	// Get the year
	date_year := fmt.Sprintf("%02d", date_time.Year())
	// Get the hour
	time_hour := fmt.Sprintf("%02d", date_time.Hour())
	// Get the minute
	time_minute := fmt.Sprintf("%02d", date_time.Minute())
	// Get the second
	time_seconds := fmt.Sprintf("%02d", date_time.Second())

	/*
		Get the date in dd-mm-yyyy format. This should be re-generated every
		run of Call().
	*/
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_dmy"] = fmt.Sprintf(
		"%s-%s-%s", date_day, date_month, date_year,
	)

	/*
		Get the date in yyyy-mm-dd format. This should be re-generated every
		run of Call().
	*/
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_ymd"] = fmt.Sprintf(
		"%s-%s-%s", date_year, date_month, date_day,
	)

	// Set the date_day
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_day"] = fmt.Sprintf(
		"%s", date_day,
	)

	// Set the date_month
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_month"] = fmt.Sprintf(
		"%s", date_month,
	)

	// Set the date_year
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"date_year"] = fmt.Sprintf(
		"%s", date_year,
	)

	/*
		Get the time in hh-mm-ss in 24 hour format. This should be re-generated
		every run of Call().
	*/
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"time_full"] = fmt.Sprintf(
		"%s-%s-%s", time_hour, time_minute, time_seconds,
	)

	// Set the date_day
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"time_hour"] = fmt.Sprintf(
		"%s", time_hour,
	)

	// Set the date_month
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"time_minute"] = fmt.Sprintf(
		"%s", time_minute,
	)

	// Set the date_year
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"time_seconds"] = fmt.Sprintf(
		"%s", time_seconds,
	)

	/*
		Create the logstamp by combining the date_ymd and time reserved
		variables.
	*/
	VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"logstamp"] = fmt.Sprintf(
		"%s/%s/%s, %s:%s:%s",
		date_year,
		date_month,
		date_day,
		time_hour,
		time_minute,
		time_seconds,
	)

	/*
		Check to see if we've got the b_scriptname_full which also serves to
		check if we have the b_scriptname_only variable set.
	*/
	_, cur_script_names := CheckVariableExistence(
		SYMBOL_RESERVED_VARIABLE_PREFIX + "scriptname_full")

	if !cur_script_names {
		/*
			This needs to be set here as it can't be set in the creation of the
			VARIABLE map because that map is created before.
		*/
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"scriptname_full"] = SCRIPT_NAME
		// Get just the file name
		_, name_only := filepath.Split(SCRIPT_NAME)
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"scriptname_only"] = name_only
	}

	_, cur_time_zone := CheckVariableExistence(
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"zone"],
	)

	if !cur_time_zone {
		/*
			Get the time zone, ignoring the offset as this isn't needed. This
			should be re-generated every run of Call(). While it is unlikely
			that someone will move between time zones or have the machine's
			time zone change during execution, it is possible.
		*/
		// Get the timezone
		time_zone, _ := date_time.Zone()

		// Get the timezone
		VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"zone"] = time_zone
	}

	/*
		Check to see if the hostname is set. This doesn't need to be
		re-generated each call of the function.
	*/
	_, host_value := CheckVariableExistence(
		SYMBOL_RESERVED_VARIABLE_PREFIX + "hostname")
	// Get the hostname
	if !host_value {
		host, err := os.Hostname()
		// If there's no error, set the b_host to the hostname.
		if err == nil {
			VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"hostname"] = host
		}
	}

	/*
		Check if the user's home directory is already set. This won't need to
		be re-set each time as this won't change.
	*/
	_, home_dir_value := CheckVariableExistence(
		SYMBOL_RESERVED_VARIABLE_PREFIX + "home")

	if !home_dir_value {
		// Get the user home directory
		home, err := os.UserHomeDir()
		// If there's no error, set the b_home to the home directory.
		if err == nil {
			VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"home"] = home
		}
	}

	/*
		Check to see if the working directory is set. For now, this doesn't
		need to be updated; while future statements may allow for a change in
		the working directory, this isn't true yet so we don't need to update
		this each time.
	*/
	_, work_dir_value := CheckVariableExistence(
		SYMBOL_RESERVED_VARIABLE_PREFIX + "wd")

	if !work_dir_value {
		// Get the working directory
		wd, err := syscall.Getwd()
		// If there's no error, set the b_wd to the working directory.
		if err == nil {
			VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"wd"] = wd
		}
	}

	/*
		House the ip address. This loops over addresses attached to network
		interfaces. This is rather "fragile" in that it's very possible that
		this block of code will not yield an expected result if and where there
		are more than one network devices in use.
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
		if ip, ip_ok := ipv4_addresses.(*net.IPNet); ip_ok &&
			!ip.IP.IsLoopback() {
			// If converting it to an IPv4 address doesn't yield an error...
			if ip.IP.To4() != nil {
				// Get the IP address as a string
				ipv4_addr := ip.IP.String()
				// Set the IPv4 address reserved variable
				VARIABLES[SYMBOL_RESERVED_VARIABLE_PREFIX+"ipv4"] = ipv4_addr
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
Replace variables inside of a string. In short, this takes a string such as
"Your name is #name" and converts it to "Your name is Appetit" (where #name
is "Appetit"). Parameters include the input line of code to fix. Returns a
templated string where variables have been fixed.
*/
func VariableTemplater(input string) string {
	// For each key-value pair in the map of variables
	for key, value := range VARIABLES {
		// Get the string value of the variable
		value = string(value)
		/*
			Replace the value in the string if the value is found in the
			string prepended by the variable replacement symbol. The
			SYMBOL_VARIABLE_SUBSTITUTION+key checks that the variable symbol
			precedes the key to ensure that words that happen to have similar
			names don't also get replaced. In other words, '#name' is
			completely different than 'name'.
		*/
		input = strings.ReplaceAll(
			input,
			SYMBOL_VARIABLE_SUBSTITUTION+key,
			value,
		)
	}
	// Return the substituted string
	return input
}
