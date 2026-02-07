/*
This module deals with file related statements: copyfile, deletefile, makefile,
and movefile.
*/
package parser

import (
	"appetit/utils"
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
func CopyFile(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("copyfile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("copyfile")+" "+
				utils.ColouriseGreen("\"[path]\"")+" to "+
				utils.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				utils.ColouriseMagenta(SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				utils.ColouriseCyan("copyfile")+
				utils.ColouriseGreen(" \"test.txt\"")+" to "+
				utils.ColouriseGreen(" \"test_new.txt\"")+"\n\nLine of "+
				"Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	if action_error != nil {
		Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix the destination string
	destination := FixStringCombined(tokens[4].TokenValue)
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

	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Copying"),
			utils.ColouriseGreen(source),
			utils.ColouriseGreen(destination),
		)
	}

	source_file, err := os.Open(source)
	if err != nil {
		Report(
			"Can't open "+utils.ColouriseYellow(source)+
				"! Are you sure that the file exists?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	destination_file, err := os.Create(destination)
	if err != nil {
		Report(
			"The destination - "+utils.ColouriseYellow(destination)+
				" - is invalid. Are you source that the destination exists? If "+
				"you're trying to copy to a directory, make sure to put in a "+
				"trailing "+utils.ColouriseYellow(string(os.PathSeparator)),
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	bytes, err := io.Copy(destination_file, source_file)
	if err != nil {
		Report(
			"There was an error doing the file copy/move. Check to ensure that "+
				"source - "+source+" - and the destination - "+destination+
				" - are valid.",
			loc,
			"n/a",
			full_loc,
		)
	}

	if MODE_VERBOSE {
		fmt.Printf(
			"done! "+
				utils.ColouriseMagenta(
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
func MoveFile(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("movefile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("movefile")+" "+
				utils.ColouriseGreen("\"[path]\"")+" to "+
				utils.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				utils.ColouriseMagenta(SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				utils.ColouriseCyan("movefile")+
				utils.ColouriseGreen(" \"test.txt\"")+" to "+
				utils.ColouriseGreen(" \"test_new.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	if action_error != nil {
		Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix the destination string
	destination := FixStringCombined(tokens[4].TokenValue)
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

	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Moving"),
			utils.ColouriseGreen(source),
			utils.ColouriseGreen(destination),
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
			Report(
				"There was an error removing the source file: "+
					utils.ColouriseYellow(source)+". It will be worth "+
					"trying to remove it manually.",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}
	}

	if MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
Delete a file. The tokens are passed to get the file that will be deleted
and the full line of code is passed for error reporting. Returns nothing.
*/
func DeleteFile(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("deletefile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("deletefile")+" "+
				utils.ColouriseGreen("\"[path]\"")+". An example of a working "+
				"version might be "+utils.ColouriseCyan("deletefile")+
				utils.ColouriseGreen(" \"test.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the source string
	source := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)
	// Check to see if the file exists
	file_exists := CheckFileExists(source)

	// If the file doesn't exist, error out
	if !file_exists {
		// Report the error
		Report(
			utils.ColouriseMagenta(source)+" does not exist.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
		// If it does exist, remove it
	} else {
		// If we're in verbose mode, report back some more info
		if MODE_VERBOSE {
			fmt.Printf(":: Deleting %s...", utils.ColouriseMagenta(source))
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
				Report(
					"So, I'm having difficulties removing the file and even "+
						"getting some information on the file to help you "+
						"understand why.",
					loc,
					tokens[2].TokenPosition,
					full_loc,
				)
			}
			// Report back some info about the permissions on the file
			Report(
				"There was an error deleting the file: "+
					utils.ColouriseMagenta(source)+". It looks like the "+
					"permissions on the file are "+info.Mode().Perm().String()+
					". Check to make sure that you have the right permissions "+
					"to delete the file.",
				loc,
				tokens[2].TokenPosition,
				full_loc,
			)
		}
		// If we're in verbose mode, report back some more info
		if MODE_VERBOSE {
			fmt.Println("done!")
		}

	}

}

/*
Create a file. The tokens are passed to get the file that will be deleted
and the full line of code is passed for error reporting. Returns nothing.
*/
func MakeFile(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("makefile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("create")+" "+
				utils.ColouriseGreen("\"[path]\"")+". An example of a working "+
				"version might be "+utils.ColouriseCyan("makefile")+
				utils.ColouriseGreen(" \"test.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the file name string
	file_name := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	file_name = VariableTemplater(file_name)
	// Check to see if the file exists
	file_exists := CheckFileExists(file_name)

	// If the file exists already exist, error out
	if file_exists {
		// Report the error
		Report(
			utils.ColouriseMagenta(file_name)+" exists already.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// If we're in verbose mode, report back some more info
	if MODE_VERBOSE {
		fmt.Printf(":: Making %s...", utils.ColouriseMagenta(file_name))
	}
	// Create the file
	_, create_err := os.Create(file_name)
	// If there is an issue with creating the file...
	if create_err != nil {
		// Report the error
		Report(
			utils.ColouriseMagenta(file_name)+" could not be created: "+
				create_err.Error(),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// If we're in verbose mode, report back some more info
	if MODE_VERBOSE {
		fmt.Println("done!")
	}
}
