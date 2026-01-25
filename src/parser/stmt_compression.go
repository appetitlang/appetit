/*
This module deals with the zipdirectory and zipfile statements.
*/
package parser

import (
	"appetit/investigator"
	"appetit/tools"
	"appetit/values"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
Make a zip archive of a file. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate.
Returns nothing. Thanks to https://earthly.dev/blog/golang-zip-files/
*/
func ZipFromFile(tokens []values.Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("zipfile")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("zipfile")+" "+
				tools.ColouriseGreen("\"[path]\"")+" to "+
				tools.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				tools.ColouriseMagenta(values.SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				tools.ColouriseCyan("zipfile")+
				tools.ColouriseGreen(" \"/Users/user/test_dir.txt\"")+" to "+
				tools.ColouriseGreen(" \"test_dir.zip\"")+"\n\nLine of "+
				"Code: "+tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the source string
	source := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := investigator.CheckAction(loc, action)
	// If there's an error in the action keyword, report it
	if action_error != nil {
		investigator.Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix up the destination string
	destination := tools.FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	// If verbose mode is set, note that we're zipping a file
	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			tools.ColouriseBlue("Zipping"),
			tools.ColouriseGreen(source),
			tools.ColouriseGreen(destination),
		)
	}

	// Create the destination, ie. the zip file that will house our file
	archive_name, archive_name_error := os.Create(destination)
	// If there's an error with creating the archive, report it
	if archive_name_error != nil {
		investigator.Report(
			"The archive name you provided - "+
				tools.ColouriseYellow(destination)+" - could "+
				" not be created. Is it possible that you can't write to that "+
				"path?",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	// Defer the closing of the archive
	defer archive_name.Close()

	// Create a zip writer
	zip_writer := zip.NewWriter(archive_name)

	// Open up the source file so that we can add it to our zip file
	file_to_zip, file_to_zip_error := os.Open(source)
	// If there's an error opening up the file that will be zipped, report it
	if file_to_zip_error != nil {
		investigator.Report(
			"Couldn't open "+tools.ColouriseYellow(source)+"! Is it "+
				"possible that this file doesn't exist?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// Defer closing the file that will be zipped
	defer file_to_zip.Close()

	/* Split the origin by the os path separator so that we can get the
	file name in case we need to append it
	*/
	file_path_parts := strings.Split(source, string(os.PathSeparator))
	// Get the filename from the parts of the file path
	filename := file_path_parts[len(file_path_parts)-1]

	// Create the file that will house the contents in the zip file
	add_file, add_file_error := zip_writer.Create(filename)
	// If there was an error creating the file in the zip archive, report it
	if add_file_error != nil {
		investigator.Report(
			"Couldn't add "+tools.ColouriseYellow(source)+" to "+
				tools.ColouriseYellow(destination)+". Something went wrong "+
				"with adding the file to the archive.",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}

	/* Start the copy of the original file into the file that lives in the zip
	archive
	*/
	copy_file_size, copy_file_err := io.Copy(add_file, file_to_zip)
	/* If there was an error finalising the copy of the original file into the
	archive, report it
	*/
	if copy_file_err != nil {
		investigator.Report(
			"Couldn't copy data from "+source+" to "+destination+
				"Check to make sure that the original file can be read.",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}

	/* If verbose mode is set, report back that we're done along with how many
	bytes were written
	*/
	if values.MODE_VERBOSE {
		fmt.Printf(
			"done! "+
				tools.ColouriseMagenta(
					"[%s bytes written]\n",
				),
			strconv.FormatInt(copy_file_size, 10),
		)
	}

	// Close the zip writer
	zip_writer.Close()
}

/*
Make a zip archive of a folder. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate.
Returns nothing. Thanks to https://stackoverflow.com/a/63233911
*/
func ZipFromPath(tokens []values.Token) {

	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := investigator.ValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		investigator.Report(
			"The "+tools.ColouriseCyan("zipdirectory")+" statement needs "+
				"to follow the form "+tools.ColouriseCyan("zipdirectory")+" "+
				tools.ColouriseGreen("\"[path]\"")+" to "+
				tools.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				tools.ColouriseMagenta(values.SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				tools.ColouriseCyan("zipdirectory")+
				tools.ColouriseGreen(" \"/Users/user/test_dir/\"")+" to "+
				tools.ColouriseGreen(" \"test_dir.zip\"")+"\n\nLine of "+
				"Code: "+tools.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the source string
	source := tools.FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := investigator.CheckAction(loc, action)
	// If there's an error in the action keyword, report it
	if action_error != nil {
		investigator.Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix up the destination string
	destination := tools.FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	// If verbose mode is set, note that we're zipping a file
	if values.MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...\n",
			tools.ColouriseBlue("Zipping"),
			tools.ColouriseGreen(source),
			tools.ColouriseGreen(destination),
		)
	}

	// Create the destination, ie. the zip file that will house our file
	archive_name, archive_name_error := os.Create(destination)
	// If there's an error with creating the archive, report it
	if archive_name_error != nil {
		investigator.Report(
			"The archive name you provided - "+
				tools.ColouriseYellow(destination)+" - could "+
				" not be created. Is it possible that you can't write to that "+
				"path?",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	// Defer the closing of the archive
	defer archive_name.Close()

	// Create a zip writer
	zip_writer := zip.NewWriter(archive_name)
	defer zip_writer.Close()

	path_walker := func(path string, info os.FileInfo, err error) error {

		// If verbose mode is set, note that we're zipping a file
		if values.MODE_VERBOSE {
			fmt.Printf(
				":: %s %s...",
				tools.ColouriseBlue("Adding"),
				tools.ColouriseGreen(path),
			)
		}

		// If there's an error walking the path, report it
		if err != nil {
			investigator.Report(
				"There was an error traversing the  "+
					tools.ColouriseYellow(path)+". Is it possible that "+
					"you can't read from that path?",
				loc,
				tokens[2].TokenPosition,
				full_loc,
			)
		}

		/* If the object that we've encountered is a path, return nil to note
		that we don't have an error but we don't want to do anything
		*/
		if info.IsDir() {
			return nil
		}

		// Open up the file to prep it for addition to our zip file
		file_path, file_path_error := os.Open(path)
		// If there's an error opening our file, report it
		if file_path_error != nil {
			investigator.Report(
				"There was an error opening  "+tools.ColouriseYellow(path)+
					". Is it possible that you can't read from that path?",
				loc,
				tokens[2].TokenPosition,
				full_loc,
			)
		}
		// Defer the close of the file
		defer file_path.Close()

		/* Replace the home directory, where need be. By default, the whole
		path is included, sometimes creating a needlessly complex set of
		nested folders. This trims off what is considered a reasonable
		number of folders (ie. the user home directory).
		*/
		path = strings.TrimPrefix(path, values.VARIABLES["b_home"])

		/* To give the unzipped folder some nice naming, we want to get the
		name of the destination zip file. So, we're going to split the name
		and then get just the file name, independent of the zip extension.
		*/
		// Get the path to the destination and split it by the path separator
		dest_split := strings.Split(destination, string(os.PathSeparator))
		// Get just the file name
		dest_name := dest_split[len(dest_split)-1]
		// Split by the period
		dest_name_split := strings.Split(dest_name, ".")
		// Get the first part of the split (ie. the file name)
		dest_name_root := dest_name_split[0]

		// Create the file that will be compressed
		zip_path := dest_name_root + string(os.PathSeparator) + path
		file, file_error := zip_writer.Create(zip_path)
		// If there's an error doing this, report it
		if file_error != nil {
			investigator.Report(
				"There was an error creating the file:  "+
					tools.ColouriseYellow(path)+". Is it possible that you "+
					"can't write to that path?",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}

		//file_path_parts := strings.Split(path, string(os.PathSeparator))
		// Get the filename from the parts of the file path
		//filename := file_path_parts[len(file_path_parts)-1]

		// Copy over the file to the archive
		bytes_written, bytes_error := io.Copy(file, file_path)
		// If there was an error, report it
		if bytes_error != nil {
			investigator.Report(
				"There was an error creating the file:  "+
					tools.ColouriseYellow(path)+". Is it possible that you "+
					"can't write to that path?",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}

		// If verbose mode is set, note that we're zipping a file
		if values.MODE_VERBOSE {
			fmt.Printf(
				":: done [%s bytes written]\n",
				tools.ColouriseMagenta(strconv.Itoa(int(bytes_written))),
			)
		}

		// Return nil as we can assume that we haven't hit any errors
		return nil
	}

	// Walk the path and call the walker function above to do all the work
	walker_error := filepath.Walk(source, path_walker)
	/* If there was an issue reported above, report it but note that this
	won't be called since it any errors will have been reported earlier.
	This is here in case there's a change above in the walker function to
	reporting errors instead.
	*/
	if walker_error != nil {
		investigator.Report(
			"There was an error traversing the  "+
				tools.ColouriseYellow(source)+". Is it possible that "+
				"you can't read from that path?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
}
