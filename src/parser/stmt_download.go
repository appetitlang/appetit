/*
Much of this module is built on the following gist from GitHub:
https://gist.github.com/cnu/026744b1e86c6d9e22313d06cba4c2e9

This module deals with the download statement by sending a GET request and
handling progress tracking and reporting back the progress.
*/
package parser

import (
	"appetit/utils"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
)

/*
Hold values of the progress of the writing of downloaded data. This has two
values: TotalBytes (which holds how many bytes have been downloaded) and
FileSize (which holds the total number of bytes of the file being
downloaded). Somewhere down the line, a 64-bit integer is returned as the
response's content length and so, we are sticking with 64-bit numbers
throughout.
*/
type WriteProgress struct {
	TotalBytes float64
	FileSize   float64
}

/*
Handle writing, to the WriteProgress struct, the progress (TotalBytes).
This takes in the progress and adds that to the total. Additionally, it
prints out the progress to the screen for the user. This returns the length
and returns nil as an error.
*/
func (wp *WriteProgress) Write(progress []byte) (int, error) {
	// Get the length of the progress
	length := len(progress)
	// Add the length of the progress byte slice to the total bytes
	wp.TotalBytes += float64(length)
	/* Create an output writer that uses stdout as the output. The reason we
	aren't using the fmt module here is to ensure that text can be flushed
	from the buffer properly which allows writing over the lines cleanly.
	*/
	writer := bufio.NewWriter(os.Stdout)
	// Calculate the percentage
	percentage := wp.TotalBytes / wp.FileSize
	// Format the progress as a percentage for printing.
	// TODO: Add comma seperators for download sizes
	progress_output := fmt.Sprintf(
		"\rDownloaded %s (%s KB of %s KB)",
		utils.ColouriseMagenta(
			strconv.FormatFloat(percentage*100, 'f', 2, 32)+"%",
		),
		CommaSeperator(wp.TotalBytes/1024),
		CommaSeperator(wp.FileSize/1024),
	)
	// Write out the progress
	writer.WriteString(progress_output)
	// Flush out the standard output
	writer.Flush()
	/* Return the length of the error and an error value of nil here. While
	it may be poor practice to return nil here without any other type of
	value, there needs to be space for an error in case this Write()
	function get's more elaborate and/or something, in the future, reveals
	a real possibility that the tracking might cause an error.
	*/
	return length, nil

}

/*
This function deals with the download itself. It takes in the conventional
set of tokens as the parameter. There is no return here as there is little
reason to have one.
*/
func Download(tokens []Token) {
	// Get the full line of code
	full_loc := tokens[0].FullLineOfCode
	// Get the line of code
	loc := strconv.Itoa(tokens[0].LineNumber)
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
