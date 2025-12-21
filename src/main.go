/*
This is the main package that serves as the (conventional) entry point.
Nothing of note here needs to be mentioned that is unique to this package.
*/
package main

import (
	"appetit/investigator"
	"appetit/parser"
	"appetit/tools"
	"appetit/values"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unsafe"
)

// This is set as the build date but this is changed with the Makefile
var BuildDate string = "-development"

/* Thanks to https://mblessed.hashnode.dev/go-embed-embed-your-html-frontend-
	in-golang for the embed info.
*/
//go:embed docs
var template_path embed.FS

/*
	Handle serving the documentation. The parameters are the conventional
	response writer and the request. No returns.
*/
func DocsHandler(writer http.ResponseWriter, request *http.Request) {
	template, _ := template.ParseFS(template_path, "docs/index.html")
	template.Execute(writer, nil)
}



/*
	Open up a script and produce a list of lines. The only parameter is the
	filename. This returns a slice of the lines.
*/
func OpenScript(filename string) []string {

	// Read the file
	script, err := os.ReadFile(filename)
	// If the file couldn't be opened
	if err != nil {
		// Report the error
		investigator.Report(
			"Unknown file: " + tools.ColouriseMagenta(filename) + ".",
			"n/a",
			"n/a",
			"n/a",
		)
		// Exit the script
		os.Exit(0)
	}
	//fmt.Print(string(script))
	//os.Exit(0)
	// Return the lines of the script
	return strings.Split(string(script), "\n")
}

/*
	Get memory stats. This takes no parameters and returns nothing. Thanks to
	https://reintech.io/blog/introduction-to-gos-runtime-package-memory-
	management-performance
*/
func PrintDevInfo() {
	var mem_stats runtime.MemStats
	runtime.ReadMemStats(&mem_stats)

	// Print out memory info
	fmt.Println(tools.ColouriseYellow("\n\nToken Summary"))
	fmt.Printf(
		tools.ColouriseCyan(":: Total Tokens (incl. line number tokens):") +
		" %d",
		len(values.TOKEN_TREE),
	)

	// Get the size of a single token
	token_memory_size := unsafe.Sizeof(values.TOKEN_TREE[0])
	// Calculate the size of the token tree as a whole
	memory_token_tree := uintptr(cap(values.TOKEN_TREE)) * token_memory_size
	fmt.Printf(
		tools.ColouriseCyan("\n:: Total Memory Usage of TOKEN_TREE:") +
		" %d bytes (single token: %d bytes)",
		memory_token_tree,
		token_memory_size,
	)
	
	// Print out memory info
	fmt.Println(tools.ColouriseYellow("\n\nMemory Information"))
	// Print out the allocated memory
	fmt.Printf(
		tools.ColouriseCyan(":: Allocated Memory: ") +
		"%d bytes, %d kilobytes\n",
		mem_stats.Alloc,
		(mem_stats.Alloc/1024),
	)
	// Print out the total allocated memory
	fmt.Printf(
		tools.ColouriseCyan(":: Total Allocated Memory: ") +
		"%d bytes, %d kilobytes\n",
		mem_stats.TotalAlloc,
		(mem_stats.TotalAlloc/1024),
	)
	// Print out the memory requested from the OS
	fmt.Printf(
		tools.ColouriseCyan(":: Memory Requested: ") +
		"%d bytes, %d kilobytes\n",
		mem_stats.Sys,
		(mem_stats.Sys/1024),
	)
	// Print out the garbage collections
	fmt.Printf(
		tools.ColouriseCyan(":: Garbage Collections: ") + "%d\n",
		mem_stats.NumGC,
	)
}

