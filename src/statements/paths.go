/*
The statements package controls the execution of actual statements.

This module deals with path related statements.
*/
package statements

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
	This is the function that walks the source directory for a copy path call
	and does the work of copying files. This is called from CopyPath().
	Thanks to https://xojoc.pw/blog/golang-file-tree-traversal
*/
func CopyPathWalker(token_info map[string]string) filepath.WalkFunc {

	// Extract values from the token_info passed to the function
	source := token_info["source"]
	dest := token_info["destination"]
	loc := token_info["loc"]
	full_loc := token_info["full_loc"]

	// Get the list of directories that make up the source path
	list_of_src_dirs := strings.Split(source, string(os.PathSeparator))
	/* Get the folder name. The index is the length minus 2 given that the path
		seperator is used. So, something like /Users/user/Downloads/test/ would
		split into _ [0] Users [1] user [2] Downloads [3] test [4] _ [5] so the
		length is six but we want element four.
	*/
	source_directory := list_of_src_dirs[len(list_of_src_dirs)-2] +
						string(os.PathSeparator)

	dest = dest + source_directory

	return func(path string, info fs.FileInfo, err error) error {
		// Get the relative path of the file that we are looking at here
		relative_path := strings.TrimPrefix(path, source)
		/* If the object being traversed is a directory, create that and
			any parents as need be
		*/
		if info.IsDir() {
			/* If verbose mode is set, notify the user that we are making a
				directory
			*/
			if values.MODE_VERBOSE {
				fmt.Println(
					":: Making " + tools.ColouriseGreen(relative_path) +
					"...",
				)
			}
			// Make the path with some sensible permissions
			os.MkdirAll(dest + relative_path, 0750)
		} else {
			// Get the name of the file to copy
			file_to_copy := source + relative_path
			// Open the source file
			source_file, source_err := os.Open(file_to_copy)
			// If there is an error opening the source file, report that
			if source_err != nil {
				investigator.Report(
					"Can't open " + tools.ColouriseYellow(file_to_copy) +
					". Perhaps you don't have read permissions? " +
					source_err.Error(),
					loc,
					"n/a",
					full_loc,
				)
			}
			// Establish where we are creating the files
			create_path := dest + relative_path
			// Create the new file
			create, create_err := os.Create(create_path)
			// If there was an error in creating the new file, report that
			if create_err != nil {
				investigator.Report(
					"Couldn't create the file in " +
					tools.ColouriseYellow(create_path) + ". Check that " +
					"you have write permissions and/or that there is " +
					"enough space available for you to copy the file(s) " +
					"over. " + create_err.Error(),
					loc,
					"n/a",
					full_loc,
				)
			}
			/* If verbose mode is set, note that we are copying a file and
				report back the file size
			*/
			if values.MODE_VERBOSE {
				fmt.Printf(
					"    :: Copying %s %s...",
					tools.ColouriseGreen(info.Name()),
					tools.ColouriseMagenta(
						"[" + strconv.FormatInt(info.Size(), 10) +
						" bytes]",
					),
				)
			}
			
			// Copy the file itself
			_, copy_err := io.Copy(create, source_file)
			// If there's an error, report it
			if copy_err != nil {
				investigator.Report(
					"There was an error doing the copy for " +
					source_file.Name(),
					loc,
					"n/a",
					full_loc,
				)
			}
			// If verbose mode is set, note that we are done copying the file
			if values.MODE_VERBOSE {
				fmt.Println("done.")
			}

		}
		/* We return no error here as it is assumed that any errors are handled
			above
		*/
		return nil
	}
}

