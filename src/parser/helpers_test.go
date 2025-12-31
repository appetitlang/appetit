package parser

import (
	"slices"
	"testing"
)

func TestCheckValidMinverLocationCount(t *testing.T) {
	valid_sample := []string{
		"minver 1",
		"writeln \"Hello world!\"",
		"ask \"Name: \" to name",
		"writeln \"Hello #name!\"",
	}

	// minver not on the first line
	invalid_sample_one := []string{
		"writeln \"Hello world!\"",
		"minver 1",
		"ask \"Name: \" to name",
		"writeln \"Hello #name!\"",
	}

	// minver provided twice
	invalid_sample_two := []string{
		"minver 1",
		"minver 1",
		"writeln \"Hello world!\"",
		"ask \"Name: \" to name",
		"writeln \"Hello #name!\"",
	}

	valid, _ := CheckValidMinverLocationCount(valid_sample)
	invalid_one, _ := CheckValidMinverLocationCount(invalid_sample_one)
	invalid_two, _ := CheckValidMinverLocationCount(invalid_sample_two)

	if !valid {
		t.Errorf(
			"[CheckValidMinverLocationCount] Valid sample returned false, " +
				"expected true",
		)
	}

	if invalid_one {
		t.Errorf(
			"[CheckValidMinverLocationCount] Invalid location returned " +
				"true, expected false",
		)
	}

	if invalid_two {
		t.Errorf(
			"[CheckValidMinverLocationCount] Invalid count returned " +
				"true, expected false",
		)
	}
}

func TestRemoveComments(t *testing.T) {
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
	after a comment.
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

	comments_removed := RemoveComments(commented_script)

	if !slices.Equal(comments_removed, comments_stripped) {
		t.Errorf(
			"[RemoveComments] Expected %s, got %s",
			comments_stripped,
			comments_removed,
		)
	}
}
