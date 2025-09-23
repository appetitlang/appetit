/*
The testing module holds values that are used in testing.
*/
package values

// A simple set of tokens for a minver call
var TEST_MINVER = []Token{
	{
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "-1",
                TokenValue: "",
                TokenType: "string",
        },
        {
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "0",
                TokenValue: "minver",
                TokenType: "string",
        },
        {
                FullLineOfCode: "minver 1",
                LineNumber: 1,
                TokenPosition: "7",
                TokenValue: "1",
                TokenType: "string",
        },
}

// A simple Hello World! set of tokens for a writeln call
var TEST_SET = []Token{
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "-1",
			TokenValue: "",
			TokenType: "string",
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "set",
			TokenType: "string",
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "4",
			TokenValue: "name",
			TokenType: "string",
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "9",
			TokenValue: "=",
			TokenType: "string",
	},
	{
			FullLineOfCode: "set name = \"Hello World!\"",
			LineNumber: 1,
			TokenPosition: "11",
			TokenValue: "\"Hello World!\"",
			TokenType: "string",
	},
}

// A simple ask statement call
var TEST_ASK = []Token{
	 {
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "-1",
			TokenValue: "",
			TokenType: "string",
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "ask",
			TokenType: "string",
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "4",
			TokenValue: "\"Greeting: \"",
			TokenType: "string",
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "17",
			TokenValue: "to",
			TokenType: "string",
	},
	{
			FullLineOfCode: "ask \"Greeting: \" to greeting",
			LineNumber: 1,
			TokenPosition: "20",
			TokenValue: "greeting",
			TokenType: "string",
	},
}

// A simple Hello World! set of tokens for a writeln call
var TEST_WRITELN = []Token{
	{
			FullLineOfCode: "writeln \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "-1",
			TokenValue: "",
			TokenType: "int",
	},
	{
			FullLineOfCode: "writeln \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "0",
			TokenValue: "writeln",
			TokenType: "string",
	},
	{
			FullLineOfCode: "writeln \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "8",
			TokenValue: "\"Hello World!\"",
			TokenType: "string",
	},
}

// A simple Hello World! set of tokens for a writeln call
var TEST_WRITE = []Token{
	{
			FullLineOfCode: "write \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "-1",
			TokenValue: "",
			TokenType: "int",
	},
	{
			FullLineOfCode: "write \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "0",
			TokenValue: "writeln",
			TokenType: "string",
	},
	{
			FullLineOfCode: "write \"Hello World!\"",
			LineNumber: 5,
			TokenPosition: "8",
			TokenValue: "\"Hello World!\"",
			TokenType: "string",
	},
}