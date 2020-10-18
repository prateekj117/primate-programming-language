package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Token/character not defined already
	ILLEGAL = "ILLEGAL"

	// End of File
	EOF = "EOF"

	// Identifiers + literals
	INDENT = "INDENT"
	INT    = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
