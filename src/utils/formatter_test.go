package utils

import "testing"

/*
This is a simple test to ensure that the CommaSeperator() function is working.
*/
func TestCommaSeperator(t *testing.T) {
	// Set up some samples to test
	num_one := 1234567890
	num_one_seperated := "1,234,567,890"

	num_two := 12345
	num_two_seperated := "12,345"

	num_three := 123
	num_three_seperated := "123"

	// Comma seperate the three values
	result_one := CommaSeperator(float64(num_one))
	result_two := CommaSeperator(float64(num_two))
	result_three := CommaSeperator(float64(num_three))

	// Report back any errors
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
