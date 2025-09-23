package parser

import (
	"appetit/values"
	"testing"
)

// Test the Tokeniser() function.
func TestTokeniser(t *testing.T) {
	sample_tokens := []values.Token{
		{
                FullLineOfCode: "writeln \"Hello World!\"",
                LineNumber: 1,
                TokenPosition: "0",
                TokenValue: "",
                TokenType: "string",
        },
        {
                FullLineOfCode: "writeln \"Hello World!\"",
                LineNumber: 1,
                TokenPosition: "1",
                TokenValue: "writeln",
                TokenType: "string",
        },
        {
                FullLineOfCode: "writeln \"Hello World!\"",
                LineNumber: 1,
                TokenPosition: "9",
                TokenValue: "\"Hello World!\"",
                TokenType: "string",
        },
	}

        results := Tokeniser(1, "writeln \"Hello World!\"")

        for index := range results {
                if results[index] != sample_tokens[index] {
                        t.Errorf(
                                "Tokenisation failed, got %v, expected %v",
                                results[index],
                                sample_tokens[index],
                        )
                }
        }

}