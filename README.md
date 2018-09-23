# Lexical analyzer

This package gets the source code from the input and produces the tokens:

1. Keywords
2. Identifiers
3. Punctuators
4. Operators
5. Literals


## Output format specification

The groups listed above are transformed to the output as follows:

1. "KEYWORD       : int "
2. "IDENTIFIER      : reverse"
3. "PUNCTUATOR : ;"
4. "OPERATOR      : ="
5. "LITERAL           : 10"


## How to run

To run the program you can use the following scripts:
* `go run *`
or
* `go build .`
* `c-sharp-lex-anal`

After that the lexical analysis results will be written to `out.txt`
line by line based on the provided input C Sharp sources in `in.txt`.

In order to run only the implemented Unit tests you can simply use by this command in the project direcetory:

* `go test ./...`
