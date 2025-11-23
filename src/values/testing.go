/*
The testing module holds values that are used in testing.
*/
package values

// A simple set of tokens for a minver call
var TEST_MINVER = []Token{
	{
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "0",
                TokenValue: "",
                TokenType: "string",
				NonCommentLineNumber: 1,
        },
        {
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "1",
                TokenValue: "minver",
                TokenType: "string",
				NonCommentLineNumber: 1,
        },
        {
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "8",
                TokenValue: "1",
                TokenType: "string",
				NonCommentLineNumber: 1,
        },
}

// A simple Hello World! set of tokens for a writeln call
var TEST_SET = []Token{
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "set",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "5",
			TokenValue: "name",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "10",
			TokenValue: "=",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "12",
			TokenValue: "\"Hello World!\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple ask statement call
var TEST_ASK = []Token{
	 {
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "ask",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "5",
			TokenValue: "\"Greeting: \"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "18",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "21",
			TokenValue: "greeting",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple Hello World! set of tokens for a writeln call
var TEST_WRITELN = []Token{
	{
		FullLineOfCode: "writeln \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "writeln \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "writeln",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "writeln \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "9",
		TokenValue: "\"Hello World!\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}

// A simple Hello World! set of tokens for a writeln call
var TEST_WRITE = []Token{
	{
		FullLineOfCode: "write \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "write \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "write",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "write \"Hello World!\"",
		LineNumber: 1,
		TokenPosition: "7",
		TokenValue: "\"Hello World!\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}