/*
This files houses the 'main' statement methods, that is, those called directly
by the engine's Call() function.
*/
package parser

import (
	"appetit/utils"
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/*
ask statement

Set a variable. Parameters include the tokens. Returns the final value of
the variable.
*/
func Ask(tokens []Token) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("ask")+" statement needs "+
				"to follow the form:\n\n\t"+utils.ColouriseCyan("ask")+" "+
				utils.ColouriseGreen("\"[question/prompt]\"")+
				utils.ColouriseMagenta(" to ")+
				utils.ColouriseYellow("\"[variable name]\"")+"\n\nAn example "+
				"of a working version check might be:\n\n\t"+
				utils.ColouriseCyan("ask")+" "+
				utils.ColouriseGreen("\"What is your name?\"")+
				utils.ColouriseMagenta(" to ")+
				utils.ColouriseGreen("\"name\"")+"\n\n"+
				"Your line of code looks like the following:\n\n\t"+
				utils.ColouriseRed(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Fix the prompt to ensure that quotation marks and escapes are handled
	properly.
	*/
	prompt := FixStringCombined(tokens[2].TokenValue)
	// Get a templated value for the prompt
	prompt = VariableTemplater(prompt)

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

	/* Fix the variable name to ensure that quotation marks and escapes are
	handled properly.
	*/
	variable_name := FixStringCombined(tokens[4].TokenValue)

	/* Get the prefix of the variable so that we can check that it isn't
	reserved
	*/
	// Hold the (possible) prefix for checking
	var variable_prefix string
	// If the length of the variable is less than the RESERVED_VARIABLE_PREFIX
	if len(variable_name) < len(SYMBOL_RESERVED_VARIABLE_PREFIX) {
		// Just set the prefix to the variable
		variable_prefix = string(variable_name)
	} else {
		// Otherwise, create a prefix to check against
		variable_prefix = string(
			variable_name[0:len(SYMBOL_RESERVED_VARIABLE_PREFIX)],
		)
	}

	// Check the variable prefix
	var_prefix_error := CheckVariablePrefix(
		loc, variable_prefix, variable_name)
	if var_prefix_error != nil {
		ReportWithFixes(
			var_prefix_error.Error(),
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}

	// Check that the variable name is not one of the statement names
	statement := CheckIsStatement(variable_name)
	// If it is a statement
	if statement {
		ReportWithFixes(
			"The variable - "+utils.ColouriseYellow(variable_name)+" - "+
				"is not a valid variable name as it conflicts with a statement "+
				"name.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s user \"%s\" and saving to variable %s...\n",
			utils.ColouriseBlue("Asking"),
			utils.ColouriseGreen(prompt),
			utils.ColouriseYellow(variable_name),
		)
	}

	/* Create a reader to get the input from user. Create a buffer size of
	65,536 bytes which doesn't seem to be acknowledged by any operating
	system. See issue #1 on GitHub for more. Leave this as-is though as
	it does allow for some extra space for input on platforms such as
	Windows. This is also potentially an issue with stdin limitations on
	any one given platform.
	*/
	input_reader := bufio.NewReaderSize(os.Stdin, 65536)
	// Prompt as per the prompt provided by the script
	fmt.Print(prompt)
	/* Read in the line while looking for the new line character as the
	delimiter
	*/
	//user_input, user_input_error := input_reader.ReadString('\n')
	user_input, user_input_error := input_reader.ReadString('\n')
	// Convert the user_input_bytes to a string
	user_input = strings.TrimSuffix(user_input, "\n")

	if user_input_error != nil {
		Report(
			"There was an error getting the user input. Please report the "+
				"following error in yellow to the project's GitHub repository "+
				"and a copy of the script:\n\n"+
				utils.ColouriseYellow(user_input_error.Error()),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Get the final variable value here by checking to see if the value is
	a math expression
	*/
	final_variable_value := CalculateValue(loc, user_input)

	// Set the variable
	VARIABLES[variable_name] = final_variable_value

	// Return the final value of the variable
	return final_variable_value
}

/*
copyfile statement

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
copydirectory statement

Copy a directory from one place to another. The parameters are the
conventional set of tokens. Returns nothing.
*/
func CopyPath(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, token_err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if token_err != nil {
		Report(
			"The "+utils.ColouriseCyan("copydirectory")+
				"statement needs to follow the form "+
				utils.ColouriseCyan("copydirectory")+
				utils.ColouriseGreen(" \"[path]\"")+" to "+
				utils.ColouriseGreen("\"[path]\"")+". A common issue "+
				"is the  use of an inappropriate action symbol ("+
				utils.ColouriseMagenta(SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				utils.ColouriseCyan("copyfcopydirectoryile")+
				utils.ColouriseGreen(" \"test_dir\"")+" to "+
				utils.ColouriseGreen(" \"new_dir\""),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Get the source folder to copy and fix the string where need be
	source_path := FixStringCombined(tokens[2].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	source_path = FixPathSeperators(source_path)
	/* Get a templated value, that is, a variable where values have
	been substituted.
	*/
	source_path = VariableTemplater(source_path)

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
	// Get the source folder to copy and fix the string where need be
	dest_path := FixStringCombined(tokens[4].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	dest_path = FixPathSeperators(dest_path)
	/* Get a templated value, that is, a variable where values have
	been substituted.
	*/
	dest_path = VariableTemplater(dest_path)

	// Set up a map of values to be passed to the file walker
	walker_values := make(map[string]string)
	walker_values["source"] = source_path
	walker_values["destination"] = dest_path
	walker_values["loc"] = loc
	walker_values["full_loc"] = full_loc
	walker_values["source_position"] = tokens[2].TokenPosition
	walker_values["dest_position"] = tokens[4].TokenPosition

	// Walk the files are start copying
	filepath.Walk(source_path, CopyPathWalker(walker_values))
}

/*
makedirectory statement

Make a directory. The tokens are passed to get the file that will be moved.
Returns nothing.
*/
func CreatePath(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("makedirectory")+" statement "+
				"needs to follow the form "+
				utils.ColouriseCyan("makedirectory")+" "+
				utils.ColouriseYellow("\"[path]\"")+". A common error here "+
				"is trying to concatenate multiple values into one statement "+
				"call here. An example of a working version might be "+
				utils.ColouriseCyan("makedirectory ")+
				utils.ColouriseGreen("\"test_dir\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the path string
	path := FixStringCombined(tokens[2].TokenValue)
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

	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...",
			utils.ColouriseBlue("Making"),
			utils.ColouriseGreen(path),
		)
	}

	mk_err := os.MkdirAll(path, 0750)
	if mk_err != nil {
		Report(
			"Error creating the directory "+utils.ColouriseYellow(path)+
				". Check to make sure that you have the right permissions to "+
				"the parent directory.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	if MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
deletefile statement

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
deletedirectory statement

Delete a directory. The parameters are the conventional set of tokens.
Returns nothing.
*/
func DeletePath(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("deletedirectory")+" statement "+
				"needs to follow the form "+
				utils.ColouriseCyan("deletedirectory")+" "+
				utils.ColouriseYellow("\"[path]\"")+". A common error here "+
				"is trying to concatenate multiple values into one statement "+
				"call here. An example of a working version might be "+
				utils.ColouriseCyan("deletedirectory ")+
				utils.ColouriseGreen("\"test_dir\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the path string
	path := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	path = VariableTemplater(path)

	// If verbose mode is set, print out what's happening
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...",
			utils.ColouriseBlue("Deleting"),
			utils.ColouriseGreen(path),
		)
	}

	remove_err := os.RemoveAll(path)
	if remove_err != nil {
		Report(
			"There was an error removing "+
				utils.ColouriseMagenta(path)+". The path does not exist.",
			loc,
			"n/a",
			full_loc,
		)
	}

	if MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
download statement

This function deals with the download itself. It takes in the conventional
set of tokens as the parameter. There is no return here as there is little
reason to have one.
*/
func Download(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)

	// Check the number of tokens and ensure that it's a proper amount
	_, num_tokens_error := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if num_tokens_error != nil {
		Report(
			"The "+utils.ColouriseCyan("download")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("download")+" "+
				utils.ColouriseGreen("\"[url]\"")+"to"+
				utils.ColouriseGreen("\"[path]\"")+". An example of a "+
				"working version might be "+utils.ColouriseCyan("download")+
				utils.ColouriseGreen(" \"http://file.com/file.txt\"")+"to"+
				utils.ColouriseGreen(" \"#b_home/file.txt\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Create a temp file to hold the download before it is moved into place
	temp_file, temp_file_err := os.CreateTemp("", "appetit_dl_temp")
	// Hold the file temporarily
	temp_loc := temp_file.Name()

	if temp_file_err != nil {
		Report(
			"Issue with creating a temporary file to store the download. "+
				"The temp file that I tried to make was "+temp_loc+
				". Check to make sure that "+utils.ColouriseCyan(os.TempDir())+
				" is writeable.",
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the remote file name
	file_to_get := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	file_to_get = VariableTemplater(file_to_get)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	/* If the action is not a valid action keyword (ie. "to"), report back the
	error
	*/
	if action_error != nil {
		ReportWithFixes(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}
	// Fix the local save file name
	save_name := FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	save_name = VariableTemplater(save_name)

	// If verbose mode is set, notify the user of what is happening
	if MODE_VERBOSE {
		fmt.Println(
			":: Creating a temp file - " + temp_loc + " - to store the " +
				"download before it's moved to its final home: " + save_name +
				".",
		)
	}

	/* If the user is not using Windows, we can defer the file close. Below,
	you'll see that we explicitly close the handler for Windows users.
	*/
	if VARIABLES["b_os"] != "windows" {
		// Defer the file close.
		defer temp_file.Close()
	}

	/* Set up a client with some default transport options. Thanks to https://
	www.zenrows.com/blog/golang-net-http-user-agent#customize-ua for this
	one.
	*/
	client := &http.Client{
		Transport: &http.Transport{},
	}

	// Set up the GET request
	request, err := http.NewRequest("GET", file_to_get, nil)
	if err != nil {
		Report(
			"There was an error initiating the request to "+
				utils.ColouriseCyan(file_to_get)+". Make sure that the URL "+
				"is valid.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* This is here to ensure that the request going through looks like it is
	coming from a web browser and not Go. Some suggestions online point to
	the default Go user agent being flagged as something that might trigger
	a 403 response.
	*/
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
			"(KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	)

	// Do the request itself
	response, err := client.Do(request)
	if err != nil {
		Report(
			"There was an error getting the file - "+
				utils.ColouriseCyan(file_to_get)+". Make sure that the URL "+
				"is valid.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	defer response.Body.Close()

	/* Get the remote file name absent much of the URL so that we can report
	back which file we are downloading.
	*/
	remote_file_name := path.Base(response.Request.URL.Path)
	// Note which file we are downloading
	fmt.Printf("Downloading %s\n", utils.ColouriseGreen(remote_file_name))
	// Set the file size in the WriteProgress struct
	size_counter := &WriteProgress{
		FileSize: float64(response.ContentLength),
	}
	/* Copy the chunk downloaded to our temp file. Here, we're copying to the
	temp_file and the source is set as a TeeReader which returns a reader
	that reads the body of the response and writes, via the WriteProgress
	type, the size of what has been downloaded.
	*/
	_, io_err := io.Copy(temp_file, io.TeeReader(response.Body, size_counter))
	// If there is an error in the saving of the chunk, report it
	if io_err != nil {
		Report(
			"There is an error saving the downloaded chunk: "+io_err.Error(),
			loc,
			tokens[1].TokenPosition,
			full_loc,
		)
	}

	/* Windows, it would seem, does not like deferred file closing so we need
	to do it manually here. This was noted above so this is handled here.
	Thanks to https://www.reddit.com/r/golang/comments/g5ftg0/
	osrename_on_windows/ for that one. As much as the implicit goal of this
	code is not to have OS specific bits, sometimes "Windows gonna Windows"
	and its idiosyncracies need to be accounted for.
	*/
	if VARIABLES["b_os"] == "windows" {
		// Close the file handler.
		temp_file.Close()
	}

	// Check if the save name is a directory
	info, info_err := os.Stat(save_name)

	if info_err == nil {
		/* If it is a directory, add the remote file name so that the downloaded
		file has the same name as the remote name.
		*/
		if info.IsDir() {
			// If it is a path, get the last character
			last_char_save_name := string(save_name[len(save_name)-1])
			// See if the last character is a path seperator
			if last_char_save_name == string(os.PathSeparator) {
				// If it is, just combine the save path with the file name
				save_name = save_name + remote_file_name
			} else {
				/* Otherwise, add a path seperator between the save path and file
				name
				*/
				save_name = save_name + string(os.PathSeparator) + remote_file_name
			}
		}
	}

	// Move the temp file to where the user wants it
	rename_err := os.Rename(temp_loc, save_name)
	// If there was an error renaming the file, report that
	if rename_err != nil {
		/* Make the second token value set to the temp location so that we
		can pass the tokens to the copyfile function. We also need to
		fix the value of the destination token to include the proper name.
		*/
		tokens[2].TokenValue = temp_loc
		tokens[4].TokenValue = save_name
		// Copy the file
		CopyFile(tokens)
		// Remove the temp file
		remove_err := os.Remove(temp_loc)
		if remove_err != nil {
			Report(
				"There was an error removing the temp file: "+
					utils.ColouriseYellow(save_name)+". It will be worth "+
					"trying to remove it manually.",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}
	}

	// Report that the file is downloaded
	fmt.Printf("\nFile downloaded to %s\n", utils.ColouriseGreen(save_name))

	/* This is a macOS specific fix to accommodate the fact that macOS seems to
	make the file hidden when the file is moved. Here, "macOS gonna macOS"
	and, again, as much as OS specific code is implicitly avoided, Finder's
	desire to hide things by default makes the download statement largely
	useless for those preferring the graphical user interface for file
	management.
	*/
	if VARIABLES["b_os"] == "darwin" {
		// Unhide the file
		macos_unhide := exec.Command("chflags", "nohidden", save_name)
		/* Capture the output and suppress it as there isn't any but we may
		need to capture any error that is returned.
		*/
		_, unhide_err := macos_unhide.Output()
		/* If there was an error, report back that the issue here means that
		the file is hidden but is still there.
		*/
		if unhide_err != nil {
			unhide_keycombo := fmt.Sprintf(
				"%s-%s-%s",
				utils.ColouriseGreen("Command"),
				utils.ColouriseGreen("Shift"),
				utils.ColouriseGreen("."),
			)
			Report(
				"On macOS, the file is hidden by the operating system by "+
					"default. There was an attempt to unhide it but it failed. "+
					"You will want to enable the showing of hidden files in "+
					"Finder. Pressing "+unhide_keycombo+"will show you the "+
					"file, after which you can restore Finder to \"normal\" by"+
					"pressing "+unhide_keycombo+" again.",
				loc,
				tokens[4].TokenPosition,
				full_loc,
			)
		}
	}

}

/*
execute statement

Execute a system command. Parameters include the tokens. Returns nothing.
*/
func ExecuteCommand(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("execute")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("execute")+" "+
				utils.ColouriseGreen("\"[command]\"")+". A common "+
				"issue here is excluding a command. An example of a working "+
				"statement might be "+utils.ColouriseCyan("execute")+
				utils.ColouriseGreen(" \"ls\"")+"."+"\n\nLine of Code: "+
				utils.ColouriseMagenta(full_loc),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Get the command and fix the string
	command := FixStringCombined(tokens[2].TokenValue)

	/* Check if the -allowexec flag was passed to the app and if not, throw
	an error
	*/
	if !MODE_ALLOW_EXEC {
		Report(
			"You are unable to execute system commands. If you would like "+
				"to do so, you need to run with the "+
				utils.ColouriseYellow("-allowexec")+" flag.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s...\n",
			utils.ColouriseBlue("Executing"),
			utils.ColouriseYellow(command),
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
		Report(
			"The application "+utils.ColouriseYellow(command)+
				" was not found. Perhaps it was a typo?",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
	// Output the results of the command
	fmt.Println(string(output))
}

/*
exit statement

Handle an exit statement call. This one is very basic and doesn't require
much of the end user other than the statement call itself.
*/
func Exit(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 1)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("exit")+" statement needs "+
				"to follow the form:\n\n\t"+utils.ColouriseCyan("exit")+
				"\n\nThere are no values that you can or need to pass which "+
				"is most likely the cause here.\n\n"+
				"Your line of code looks like the following:\n\n\t"+
				utils.ColouriseRed(full_loc)+"\n\n",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set
	if MODE_VERBOSE {
		fmt.Println(":: Exiting...")
	}
	// Finally, exit
	os.Exit(0)
}

/*
log statement

This will log a string to a file of the user's choosing as a helpful shorthand
for tracking executions of a script. Returns nothing.
*/
func Log(tokens []Token) {
	full_loc := tokens[0].FullLineOfCode

	loc := strconv.Itoa(tokens[0].LineNumber)

	_, num_tokens_error := CheckValidNumberOfTokens(tokens, 4)

	if num_tokens_error != nil {
		Report(
			"The "+utils.ColouriseCyan("log")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("log")+" "+
				utils.ColouriseGreen("\"[path]\"")+"to"+
				utils.ColouriseGreen("\"[path]\"")+". An example of a "+
				"working version might be "+utils.ColouriseCyan("log")+
				utils.ColouriseGreen(" \"The script is done\"")+"to"+
				utils.ColouriseGreen(" \"script_log\"")+".",
			loc,
			"n/a",
			full_loc,
		)
	}

	action := tokens[3].TokenValue
	if action != SYMBOL_ACTION {
		Report(
			"An inapportiate action symbol is used. You used "+
				utils.ColouriseMagenta(action)+" when you need to use "+
				utils.ColouriseMagenta(SYMBOL_ACTION)+".",
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	/* Get the output string, that is, the text that will be output to the log.
	Further, fix the string.
	*/
	output_string := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	output_string = VariableTemplater(output_string)

	// Get the file name
	file_name := FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	file_name = VariableTemplater(file_name)
	// Open the log file and create it if it doesn't exist
	file_handler, file_handler_error := os.OpenFile(
		file_name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// If there's an error opening or creating the file...
	if file_handler_error != nil {
		Report(
			"There was an error opening the file "+
				utils.ColouriseCyan(file_name)+". Make sure that you can "+
				"create a log file in this directory.",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	// Defer the file handler close
	defer file_handler.Close()

	_, write_handler_error := file_handler.WriteString(output_string)
	if write_handler_error != nil {
		Report(
			"There was an error writing to the file "+
				utils.ColouriseCyan(file_name)+". Make sure that you can "+
				"write to files in this directory.",
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}

	if MODE_VERBOSE {
		fmt.Println("Wrote log file to " + file_name + ".log.")
	}
}

/*
makefile statement

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

/*
minver statement

Check the minimum version required to run the script. Parameters include
the tokens. Returns nothing.
*/
func MinVer(tokens []Token) int {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, token_number_err := CheckValidNumberOfTokens(tokens, 2)
	/* Get the minimum version as a string. As we check whether this is an
	integer above, we should be good to go here.
	*/
	min_ver_string := tokens[2].TokenValue
	// Get the minver set by the user as an integer for comparison
	min_ver, int_conversion_err := strconv.Atoi(min_ver_string)
	/* If there is an error trying to do the conversion or if the min_ver is
	less than zero, report an error. This also captures negative integers
	in particular as the negative sign and the integer are tokenised as
	seperate tokens.
	*/
	if int_conversion_err != nil || min_ver <= 0 || token_number_err != nil {
		Report(
			"The "+utils.ColouriseCyan("minver")+" statement needs to "+
				"include a valid non-zero positive integer. A valid "+
				utils.ColouriseCyan("minver")+" statement needs to follow the "+
				"form:\n\t"+utils.ColouriseCyan("minver")+
				utils.ColouriseYellow(" [version number]")+"\nAn example of a "+
				"working version check might be:\n\t"+
				utils.ColouriseCyan("minver")+utils.ColouriseYellow(" 3")+
				"\nMake sure that you have none of the following for the "+
				utils.ColouriseCyan("minver")+" statement value:\n\t"+
				"- Negative number\n\t- Float (ie. decimal number)\n\t"+
				"- String\n\t- No value\n",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Check if the minimum version is greater than or equal to the language
	version.
	*/
	if min_ver > LANG_VERSION {
		Report(
			"The script you're running here requires a newer version of "+
				"the interpreter. You are running version "+
				utils.ColouriseYellow(strconv.Itoa(LANG_VERSION))+
				" but the script requires at least version "+
				utils.ColouriseYellow(min_ver_string)+". Check to see if "+
				"a newer version is available. ",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	if MODE_VERBOSE {
		fmt.Printf(
			":: Setting the minimum version of Appetit (%s) required to "+
				"run this script to %s\n",
			utils.ColouriseBlue("minver"),
			utils.ColouriseGreen(strconv.Itoa(min_ver)),
		)
	}

	// Return the minimum version required
	return min_ver
}

/*
movefile statement

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
movedirectory statement
Move a directory. The parameters are the conventional set of tokens.
Returns nothing.
*/
func MovePath(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("movedirectory")+" statement "+
				"needs to follow the form "+
				utils.ColouriseCyan("movedirectory")+" "+
				utils.ColouriseYellow("\"[path]\"")+". A common error here "+
				"is trying to concatenate multiple values into one statement "+
				"call here. An example of a working version might be "+
				utils.ColouriseCyan("movedirectory ")+
				utils.ColouriseGreen("\"test_dir\"")+" to "+
				utils.ColouriseGreen("\"actual_dir\""),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Get the source folder to copy and fix the strings
	old_path := FixStringCombined(tokens[2].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	old_path = FixPathSeperators(old_path)
	/* Get a templated value, that is, a variable where values have
	been substituted.
	*/
	old_path = VariableTemplater(old_path)

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
	// Get the destination folder to copy and fix the strings
	new_path := FixStringCombined(tokens[4].TokenValue)
	// Fix the path seperators to ensure that the last character is a seperator
	new_path = FixPathSeperators(new_path)
	/* Get a templated value, that is, a variable where values have
	been substituted.
	*/
	new_path = VariableTemplater(new_path)

	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Moving"),
			utils.ColouriseGreen(old_path),
			utils.ColouriseGreen(new_path),
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
		/*Report(
			"There was an error moving the directory. Check to ensure that " +
			"the source - " + utils.ColouriseYellow(old_path) + " - and the " +
			"destination - " + utils.ColouriseYellow(new_path) + " - " +
			"are valid locations and that " + utils.ColouriseYellow(new_path) +
			" doesn't already exist.",
			loc,
			"n/a",
			full_loc,
		)*/
	}

	if MODE_VERBOSE {
		fmt.Println("done!")
	}
}

/*
pause statement

Pause the execution of the script. Parameters include the tokens. Returns
nothing.
*/
func Pause(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("pause")+" statement needs to "+
				"follow the form "+utils.ColouriseCyan("pause")+" "+
				utils.ColouriseYellow("[version number]")+". A common "+
				"issue here is excluding a version number or passing one as a "+
				"string (eg. "+utils.ColouriseGreen("\"3\"")+"). An "+
				"example of a working version check might be "+
				utils.ColouriseCyan("pause")+utils.ColouriseYellow(" 3"),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	/* Convert the loc to a simple number to keep things simple. We can discard
	the line number as this will always be an integer.
	*/
	pause_as_string := tokens[2].TokenValue

	// Create an integer version of the pause length
	pause_int, err := strconv.Atoi(pause_as_string)
	/* If there is an error trying to do the conversion or if the pause_int is
	less than zero, report an error.
	*/
	if err != nil || pause_int < 0 {
		Report(
			"The version number "+utils.ColouriseYellow(pause_as_string)+
				" is not a valid version. You need to use a positive non-zero "+
				"integer.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// If verbose mode is set...
	if MODE_VERBOSE {
		fmt.Printf(":: Pausing for %s seconds...", pause_as_string)
	}
	// Pause execution by sleeping for the required number of seconds
	time.Sleep(time.Duration(pause_int) * time.Second)
}

/*
run statement

Run a script from elsewhere. Parameters include the tokens. Returns
nothing.
*/
func Run(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("run")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("run")+
				utils.ColouriseGreen(" \"[script]\"")+". An example of a "+
				"working version check might be "+utils.ColouriseCyan("run")+
				utils.ColouriseGreen("\"other_script.apt\"")+"\n\n"+
				"Line of Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	/* Set the path name for the script to be run, fixing any issues with the
	string.
	*/
	script_name := FixStringCombined(tokens[2].TokenValue)
	// Replace any variables in the output string
	script_name = VariableTemplater(script_name)

	file_exists := CheckFileExists(script_name)

	if !file_exists {
		Report(
			"The script - "+utils.ColouriseYellow(script_name)+" - does "+
				"not exist and/or can't be accessed. Double check to verify "+
				"that the script exists.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	contents := PrepScript(script_name)

	if MODE_DEV {
		// Start printing out the tokens
		fmt.Println(utils.ColouriseYellow("\nTokens"))
		Start(contents, true)
	} else {
		Start(contents, false)
	}
}

/*
set statement

Set a variable. Parameters include the tokens. Returns the variable value.
*/
func Set(tokens []Token) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("set")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("set")+" "+
				utils.ColouriseYellow("[variable name]")+" = "+
				utils.ColouriseGreen("\"[value]\"")+". An example of a "+
				"working version check might be "+utils.ColouriseCyan("set")+
				" name = "+
				utils.ColouriseGreen("\""+LANG_NAME+"\"")+"\n\n"+
				"Line of Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Set the variable name
	variable_name := tokens[2].TokenValue
	// The assignment operator
	assignment_operator := tokens[3].TokenValue
	// The variable value with fixes that need to be made
	variable_value := FixStringCombined(tokens[4].TokenValue)

	/* Get the prefix of the variable so that we can check that it isn't
	reserved
	*/
	// Hold the (possible) prefix for checking
	var variable_prefix string
	// If the length of the variable is less than the RESERVED_VARIABLE_PREFIX
	if len(variable_name) < len(SYMBOL_RESERVED_VARIABLE_PREFIX) {
		// Just set the prefix to the variable
		variable_prefix = string(variable_name)
	} else {
		// Otherwise, create a prefix to check against
		variable_prefix = string(variable_name[0:2])
	}

	// Check the variable prefix
	var_prefix_error := CheckVariablePrefix(
		loc, variable_prefix, variable_name)
	if var_prefix_error != nil {
		ReportWithFixes(
			var_prefix_error.Error(),
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Check that the variable name is not one of the statement names
	statement := CheckIsStatement(variable_name)
	// If it is a statement
	if statement {
		ReportWithFixes(
			"The variable - "+utils.ColouriseYellow(variable_name)+" - "+
				"is not a valid variable name as it conflicts with a statement "+
				"name.",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}

	// Check for a valid assignment operator
	assignment_error := CheckValidAssignment(loc, assignment_operator)
	if assignment_error != nil {
		ReportWithFixes(
			assignment_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	templated_variable := VariableTemplater(variable_value)

	/* Get the final variable value here by checking to see if the value is
	a math expression
	*/
	final_variable_value := CalculateValue(loc, templated_variable)

	// If verbose mode is set...
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Setting"),
			utils.ColouriseYellow(variable_name),
			utils.ColouriseGreen(final_variable_value),
		)
	}
	// Set the variable
	VARIABLES[variable_name] = final_variable_value

	// If verbose mode is set, report that things are done.
	if MODE_VERBOSE {
		fmt.Println("done!")
	}

	return final_variable_value

}

/*
write and writeln statement

Write output to its own line or on one line. This handles both the write
and writeln statement. Parameters include the tokens, and newline as a bool
for whether output needs to add a new line (writeln) or leave the line
without a newline character at the end. Returns nothing.
*/
func Writeln(tokens []Token, newline bool) string {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 2)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("write/writeln")+" statement "+
				"needs to follow the form "+
				utils.ColouriseCyan("write/writeln")+" "+
				utils.ColouriseYellow("[content to be written]")+". A "+
				"common error here is trying to concatenate multiple values "+
				"into one statement call here. An example of a working version "+
				"might be "+utils.ColouriseCyan("write/writeln ")+
				utils.ColouriseGreen("\"Hello World\"")+"\n\nLine of Code: "+
				utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}

	// Fix the string to be printed
	trimmed_output := FixStringCombined(tokens[2].TokenValue)
	// Replace any variables in the output string
	trimmed_output = VariableTemplater(trimmed_output)

	/* If newline is true, we are parsing a writeln, otherwise, we are parsing
	a write
	*/
	if newline {
		// Print out the output with a newline as we are parsing a writeln
		fmt.Printf("%s\n", trimmed_output)
		return fmt.Sprintf("%s\n", trimmed_output)
	} else {
		// Print out the output with a newline as we are parsing a write
		fmt.Print(trimmed_output)
		return trimmed_output
	}
}

/*
zipfile statement

Make a zip archive of a file. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate.
Returns nothing. Thanks to https://earthly.dev/blog/golang-zip-files/
*/
func ZipFromFile(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("zipfile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("zipfile")+" "+
				utils.ColouriseGreen("\"[path]\"")+" to "+
				utils.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				utils.ColouriseMagenta(SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				utils.ColouriseCyan("zipfile")+
				utils.ColouriseGreen(" \"/Users/user/test_dir.txt\"")+" to "+
				utils.ColouriseGreen(" \"test_dir.zip\"")+"\n\nLine of "+
				"Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the source string
	source := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	// If there's an error in the action keyword, report it
	if action_error != nil {
		Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix up the destination string
	destination := FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	// If verbose mode is set, note that we're zipping a file
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Zipping"),
			utils.ColouriseGreen(source),
			utils.ColouriseGreen(destination),
		)
	}

	// Create the destination, ie. the zip file that will house our file
	archive_name, archive_name_error := os.Create(destination)
	// If there's an error with creating the archive, report it
	if archive_name_error != nil {
		Report(
			"The archive name you provided - "+
				utils.ColouriseYellow(destination)+" - could "+
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
		Report(
			"Couldn't open "+utils.ColouriseYellow(source)+"! Is it "+
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
		Report(
			"Couldn't add "+utils.ColouriseYellow(source)+" to "+
				utils.ColouriseYellow(destination)+". Something went wrong "+
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
		Report(
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
	if MODE_VERBOSE {
		fmt.Printf(
			"done! "+
				utils.ColouriseMagenta(
					"[%s bytes written]\n",
				),
			strconv.FormatInt(copy_file_size, 10),
		)
	}

	// Close the zip writer
	zip_writer.Close()
}

/*
zipdirectory statement

Make a zip archive of a file. The tokens are passed to get
the origin, destination, and to ensure that the 'action' is appropriate. This
is a modified version of this function (consider it version 2) as it migrates
from a file path walker to the os.DirFS and zip writer AddFS() functions.
Returns nothing.
*/
func ZipFromPath(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
	// Check the number of tokens and ensure that it's a proper amount
	_, err := CheckValidNumberOfTokens(tokens, 4)
	// If not a valid number of tokens, report an error
	if err != nil {
		Report(
			"The "+utils.ColouriseCyan("zipfile")+" statement needs "+
				"to follow the form "+utils.ColouriseCyan("zipfile")+" "+
				utils.ColouriseGreen("\"[path]\"")+" to "+
				utils.ColouriseGreen("\"[path]\"")+". A common issue is the "+
				"use of an inappropriate action symbol ("+
				utils.ColouriseMagenta(SYMBOL_ACTION)+"). An "+
				"example of a working version might be "+
				utils.ColouriseCyan("zipfile")+
				utils.ColouriseGreen(" \"/Users/user/test_dir.txt\"")+" to "+
				utils.ColouriseGreen(" \"test_dir.zip\"")+"\n\nLine of "+
				"Code: "+utils.ColouriseMagenta(full_loc),
			loc,
			"n/a",
			full_loc,
		)
	}
	// Fix up the source string
	source := FixStringCombined(tokens[2].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	source = VariableTemplater(source)

	// Get the action to ensure that it can be checked
	action := tokens[3].TokenValue
	// Check the action keyword to ensure that it's valid
	action_error := CheckAction(loc, action)
	// If there's an error in the action keyword, report it
	if action_error != nil {
		Report(
			action_error.Error(),
			loc,
			tokens[3].TokenPosition,
			full_loc,
		)
	}

	// Fix up the destination string
	destination := FixStringCombined(tokens[4].TokenValue)
	/* Get a templated value, that is, a variable where values have been
	substituted
	*/
	destination = VariableTemplater(destination)

	// If verbose mode is set, note that we're zipping a file
	if MODE_VERBOSE {
		fmt.Printf(
			":: %s %s to %s...",
			utils.ColouriseBlue("Zipping"),
			utils.ColouriseGreen(source),
			utils.ColouriseGreen(destination),
		)
	}

	// Create the archive file
	archive_file, archive_file_error := os.Create(destination)
	if archive_file_error != nil {
		Report(
			"Error creating the archive at "+destination,
			loc,
			tokens[4].TokenPosition,
			full_loc,
		)
	}
	defer archive_file.Close()

	// Create an archive writer
	archive_writer := zip.NewWriter(archive_file)
	defer archive_writer.Close()

	// Set a filesystem path to the source
	archive_path := os.DirFS(source)
	// Ad the filesystem path to the archive
	archive_fs_error := archive_writer.AddFS(archive_path)
	if archive_fs_error != nil {
		Report(
			"Error adding path to the archive",
			loc,
			tokens[2].TokenPosition,
			full_loc,
		)
	}
}
