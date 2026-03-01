package parser

import (
	"testing"
)

/*
Benchmarking the Call() function. Run this against the TEST_SET of tokens as
this doesn't produce any output.
*/
func BenchmarkCall(b *testing.B) {
	//q := quiet()
	//defer q()

	for b.Loop() {
		Call(TEST_SET)
	}
}

/*
Benchmarking the Start() function. Run this against a set statement as this
doesn't produce any output.
*/
func BenchmarkStart(b *testing.B) {
	//q := quiet()
	//defer q()
	for b.Loop() {
		lines := []string{
			"set name = \"Hello\"",
		}
		Start(lines, false)
	}
}

/*
Benchmarking the Tokenise() function. Run this against a set statement as this
doesn't produce any output.
*/
func BenchmarkTokenise(b *testing.B) {
	for b.Loop() {
		Tokenise(
			"set greeting = \"Hello World!\"",
			1,
			1,
		)
	}
}
