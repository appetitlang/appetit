/*
The testing module holds values that are used in testing. This includes
statement calls.
*/
package values

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

// A simple copydirectory statement call
var TEST_COPYDIR = []Token{
	 {
			FullLineOfCode: "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "copydirectory",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "15",
			TokenValue: "\"/home/user/test\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "33",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copydirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "36",
			TokenValue: "\"/home/user/test2\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple copyfile statement call
var TEST_COPYFILE = []Token{
	 {
			FullLineOfCode: "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "copyfile",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
			LineNumber: 1,
			TokenPosition: "10",
			TokenValue: "\"/home/user/test.txt\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
			LineNumber: 1,
			TokenPosition: "32",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"",
			LineNumber: 1,
			TokenPosition: "35",
			TokenValue: "\"/home/user/test2.txt\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple deletedirectory statement call
var TEST_DELETEDIR = []Token{
	 {
			FullLineOfCode: "deletedirectory \"/home/user/test/\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "deletedirectory \"/home/user/test/\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "deletedirectory",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "deletedirectory \"/home/user/test/\"",
			LineNumber: 1,
			TokenPosition: "17",
			TokenValue: "\"/home/user/test/\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple deletefile statement call
var TEST_DELETEFILE = []Token{
	 {
			FullLineOfCode: "deletefile \"/home/user/test.txt\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "deletefile \"/home/user/test.txt\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "deletefile",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "deletefile \"/home/user/test.txt\"",
			LineNumber: 1,
			TokenPosition: "12",
			TokenValue: "\"/home/user/test.txt\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple downlaod statement call
var TEST_DOWNLOADFILE = []Token{
	 {
			FullLineOfCode: "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "download",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
			LineNumber: 1,
			TokenPosition: "10",
			TokenValue: "\"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
			LineNumber: 1,
			TokenPosition: "196",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "download \"http://upload.wikimedia.org/wikipedia/commons/0/02/La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024.jpg\" to \"#b_home/Desktop/del.jpg\"",
			LineNumber: 1,
			TokenPosition: "199",
			TokenValue: "\"#b_home/Desktop/del.jpg\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

var TEST_EXECUTE = []Token{
	{
		FullLineOfCode: "execute \"ls -l\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "execute \"ls -l\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "execute",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "execute \"ls -l\"",
		LineNumber: 1,
		TokenPosition: "9",
		TokenValue: "\"ls -l\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}

// A simple exit statement call
var TEST_EXIT = []Token{
	 {
			FullLineOfCode: "exit",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "exit",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "exit",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple makedirectory call
var TEST_MAKEDIR = []Token{
	{
		FullLineOfCode: "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "makedirectory",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "makedirectory \"#b_home/Downloads/testdir2\"",
		LineNumber: 1,
		TokenPosition: "15",
		TokenValue: "\"#b_home/Downloads/testdir2\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}

// A simple makefile call
var TEST_MAKEFILE = []Token{
	{
		FullLineOfCode: "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "makefile",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "makefile \"#b_home/Downloads/testdir2.txt\"",
		LineNumber: 1,
		TokenPosition: "10",
		TokenValue: "\"#b_home/Downloads/testdir2.txt\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}

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

// A simple movedir statement call
var TEST_MOVEDIR = []Token{
	{
			FullLineOfCode: "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "movedirectory",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "15",
			TokenValue: "\"/home/user/test\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "33",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movedirectory \"/home/user/test\" to \"/home/user/test2\"",
			LineNumber: 1,
			TokenPosition: "36",
			TokenValue: "\"/home/user/test2\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple movefile call
var TEST_MOVEFILE = []Token{
	 {
			FullLineOfCode: "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "movefile",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
			LineNumber: 1,
			TokenPosition: "10",
			TokenValue: "\"/home/user/test.txt\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
			LineNumber: 1,
			TokenPosition: "32",
			TokenValue: "to",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "movefile \"/home/user/test.txt\" to \"/home/user/test\"",
			LineNumber: 1,
			TokenPosition: "35",
			TokenValue: "\"/home/user/test\"",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

var TEST_PAUSE = []Token{
	{
			FullLineOfCode: "pause 3",
			LineNumber: 1,
			TokenPosition: "0",
			TokenValue: "",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "pause 3",
			LineNumber: 1,
			TokenPosition: "1",
			TokenValue: "pause",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
	{
			FullLineOfCode: "pause 3",
			LineNumber: 1,
			TokenPosition: "7",
			TokenValue: "3",
			TokenType: "string",
			NonCommentLineNumber: 1,
	},
}

// A simple set call
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

// A simple write call
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

// A simple writeln call
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

// A simple zipdirectory call
var TEST_ZIPDIR = []Token{
	{
		FullLineOfCode: "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "zipdirectory",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "14",
		TokenValue: "\"/home/user/test\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "32",
		TokenValue: "to",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "35",
		TokenValue: "\"/home/user/test2.zip\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}

// A simple zipdirectory call
var TEST_ZIPFILE = []Token{
	{
		FullLineOfCode: "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "0",
		TokenValue: "",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "1",
		TokenValue: "zipfile",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "9",
		TokenValue: "\"/home/user/test.txt\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "31",
		TokenValue: "to",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
	{
		FullLineOfCode: "zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"",
		LineNumber: 1,
		TokenPosition: "34",
		TokenValue: "\"/home/user/test2.zip\"",
		TokenType: "string",
		NonCommentLineNumber: 1,
	},
}