/*
	Copy a directory from one place to another. The parameters are the
	conventional set of tokens. Returns nothing.
*/
func CopyPath(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, token_err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if token_err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("copydirectory") +
			"statement needs to follow the form " +
			tools.ColouriseCyan("copydirectory") +
			tools.ColouriseGreen(" \"[path]\"") + " to " +
			tools.ColouriseGreen("\"[path]\"") + ". A common issue " +
			"is the  use of an inappropriate action symbol (" +
			tools.ColouriseMagenta(values.SYMBOL_ACTION) + "). An " + 
			"example of a working version might be " +
			tools.ColouriseCyan("copyfcopydirectoryile") +
			tools.ColouriseGreen(" \"test_dir\"") + " to " +
			tools.ColouriseGreen(" \"new_dir\""),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Get the source folder to copy and fix the string where need be
	source_path := tools.FixStringCombined(tokens[2].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	source_path = tools.FixPathSeperators(source_path)
	/* Get a templated value, that is, a variable where values have
		been substituted.
	*/
	source_path = VariableTemplater(source_path)

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
	// Get the source folder to copy and fix the string where need be
	dest_path := tools.FixStringCombined(tokens[4].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	dest_path = tools.FixPathSeperators(dest_path)
	/* Get a templated value, that is, a variable where values have
		been substituted.
	*/
	dest_path = VariableTemplater(dest_path)

	// Set up a map of values to be passed to the file walker
	walker_values := make(map[string]string)
	walker_values["source"] = source_path
	walker_values["destination"] = dest_path
	walker_values["loc"] = loc
	walker_values["full_loc"] = loc

	// Walk the files are start copying
	filepath.Walk(source_path, CopyPathWalker(walker_values))
}

/*
	Move a directory. The parameters are the conventional set of tokens.
	Returns nothing.
*/
func MovePath(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("movedirectory") + " statement " +
			"needs to follow the form " +
			tools.ColouriseCyan("movedirectory") + " " +
			tools.ColouriseYellow("\"[path]\"") + ". A common error here " +
			"is trying to concatenate multiple values into one statement " +
			"call here. An example of a working version might be " + 
			tools.ColouriseCyan("movedirectory ") +
			tools.ColouriseGreen("\"test_dir\"") + " to " +
			tools.ColouriseGreen("\"actual_dir\""),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Get the source folder to copy and fix the strings
	old_path := tools.FixStringCombined(tokens[2].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	old_path = tools.FixPathSeperators(old_path)
	/* Get a templated value, that is, a variable where values have
		been substituted.
	*/
	old_path = VariableTemplater(old_path)

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
	// Get the destination folder to copy and fix the strings
	new_path := tools.FixStringCombined(tokens[4].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	new_path = tools.FixPathSeperators(new_path)
	/* Get a templated value, that is, a variable where values have
		been substituted.
	*/
	new_path = VariableTemplater(new_path)

	if values.MODE_VERBOSE {
		fmt.Printf(
				":: %s %s to %s...",
				tools.ColouriseBlue("Moving"),
				tools.ColouriseGreen(old_path),
				tools.ColouriseGreen(new_path),
			)
	}

	/* Thanks to https://www.geeksforgeeks.org/how-to-rename-and-move-a-file-
		in-golang/
	*/
	move_err := os.Rename(old_path, new_path)
	if move_err != nil {
		/* Trying to move files across partitions often yields an error. In
			light of that, we'll just try to move things across by actually
			copying them.
		*/
		// Give copying a go here instead.
		CopyPath(tokens)
		/*investigator.Report(
			"There was an error moving the directory. Check to ensure that " +
			"the source - " + tools.ColouriseYellow(old_path) + " - and the " +
			"destination - " + tools.ColouriseYellow(new_path) + " - " +
			"are valid locations and that " + tools.ColouriseYellow(new_path) +
			" doesn't already exist.",
			loc,
			"n/a",
			full_loc,
		)*/
	}

	if values.MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
	Delete a directory. The parameters are the conventional set of tokens.
	Returns nothing.
*/
func DeletePath(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("deletedirectory") + " statement " +
			"needs to follow the form " +
			tools.ColouriseCyan("deletedirectory") + " " +
			tools.ColouriseYellow("\"[path]\"") + ". A common error here " +
			"is trying to concatenate multiple values into one statement " +
			"call here. An example of a working version might be " + 
			tools.ColouriseCyan("deletedirectory ") +
			tools.ColouriseGreen("\"test_dir\"") + ".",
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the path string
	path := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
		substituted
	*/
	path = VariableTemplater(path)

	// If verbose mode is set, print out what's happening
	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...",
			tools.ColouriseBlue("Deleting"),
			tools.ColouriseGreen(path),
		)
	}

	remove_err := os.RemoveAll(path)
	if remove_err != nil {
		investigator.Report(
			"There was an error removing " +
			tools.ColouriseMagenta(path) + ". The path does not exist.",
			loc,
			"n/a",
			full_loc,
		)
	}

	if values.MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
	Make a directory. The tokens are passed to get the file that will be moved.
	Returns nothing.
*/
func CreatePath(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The " + tools.ColouriseCyan("makedirectory") + " statement " +
			"needs to follow the form " +
			tools.ColouriseCyan("makedirectory") + " " +
			tools.ColouriseYellow("\"[path]\"") + ". A common error here " +
			"is trying to concatenate multiple values into one statement " +
			"call here. An example of a working version might be " + 
			tools.ColouriseCyan("makedirectory ") +
			tools.ColouriseGreen("\"test_dir\"") + ".",
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the path string
	path := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
		substituted
	*/
	path = VariableTemplater(path)

	// Get the last character
	last_char := path[len(path)-1:]

	// If the last character isn't a path seperator, add it
	if last_char != string(os.PathSeparator) {
		path = path + string(os.PathSeparator)
	}

	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...",
			tools.ColouriseBlue("Making"),
			tools.ColouriseGreen(path),
		)
	}

	mk_err := os.MkdirAll(path, 0750)
	if mk_err != nil {
		investigator.Report(
			"Error creating the directory " + tools.ColouriseYellow(path) +
			". Check to make sure that you have the right permissions to " +
			"the parent directory.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	if values.MODE_VERBOSE {
		fmt.Println("done!")
	}
}

