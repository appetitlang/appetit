/*
This module deals with file related statements.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Copy a file from an origin to a destination. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate.
Returns nothing. Thanks to https://www.kelche.co/blog/go/golang-file-
handling/.
*/
func CopyFile(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("copyfile")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("copyfile")+" "+
				tools.ColouriseGreen("\"[path]\"")+" to "+
				tools.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				tools.ColouriseMagenta(values.SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				tools.ColouriseCyan("copyfile")+
				tools.ColouriseGreen(" \"test.txt\"")+" to "+
				tools.ColouriseGreen(" \"test_new.txt\"")+"\n\nLine of "+
				"Code: "+tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := investigator.CheckAction(loc, action)
	if action_error != nil {
		investigator.Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix the destination string
	destination := tools.FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	/* Split the origin by the os path separator so that we can get the
	file name in case we need to append it
	*/
	file_path_parts := strings.Split(source, string(os.PathSeparator))
	// Get the filename from the parts of the file path
	filename := file_path_parts[len(file_path_parts)-1]
	// Get the last character of the destination
	end_of_dest := string(destination[len(destination)-1])
	// If the last character is the path separator (ie. a path was passed)...
	if end_of_dest == string(os.PathSeparator) {
		// Append the file
		destination = destination + filename
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			tools.ColouriseBlue("Copying"),
			tools.ColouriseGreen(source),
			tools.ColouriseGreen(destination),
		)
	}

	source_file, err := os.Open(source)
	if err != nil {
		investigator.Report(
			"Can't open "+tools.ColouriseYellow(source)+
				"! Are you sure that the file exists?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	destination_file, err := os.Create(destination)
	if err != nil {
		investigator.Report(
			"The destination - "+tools.ColouriseYellow(destination)+
				" - is invalid. Are you source that the destination exists? If "+
				"you're trying to copy to a directory, make sure to put in a "+
				"trailing "+tools.ColouriseYellow(string(os.PathSeparator)),
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	bytes, err := io.Copy(destination_file, source_file)
	if err != nil {
		investigator.Report(
			"There was an error doing the file copy/move. Check to ensure that "+
				"source - "+source+" - and the destination - "+destination+
				" - are valid.",
			loc,
			"n/a",
			full_loc,
		)
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			"done! "+
				tools.ColouriseMagenta(
					"[%s bytes written]\n",
				),
			strconv.FormatInt(bytes, 10),
		)
	}
}

/*
Move a file from an origin to a destination. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate.
Returns nothing.
*/
func MoveFile(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("movefile")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("movefile")+" "+
				tools.ColouriseGreen("\"[path]\"")+" to "+
				tools.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				tools.ColouriseMagenta(values.SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				tools.ColouriseCyan("movefile")+
				tools.ColouriseGreen(" \"test.txt\"")+" to "+
				tools.ColouriseGreen(" \"test_new.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := investigator.CheckAction(loc, action)
	if action_error != nil {
		investigator.Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix the destination string
	destination := tools.FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	/* Split the origin by the os path separator so that we can get the
	file name in case we need to append it
	*/
	file_path_parts := strings.Split(source, string(os.PathSeparator))
	// Get the filename from the parts of the file path
	filename := file_path_parts[len(file_path_parts)-1]
	// Get the last character of the destination
	end_of_dest := string(destination[len(destination)-1])
	// If the last character is the path separator (ie. a path was passed)...
	if end_of_dest == string(os.PathSeparator) {
		// Append the file
		destination = destination + filename
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			tools.ColouriseBlue("Moving"),
			tools.ColouriseGreen(source),
			tools.ColouriseGreen(destination),
		)
	}

	/* Thanks to https://www.geeksforgeeks.org/how-to-rename-and-move-a-file-
	in-golang/
	*/
	move_err := os.Rename(source, destination)
	if move_err != nil {
		/* If there's an error renaming the file, copy it instead and delete
		the original
		*/
		CopyFile(tokens)
		remove_err := os.Remove(source)
		if remove_err != nil {
			investigator.Report(
				"There was an error removing the source file: "+
					tools.ColouriseYellow(source)+". It will be worth "+
					"trying to remove it manually.",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}
	}

	if values.MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
Delete a file. The tokens are passed to get the file that will be deleted
and the full line of code is passed for error reporting. Returns nothing.
*/
func DeleteFile(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("deletefile")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("deletefile")+" "+
				tools.ColouriseGreen("\"[path]\"")+". An example of a working "+
				"version might be "+tools.ColouriseCyan("deletefile")+
				tools.ColouriseGreen(" \"test.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)
	// Check to see if the file exists
	file_exists := investigator.FileExists(source)

	// If the file doesn't exist, error out
	if !file_exists {
		// Report the error
		investigator.Report(
			tools.ColouriseMagenta(source)+" does not exist.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
		// If it does exist, remove it
	} else {
		// If we're in verbose mode, report back some more info
		if values.MODE_VERBOSE {
			fmt.Printf(":: Deleting %s...", tools.ColouriseMagenta(source))
		}
		err := os.Remove(source)
		if err != nil {
			// Get some information on the source file if we can
			info, info_err := os.Lstat(source)
			/* If there was an error getting some info on the file, we're out
			of options for being specific on the issue so we communicate
			that
			*/
			if info_err != nil {
				investigator.Report(
					"So, I'm having difficulties removing the file and even "+
						"getting some information on the file to help you "+
						"understand why.",
					loc,
					tokens[2].TokenPosition,
					full_loc,
				)
			}
			// Report back some info about the permissions on the file
			investigator.Report(
				"There was an error deleting the file: "+
					tools.ColouriseMagenta(source)+". It looks like the "+
					"permissions on the file are "+info.Mode().Perm().String()+
					". Check to make sure that you have the right permissions "+
					"to delete the file.",
				loc,
				tokens[2].TokenPosition,
				full_loc,
			)
		}
		// If we're in verbose mode, report back some more info
		if values.MODE_VERBOSE {
			fmt.Println("done!")
		}

	}

}

/*
Create a file. The tokens are passed to get the file that will be deleted
and the full line of code is passed for error reporting. Returns nothing.
*/
func MakeFile(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("makefile")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("create")+" "+
				tools.ColouriseGreen("\"[path]\"")+". An example of a working "+
				"version might be "+tools.ColouriseCyan("makefile")+
				tools.ColouriseGreen(" \"test.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the file name string
	file_name := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	file_name = VariableTemplater(file_name)
	// Check to see if the file exists
	file_exists := investigator.FileExists(file_name)

	// If the file exists already exist, error out
	if file_exists {
		// Report the error
		investigator.Report(
			tools.ColouriseMagenta(file_name)+" exists already.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// If we're in verbose mode, report back some more info
	if values.MODE_VERBOSE {
		fmt.Printf(":: Making %s...", tools.ColouriseMagenta(file_name))
	}
	// Create the file
	_, create_err := os.Create(file_name)
	// If there is an issue with creating the file...
	if create_err != nil {
		// Report the error
		investigator.Report(
			tools.ColouriseMagenta(file_name)+" could not be created: "+
				create_err.Error(),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// If we're in verbose mode, report back some more info
	if values.MODE_VERBOSE {
		fmt.Println("done!")
	}
}
