/*
This module deals with the execute statement.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

/*
Execute a system command. Parameters include the tokens. Returns nothing.
*/
func ExecuteCommand(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("execute")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("execute")+" "+
				tools.ColouriseGreen("\"[command]\"")+". A common "+
				"issue here is excluding a command. An example of a working "+
				"statement might be "+tools.ColouriseCyan("execute")+
				tools.ColouriseGreen(" \"ls\"")+"."+"\n\nLine of Code: "+
				tools.ColouriseMagenta(full_loc),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Get the command and fix the string
	command := tools.FixStringCombined(tokens[2].TokenValue)

	/* Check if the -allowexec flag was passed to the app and if not, throw
	an error
	*/
	if !values.ALLOW_EXEC {
		investigator.Report(
			"You are unable to execute system commands. If you would like "+
				"to do so, you need to run with the "+
				tools.ColouriseYellow("-allowexec")+" flag.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...\n",
			tools.ColouriseBlue("Executing"),
			tools.ColouriseYellow(command),
		)
	}

	// Split the command into parts, needed below to account for arguments.
	cmd_split := strings.Split(command, " ")

	/* Get the output and an error from executing the command and capturing
	the output. Thanks to https://stackoverflow.com/a/23724092 for the
	argument passing here.
	*/
	output, err := exec.Command(cmd_split[0], cmd_split[1:]...).Output()
	// If the error isn't nil, throw an err
	if err != nil {
		investigator.Report(
			"The application "+tools.ColouriseYellow(command)+
				" was not found. Perhaps it was a typo?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// Output the results of the command
	fmt.Println(string(output))
}
