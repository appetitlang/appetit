package parser

import (
	"fmt"
	"slices"
	"testing"
)

/*
Check to make sure that the CalculateValue() function (a) both correctly
calculates maths expressions and; (b) ignores expressions that can't be
calculated.
*/
func TestCalculateValue(t *testing.T) {
	// Three sample expressions, the first two can be calculated, the last not
	expression_one := "1+4"
	expression_two := "(23+10)*3"
	expression_three := fmt.Sprintf("%s v%d", LANG_NAME, LANG_VERSION)

	// Calculate those expressions
	calculated_one := CalculateValue("1", expression_one)
	calculated_two := CalculateValue("1", expression_two)
	calculated_three := CalculateValue("1", expression_three)

	// Do the checks and error out if need be
	if calculated_one != "5" {
		t.Errorf("[CalculateValue] Expected %s, got %s",
			"5",
			calculated_one)
	}

	if calculated_two != "99" {
		t.Errorf("[CalculateValue] Expected %s, got %s",
			"99",
			calculated_two)
	}

	if calculated_three != expression_three {
		t.Errorf("[CalculateValue] Expected %s, got %s",
			expression_three,
			calculated_three)
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
