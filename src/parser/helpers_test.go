package parser

import (
	"slices"
	"testing"
)

/*
This is a simple test to ensure that the CommaSeperator() function is working.
*/
func TestCommaSeperator(t *testing.T) {
	num_one := 1234567890
	num_one_seperated := "1,234,567,890"

	num_two := 12345
	num_two_seperated := "12,345"

	num_three := 123
	num_three_seperated := "123"

	result_one := CommaSeperator(float64(num_one))
	result_two := CommaSeperator(float64(num_two))
	result_three := CommaSeperator(float64(num_three))

	if result_one != num_one_seperated {
		t.Errorf(
			"[CommaSeperator] Expected %s, got %s ",
			num_one_seperated,
			result_one,
		)
	}

	if result_two != num_two_seperated {
		t.Errorf(
			"[CommaSeperator] Expected %s, got %s ",
			num_two_seperated,
			result_two,
		)
	}

	if result_three != num_three_seperated {
		t.Errorf(
			"[CommaSeperator] Expected %s, got %s ",
			num_three_seperated,
			result_three,
		)
	}
}

/*
Check to make sure that the removal of comments does what it should - replace
the comments with empty comment strings.
*/
func TestRemoveComments(t *testing.T) {
	// This is a commented script
	commented_script := []string{
		"- This is a comment",
		"minver 1",
		"writeln \"Hello World!\"",
		"- Ask for a user name",
		"ask \"Name: \" to name",
		"- Write out the user's name",
		"writeln \"Hello #name!\"",
	}

	/* The "stripped" version of a script actually just removes everything
	after a comment. This is what the internal representation of the lines
	should look like after.
	*/
	comments_stripped := []string{
		"-",
		"minver 1",
		"writeln \"Hello World!\"",
		"-",
		"ask \"Name: \" to name",
		"-",
		"writeln \"Hello #name!\"",
	}

	// Remove the comments
	comments_removed := RemoveComments(commented_script)

	// Check to see if the slices are equivalent and if not...
	if !slices.Equal(comments_removed, comments_stripped) {
		// Report an error
		t.Errorf(
			"[RemoveComments] Expected %s, got %s",
			comments_stripped,
			comments_removed,
		)
	}
}
