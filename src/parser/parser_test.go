package parser

import (
	"appetit/values"
	"reflect"
	"testing"
)

func TestValidMinverCalls(t *testing.T) {
	simple_line_sets := []string{
			"minver 1",
			"writeln \"Hello World!\"",
	}

	duplicate_minvers := []string{
			"minver 1",
			"minver 3",
			"write \"Hello World!\"",
	}

	incorrect_order_minvers := []string{
			"write \"Hello World!\"",
			"pause 3",
			"minver 1",
	}

	// Check a basic minver check
	simple_minver_check, _ := CheckValidMinverLocationCount(
			simple_line_sets,
	)
	// Check that it calls duplicates of minver calls
	duplicate_minver_check, _ := CheckValidMinverLocationCount(
			duplicate_minvers,
	)
	/* Check for incorrectly ordered statement (ie. wrongly placed minver
			calls)
	*/
	incorrect_order_minver_check, _ := CheckValidMinverLocationCount(
			incorrect_order_minvers,
	)

	if !simple_minver_check {
			t.Errorf(
					"MinverCheck failed (simple), got false, " +
					"expected true",
			)
	}

	if duplicate_minver_check {
			t.Errorf(
					"MinverCheck failed (duplicate), got false, " + 
					"expected true",
			)
	}

	if incorrect_order_minver_check {
			t.Errorf(
					"MinverCheck failed (not first line), got false, " + 
					"expected true",
			)
	}
}

func TestValidWriteCall(t *testing.T) {
	results := Tokenise("write \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, values.TEST_WRITE)

	if !tokenisation_equal {
		t.Errorf(
				"[write stmt] Tokenisation failed, got %v, expected %v",
				results,
				values.TEST_WRITE,
		)
	}
}

func TestValidWriteLnCall(t *testing.T) {
	results := Tokenise("writeln \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, values.TEST_WRITELN)

	if !tokenisation_equal {
		t.Errorf(
				"[writeln stmt] Tokenisation failed, got %v, expected %v",
				results,
				values.TEST_WRITELN,
		)
	}
}

func TestValidAskCall(t *testing.T) {
	results := Tokenise("ask \"Greeting: \" to greeting", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, values.TEST_ASK)

	if !tokenisation_equal {
		t.Errorf(
				"[ask stmt] Tokenisation failed, got %v, expected %v",
				results,
				values.TEST_ASK,
		)
	}
}

func TestValidSetCall(t *testing.T) {
	results := Tokenise("set name = \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, values.TEST_SET)

	if !tokenisation_equal {
		t.Errorf(
			"[set stmt] Tokenisation failed, got %v, expected %v",
			results,
			values.TEST_SET,
		)
	}
}

func TestValidMinVerCall(t *testing.T) {
	results := Tokenise("minver 1", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, values.TEST_MINVER)

	if !tokenisation_equal {
		t.Errorf(
			"[minver stmt] Tokenisation failed, got %v, expected %v",
			results,
			values.TEST_SET,
		)
	}
}