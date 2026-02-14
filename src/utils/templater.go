package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateTemplate(file_name string, lang_ver int) {
	// Create the template with a minver and a shebang line
	template_script := []byte("#!/usr/bin/appetit\nminver " +
		strconv.Itoa(lang_ver) + "\n\n- Say hello to the " +
		"world\nwriteln \"Hello World!\"\n",
	)

	// Get the user's home directory in case they pass a tilde
	user_home, _ := os.UserHomeDir()
	/* If they provide a tilde, convert the tilde to the user home
	directory
	*/
	if strings.HasPrefix(file_name, "~") {
		// Hold the corrected file name
		file_name = user_home + file_name[1:]
	}
	// Create a write handler
	write_handler, write_err := os.Create(file_name)
	// If there was an error creating the file
	if write_err != nil {
		// Respond with an error
		fmt.Println(
			"There was an error creating the script at " +
				file_name + ". Make sure that you can save a " +
				"file in that location.",
		)
	}
	// Defer the file closing
	defer write_handler.Close()

	// Write out the template
	_, output_err := write_handler.Write(template_script)

	// If there was an error writing the template to disk
	if output_err != nil {
		// Respond with an error
		fmt.Println(
			"There was an error creating the script at " +
				file_name + ". Make sure that you can save a " +
				"file in that location.",
		)
	}

	// Report back that we're done
	fmt.Println(
		":: Created a script at " +
			ColouriseCyan(file_name),
	)
	// Abandon ship
	os.Exit(0)
}
