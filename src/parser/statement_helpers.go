/*
This houses functions used directly by statement functions as supports. In this
way, these are functions that are needed to make statements work directly and
are thus a 'dependency' of those functions.
*/
package parser

import (
	"appetit/utils"
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
/*
copydirectory statement helpers
*/

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
	source_position := token_info["source_position"]
	dest_position := token_info["dest_position"]

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
			if MODE_VERBOSE {
				fmt.Println(
					":: Making " + utils.ColouriseGreen(relative_path) +
						"...",
				)
			}
			// Make the path with some sensible permissions
			os.MkdirAll(dest+relative_path, 0750)
		} else {
			// Get the name of the file to copy
			file_to_copy := source + relative_path
			// Open the source file
			source_file, source_err := os.Open(file_to_copy)
			// If there is an error opening the source file, report that
			if source_err != nil {
				Report(
					"Can't open "+utils.ColouriseYellow(file_to_copy)+
						". Perhaps you don't have read permissions? "+
						source_err.Error(),
					loc,
					source_position,
					full_loc,
				)
			}
			// Establish where we are creating the files
			create_path := dest + relative_path
			// Create the new file
			create, create_err := os.Create(create_path)
			// If there was an error in creating the new file, report that
			if create_err != nil {
				Report(
					"Couldn't create the file in "+
						utils.ColouriseYellow(create_path)+". Check that "+
						"you have write permissions to write to "+
						utils.ColouriseYellow(dest)+" and/or that there is "+
						"enough space available for you to copy the file(s) "+
						"over.",
					loc,
					dest_position,
					full_loc,
				)
			}
			/* If verbose mode is set, note that we are copying a file and
			report back the file size
			*/
			if MODE_VERBOSE {
				fmt.Printf(
					"    :: Copying %s %s...",
					utils.ColouriseGreen(info.Name()),
					utils.ColouriseMagenta(
						"["+strconv.FormatInt(info.Size(), 10)+
							" bytes]",
					),
				)
			}

			// Copy the file itself
			_, copy_err := io.Copy(create, source_file)
			// If there's an error, report it
			if copy_err != nil {
				Report(
					"There was an error doing the copy for "+
						source_file.Name(),
					loc,
					"n/a",
					full_loc,
				)
			}
			// If verbose mode is set, note that we are done copying the file
			if MODE_VERBOSE {
				fmt.Println("done.")
			}

		}
		/* We return no error here as it is assumed that any errors are handled
		above
		*/
		return nil
	}
}

// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
/*
download statement helpers
*/

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
	aren't using fmt here is to ensure that text can be flushed from the buffer
	properly which allows writing over the lines cleanly.
	*/
	writer := bufio.NewWriter(os.Stdout)
	// Calculate the percentage
	percentage := wp.TotalBytes / wp.FileSize
	// Format the progress as a percentage for printing.
	progress_output := fmt.Sprintf(
		"\rDownloaded %s (%s KB of %s KB)",
		utils.ColouriseMagenta(
			strconv.FormatFloat(percentage*100, 'f', 2, 32)+"%",
		),
		utils.CommaSeperator(wp.TotalBytes/1024),
		utils.CommaSeperator(wp.FileSize/1024),
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

// ----------------------------------------------------------------------------
