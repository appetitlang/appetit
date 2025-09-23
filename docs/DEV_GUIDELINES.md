## Developer Guidelines

### Release Build
1. First, make sure that all the tests pass. Run `make test` or `@go test ./...` in the `src/` directory.
2. Increment version number in `Makefile` and `Make.ps1`
3. Run `Makefile` and `Make.ps1`
4. Profit


### Code
#### Formatting
* Lines of code are 80 characters wide

#### Path Organisation
* investigator - this deals with checks and errors throughout the code
* parser - this deals with tokenising the lines of the script and delegating to the statements module
* statements - this deals with executing functionality for the statements
* tools - this houses miscellaneous tools and functions that are used throughout the interpreter
* values - this deals with holding values for use across the interpreter ranging from the version number through to the struct that defines a token

#### Errors
* Errors should include the full line of code where the error is triggered by a malformed line of code, not just that the execution failed. For instance, `writeln "Hello` should trigger an error with a line of code given the missing quotation marks.
* Errors also need to be user friendly. It should be pointed as clear about what the problem is or might be. A cryptic message like `buffer overflow` or `permissions on dir are 0655, not 0755` are not helpful for those who don't know what a buffer is or what Unix style permissions are. Remember that the user is a human first, not a programmer first. If you have something that might cause an error, write any messages accordingly including, where possible, a potential explanation. Always include the full line of code.

#### Naming
* We don't follow Go's preference for exceptionally short variable names. Variables have descriptive names.
* CapWords is used for functions (ie. every word is capitalised) to accord with Go conventions (capitalised functions are public) and readability. Being in a different case scheme relative to variables makes it easy to sight the differences.
* Spellings conform to Australian English conventions. For instance, the `colourise` module is spelled correctly (ie. it's not colorize). This is non-negotiable to ensure consistency and predicatbility across the codebase.

#### Colouring Output
* Ensure consistency with colourised output
    * All error headers/footers are urgent messages should be ColouriseRed
    * All statement names should be ColouriseCyan
    * All strings in sample code should be ColouriseGreen
    * All values in errors should be ColouriseYellow
    * All else including full lines of code should be ColouriseMagenta
This is a work in progress so please do feel free to suggest errors where appropriate.