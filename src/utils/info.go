package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

func About(lang_name string, lang_ver int, lang_codename string, build_date string) {
	/* Set up the struct that will hold the JSON data for remotely checking
	the current version.
	*/
	type RemoteDetails struct {
		Version     int    `json:"version"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	// Fetch the info about the current release
	remote_response, remote_error := http.Get(
		"https://bryanabsmith.com/appetit/version_info.json",
	)
	// If there's an error getting the current version info, abandon ship
	if remote_error != nil {
		// Exit the app
		os.Exit(0)
	}
	// Defer the remote handler close
	defer remote_response.Body.Close()

	// Set up a home for the remotely pulled information.
	var remote_data RemoteDetails
	// Decode the data and place it in the remote_data RemoteDetails
	decode_error := json.NewDecoder(remote_response.Body).Decode(
		&remote_data,
	)
	// If there was an error decoding the data, abandon ship
	if decode_error != nil {
		// Exit the app
		os.Exit(0)
	}

	new_or_current_version := ""
	/* If the remote file has a version number greater than the current
	version, inform the user and give them some details.
	*/
	if remote_data.Version > lang_ver {
		new_or_current_version = ColouriseRed(
			fmt.Sprintf(
				"\n%s\nThere's a new version of Appetit "+
					"available! Version %s is available, released %s. "+
					"%s\n",
				ColouriseYellow("[New Version]"),
				strconv.Itoa(remote_data.Version),
				remote_data.Date,
				remote_data.Description,
			),
		)
		/* If the version information is the same, let them know that they are
		running the most current version.
		*/
	} else if remote_data.Version == lang_ver {
		new_or_current_version = ColouriseGreen("You're up to date!")
	}

	// Get build info
	build_info, _ := debug.ReadBuildInfo()
	// Get the go information absent the first to characters which is go
	go_version := build_info.GoVersion[2:]

	/* Print out a pretty version of the version info. During testing, the
	BuildDate variable will be "testing" and this is replaced with the
	actual build date when the Makefile is run.
	*/

	bin_dir, _ := os.Executable()

	fmt.Printf(
		"%s %s\n%s\n\n%s\n%s%s\n\n%s\n%s%s\n%s%s\n%s%d\n\n%s\n%s%s\n%s%s\n",
		ColouriseMagenta(
			lang_name+" "+strconv.Itoa(int(lang_ver)),
		),
		ColouriseMagenta(lang_codename),
		new_or_current_version,
		ColouriseYellow("[Files]"),
		ColouriseCyan("Interpreter Path: "),
		ColouriseBlue(bin_dir),
		ColouriseYellow("[Platform]"),
		ColouriseCyan("Operating System: "),
		GetHumanVersionOS(),
		ColouriseCyan("Architecture: "),
		runtime.GOARCH,
		ColouriseCyan("CPUs: "),
		runtime.NumCPU(),
		ColouriseYellow("[Build]"),
		ColouriseCyan("Go Version: "),
		go_version,
		ColouriseCyan("Build Date: "),
		build_date,
	)
	// Exit the app
	os.Exit(0)
}

/*
Get a human readable version of the underlying operating system. Returns the
human readable version as a string.
*/
func GetHumanVersionOS() string {
	// Get the GOOS platform name
	goos := runtime.GOOS
	/*
		Get the more human version of the os name reported back by the runtime
		package.
	*/
	// Set up the os variable
	var os string
	/*
		Set up the version_command variable that will hold the command executed
		to get the version number of the OS.
	*/
	var version_command string
	// Switch between them
	switch goos {
	case "darwin":
		os = "macOS"
		version_command = "sw_vers -productVersion"
	case "windows":
		/*
			This is blank as the version_command returns a full name including
			"Windows".
		*/
		os = ""
		version_command = "(Get-CimInstance Win32_OperatingSystem).Caption"
	case "linux":
		os = "Linux"
		version_command = "uname -r"
	case "freebsd":
		os = "FreeBSD"
		version_command = "uname -r"
	case "netbsd":
		os = "NetBSD"
		version_command = "uname -r"
	case "openbsd":
		os = "OpenBSD"
		version_command = "uname -r"
	}

	// Split the version_command to prep it for the execution of the command
	cmd_split := strings.Split(version_command, " ")
	// Execute the version_command
	ver, _ := exec.Command(cmd_split[0], cmd_split[1:]...).Output()
	// Format the output
	os = fmt.Sprintf("%s %s", os, string(ver))
	// Remove the newline character
	os = strings.TrimRight(os, "\n")
	// Return the os
	return os
}
