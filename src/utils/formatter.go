/*
The formatter functions deal with formatting text for output. These formatting
functions are non-language specific. For instance, the CommaSeperator()
function just creates a comma seperated string representation of a float.
*/
package utils

import (
	"strconv"
	"strings"
)

/*
This function takes in a float64 and returns a comma seperated version of it as
a string. This is helpful for producing a human readable version of the number
in places such as the download statement. Takes in a float64 and returns a
string.
*/
func CommaSeperator(number float64) string {
	// Get the integer representation of the number
	number_int := int64(number)
	// Create a string version of the number
	string_number := strconv.Itoa(int(number_int))
	// Split the number into seperate characters
	chars := []rune(string_number)

	// Hold the final number
	final_number := ""

	/*
		Count how many digits we've looked at so that we can track where the
		comma needs to go. We're starting here at one as we will be starting
		with our first digit when this is accessed for the first time.
	*/
	comma_counter := 1
	/*
		Loop over the characters starting at the end and working from the end
		back to the beginning.
	*/
	for char_count := len(chars) - 1; char_count >= 0; char_count-- {
		/*
			If the counter is less than 3, we just append the digit to the
			beginning of the final_number.
		*/
		if comma_counter < 3 {
			// Add the digit to the beginning of the final_number
			final_number = string(chars[char_count]) + final_number
			// Increment the counter
			comma_counter++
			/*
				If we've hit a comma_counter value of 3, we need to add in the
				comma and then reset the comma_counter to 1.
			*/
		} else {
			// Add a comma and the digit to the beginning of the final_number
			final_number = "," + string(chars[char_count]) + final_number
			// Reset the counter
			comma_counter = 1
		}
	}
	/*
		If the number of digits is a multiple of 3, a leading comma will be
		there so let's remove it.
	*/
	final_number = strings.TrimLeft(final_number, ",")
	// Return the final comma seperated number.
	return final_number
}
