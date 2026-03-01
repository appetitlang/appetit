/*
This is the main package that serves as the (conventional) entry point.
Nothing of note here needs to be mentioned that is unique to this package.
*/
package main

import (
	"appetit/parser"
	"appetit/utils"
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/trace"
	"time"
	"unsafe"
)

/* Thanks to https://mblessed.hashnode.dev/go-embed-embed-your-html-frontend-
in-golang for the embed info.
*/

//go:embed docs/book.pdf
var book []byte

/*
Handle serving the documentation with appropriate headers. The parameters
are the conventional response writer and the request. No returns.
*/
func DocsHandler(writer http.ResponseWriter, request *http.Request) {
	// Set the content type
	writer.Header().Set("Content-Type", "application/pdf")
	// Set the content disposition to make the PDF appear inline.
	writer.Header().Set("Content-Disposition", "inline; filename=\"book.pdf\"")
	// Write the book as the content
	writer.Write(book)
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
	fmt.Println(utils.ColouriseYellow("\n\nToken Summary"))
	fmt.Printf(
		utils.ColouriseCyan(":: Total Tokens (incl. line number tokens):")+
			" %s",
		utils.CommaSeperator(float64(len(parser.TOKEN_TREE))),
	)

	// Get the size of a single token
	token_memory_size := unsafe.Sizeof(parser.TOKEN_TREE[0])
	// Calculate the size of the token tree as a whole
	memory_token_tree := uintptr(cap(parser.TOKEN_TREE)) * token_memory_size
	fmt.Printf(
		utils.ColouriseCyan("\n:: Total Memory Usage of TOKEN_TREE:")+
			" %s bytes",
		utils.CommaSeperator(float64(memory_token_tree)),
	)

	// Print out memory info
	fmt.Println(utils.ColouriseYellow("\n\nMemory Information"))
	// Print out the allocated memory
	fmt.Printf(
		utils.ColouriseCyan(":: Allocated Memory: ")+"%s bytes\n",
		utils.CommaSeperator(float64(mem_stats.Alloc)),
	)
	// Print out the total allocated memory
	fmt.Printf(
		utils.ColouriseCyan(":: Total Allocated Memory: ")+"%s bytes\n",
		utils.CommaSeperator(float64(mem_stats.TotalAlloc)),
	)
	// Print out the memory requested from the OS
	fmt.Printf(
		utils.ColouriseCyan(":: Memory Requested: ")+"%s bytes\n",
		utils.CommaSeperator(float64(mem_stats.Sys)),
	)
	// Print out the garbage collections
	fmt.Printf(
		utils.ColouriseCyan(":: Garbage Collections: ")+"%d\n",
		mem_stats.NumGC,
	)
}

/*
The main function. No parameters and no returns.
*/
func main() {
	// Get the version of the app
	about_flag := flag.Bool(
		"about",
		false,
		"Information about the interpreter.",
	)

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
		"[Dev] See information relevant for developing the interpreter. If "+
			"you are an enduser, this information will not be helpful.",
	)

	// Serve up the documentation
	docs_flag := flag.String(
		"docs",
		"",
		"Serve up documentation for the language on port 8000.",
	)

	// Time the execution of the script
	timer_flag := flag.Bool(
		"timer",
		false,
		"[Dev] Time the execution of the script.",
	)

	// Run a trace
	trace_flag := flag.Bool(
		"trace",
		false,
		"[Dev] Execute a runtime trace on the interpreter.",
	)

	// Get whether we are being verbose with output or not, defaults to false
	verbose_flag := flag.Bool(
		"verbose",
		false,
		"Verbose mode",
	)

	// Parse the flags
	flag.Parse()

	if *trace_flag {
		// Get the date and time
		date_now := time.Now()
		// Format the date in YYYY-MM-DD format
		date := fmt.Sprintf(
			"%d-%d-%d-%d-%d-%d",
			date_now.Year(),
			date_now.Month(),
			date_now.Day(),
			date_now.Hour(),
			date_now.Minute(),
			date_now.Second(),
		)
		// Create the trace file name
		trace_file_name := date + "-function_trace.out"
		// Output information about the running of the trace
		fmt.Println(
			utils.ColouriseYellow(":: Saving trace to " + trace_file_name),
		)
		fmt.Print(
			utils.ColouriseYellow(":: When the execution is done, run "),
		)
		fmt.Print(
			utils.ColouriseCyan("go tool trace " + trace_file_name + "\n"),
		)

		// Function tracing here
		trace_file, trace_error := os.Create(trace_file_name)
		// If there is an error creating a trace, error out
		if trace_error != nil {
			fmt.Println("Error creating the trace file")
		}
		// Notify the user if the trace file can't be created
		defer func() {
			if trace_error := trace_file.Close(); trace_error != nil {
				fmt.Println("Failed to close the trace file")
			}
		}()

		// Notify the user if the trace can't be started
		if trace_error := trace.Start(trace_file); trace_error != nil {
			fmt.Println("Failed to start the trace")
		}
		// Defer the stop of the trace
		defer trace.Stop()
	}

	// Set up a start timer
	var start time.Time
	// If the timer flag is true, start a timer
	if *timer_flag {
		// Get a start time for execution speed
		start = time.Now()
	}

	// If the create flag is passed...
	if *create_template_flag != "" {
		utils.CreateTemplate(
			*create_template_flag,
			parser.LANG_VERSION,
		)
	}

	// If the docs flag is passed, serve up the docs
	if *docs_flag != "" {
		// Set the port
		port := *docs_flag
		// Print the port that the documentation is being served on
		fmt.Println(
			"Open up " + utils.ColouriseCyan("http://localhost:"+port) +
				" in your browser.",
		)
		fmt.Println(
			"Press " + utils.ColouriseMagenta("Ctrl-C") + " to quit the " +
				"server.",
		)
		// Set up a handler
		http.HandleFunc("/", DocsHandler)
		// Serve the documentation
		http.ListenAndServe(":"+port, nil)
	}

	// If the version flag is true, print version info
	if *about_flag {
		utils.About(
			parser.LANG_NAME,
			parser.LANG_VERSION,
			parser.LANG_CODENAME,
			parser.BuildDate,
		)
	}

	// Set the output to verbose
	parser.MODE_VERBOSE = *verbose_flag

	// Set the allow exec setting
	parser.MODE_ALLOW_EXEC = *allowexec_flag

	// Get the file name
	file_name := flag.Args()
	// If there are no tailing arguments (ie. the file name)
	if len(file_name) == 0 {
		// Error out
		parser.ReportSimple(
			"You need to pass a script name to the interpreter.",
		)
	}

	// Prep the script by opening it and removing the comments
	contents := parser.PrepScript(file_name[0])

	/* If the dev flag is set, run the script and report back developer
	information.
	*/
	if *dev_flag {
		parser.MODE_DEV = true
		// Start printing out the tokens
		fmt.Println(utils.ColouriseYellow("\nTokens"))
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
		fmt.Println(utils.ColouriseCyan("\nTotal Running Times"))
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
