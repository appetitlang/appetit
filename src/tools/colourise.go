/*
The colourise module provides helpful functions for colouring a string.
*/
package tools

/*
	Set up the ANSI escape characters for shell colouring
	https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
*/
var red string = "\033[31m"
var green string = "\033[32m" 
var yellow string = "\033[33m" 
var blue string = "\033[34m" 
var magenta string = "\033[35m"
var cyan string = "\033[36m" 
var grey string = "\033[37m"
// This is the reset that goes at the end
var reset string = "\033[0m"

/*
	All of the functions below should be self explanatory. In effect, they
	just return a "tools.Colourised" version of the string passed to them.
*/

/*
	Convert text to red. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseRed(text string) string {
	return red + text + reset
}

/*
	Convert text to green. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseGreen(text string) string {
	return green + text + reset
}

/*
	Convert text to yellow. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseYellow(text string) string {
	return yellow + text + reset
}

/*
	Convert text to blue. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseBlue(text string) string {
	return blue + text + reset
}

/*
	Convert text to magenta. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseMagenta(text string) string {
	return magenta + text + reset
}

/*
	Convert text to cyan. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseCyan(text string) string {
	return cyan + text + reset
}

/*
	Convert text to grey. Parameters include the text to convert. Returns a
	colourised string.
*/
func ColouriseGrey(text string) string {
	return grey + text + reset
}