/*
	The main function. No parameters and no returns.
*/
func main() {

	// Allow the user to execute system commands, defaults to false
	allowexec_flag := flag.Bool(
		"allowexec",
		false,
		"Allow execution of system commands.",
	)

	// Create a template script to work from
	create_template_flag := flag.String(
		"create",
		"",
		"Create a script at the path specified.",
	)
	
	// Print out developer relevant information
	dev_flag := flag.Bool(
		"dev",
		false,
		"[Dev] See information relevant for developing the interpreter. If " +
		"you are an enduser, this information will not be helpful.",
	)

	// Serve up the documentation
	docs_flag := flag.Bool(
		"docs",
		false,
		"Serve up documentation for the language on port 8000.",
	)

	// Time the execution of the script
	timer_flag := flag.Bool(
		"timer",
		false,
		"[Dev] Time the execution of the script.",
	)
	// Get whether we are being verbose with output or not, defaults to false
	verbose_flag := flag.Bool(
		"verbose",
		false,
		"Verbose mode",
	)
	// Get the version of the app
	version_flag := flag.Bool(
		"version",
		false,
		"The version of the interpreter/language.",
	)

	// Parse the flags
	flag.Parse()

	// Set up a start timer
	var start time.Time
	// If the timer flag is true, start a timer
	if *timer_flag {
		// Get a start time for execution speed
		start = time.Now()
	}

	// If the create flag is passed...
	if *create_template_flag != "" {
		// Create the template with a minver and a shebang line
		template_script := []byte("#!/usr/bin/appetit\nminver " +
			strconv.Itoa(values.LANG_VERSION) + "\n\n- Say hello to the " +
			"world\nwriteln \"Hello World!\"\n",
		)
		
		// Dereference the *create_template_flag
		file_name := *create_template_flag
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
				*create_template_flag + ". Make sure that you can save a " +
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
				*create_template_flag + ". Make sure that you can save a " +
				"file in that location.",
			)
		}

		// Report back that we're done
		fmt.Println(
			":: Created a script at " +
			tools.ColouriseCyan(*create_template_flag),
		)
		// Abandon ship
		os.Exit(0)
	}

	// If the docs flag is passed, serve up the docs
	if *docs_flag {
		// Set the port
		port := "8000"
		// Print the port that the documentation is being served on
		fmt.Println(
			"Open up " + tools.ColouriseCyan("http://localhost:" + port) +
			" in your browser.",
		)
		fmt.Println(
			"Press " + tools.ColouriseMagenta("Ctrl-C") + " to quit the " +
			"server.",
		)
		// Set up a handler
		http.HandleFunc("/", DocsHandler)
		// Serve the documentation
		http.ListenAndServe(":" + port, nil)
	}

	// If the version flag is true, print version info
	if *version_flag {
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
			"%s\nInstalled to %s\n\n%s\n\t%s%s\n\t%s%s\n\t%s%d\n\n%s\n\t%s%s\n\t%s%s\n",
			tools.ColouriseMagenta(
				values.LANG_NAME + " " +
				strconv.Itoa(int(values.LANG_VERSION)),
			),
			tools.ColouriseBlue(bin_dir),
			tools.ColouriseYellow("[Platform]"),
			tools.ColouriseCyan("Operating System: "),
			runtime.GOOS,
			tools.ColouriseCyan("Architecture: "),
			runtime.GOARCH,
			tools.ColouriseCyan("CPUs: "),
			runtime.NumCPU(),
			tools.ColouriseYellow("[Build]"),
			tools.ColouriseGreen("Go Version: "),
			go_version,
			tools.ColouriseGreen("Build Date: "),
			BuildDate,
		)
		/* Set up the struct that will hold the JSON data for remotely checking
			the current version.
		*/
		type RemoteDetails struct {
			Version	int	`json:"version"`
			Date	string	`json:"date"`
			Description	string	`json:"description"`
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
		/* If the remote file has a version number greater than the current
			version, inform the user and give them some details.
		*/
		if remote_data.Version > values.LANG_VERSION {
			fmt.Printf(
				"\n%s\nThere's a new version of Appetit " +
				"available! Version %s is available, released %s. " +
				"%s\n",
				tools.ColouriseYellow("[New Version]"),
				strconv.Itoa(remote_data.Version),
				remote_data.Date,
				remote_data.Description,
			)
		/* If the version information is the same, let them know that they are
			running the most current version.
		*/
		} else if remote_data.Version == values.LANG_VERSION {
			fmt.Println("\nYou're up to date!")
		}
		// Exit the app
		os.Exit(0)
	}

	// Set the output to verbose
	values.MODE_VERBOSE = *verbose_flag

	// Set the allow exec setting
	values.ALLOW_EXEC = *allowexec_flag

	// Get the file name
	file_name := flag.Args()
	// If there are no tailing arguments (ie. the file name)
	if len(file_name) == 0 {
		// Error out
		investigator.ReportSimple(
			"You need to pass a script name to the interpreter.",
		)
	}

	/* Before we start parsing, set any reserved variables that require
		"computation".
	*/
	values.BuildReservedVariables()

	// Open up the script
	contents := OpenScript(file_name[0])
	// Remove the comments from the script
	contents = parser.RemoveComments(contents)
	/* Do a quick check to make sure that the minver statement call, if
		present, is the first language specific call.
	*/
	valid_minver, message := parser.CheckValidMinverLocationCount(contents)
	// If it's not appropriately located in the script, error out
	if !valid_minver {
		investigator.ReportSimple(message)
	}

	/* If the dev flag is set, run the script and report back developer
		information.
	*/
	if *dev_flag {
		// Start printing out the tokens
		fmt.Println(tools.ColouriseYellow("\nTokens"))
		parser.Start(contents, true)
	} else {
		parser.Start(contents, false)
	}

	// If the timer flag is true, print the results
	if *timer_flag {
		// Get the time now for calculating the end
		end := time.Now()

		// Get the time
		total_running_time := end.Sub(start)
		fmt.Println(tools.ColouriseCyan("\nTotal Running Times"))
		fmt.Println("\tReported value: " + total_running_time.String())
		fmt.Println(
			"\tRounded (millisecond): " +
			time.Since(start).Round(time.Millisecond).String(),
		)
		fmt.Println(
			"\tRounded (nanosecond): " +
			time.Since(start).Round(time.Nanosecond).String(),
		)
	}

	// If the developer flag is set, print out developer information.
	if *dev_flag {
		PrintDevInfo()
	}
		
}
