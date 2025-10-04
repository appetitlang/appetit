package parser

import (
	"appetit/values"
	"testing"
)

// Test the Tokenise() function.
func TestTokenise(t *testing.T) {
	sample_tokens := values.Token{
                FullLineOfCode: "writeln \"Hello World!\"",
                LineNumber: 1,
                TokenPosition: "0",
                TokenValue: "",
                TokenType: "string",
                NonCommentLineNumber: 1,
	}

        results := Tokenise("writeln \"Hello World!\"", 1, 1)

        if results[0] != sample_tokens {
                t.Errorf(
                        "Tokenisation failed, got %v, expected %v",
                        results[0],
                        sample_tokens,
                )
        }
}

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