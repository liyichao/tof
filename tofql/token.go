package tofql

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT
	INTEGER
	FLOAT
	SCINUMBER
	STRING
	NUMBER
	TRUE
	FALSE

	// Symbols
	LPAREN    // (
	RPAREN    // )
	COMMA     // ,
	EQUAL     // =
	LBRACE    // {
	RBRACE    // }
	LBRACEKET // [
	RBRACEKET // ]
	BACKSLASH // \

)
