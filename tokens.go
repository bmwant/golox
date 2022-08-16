package main

//go:generate stringer -type=TokenType
type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	// Literals
	IDENTIFIER
	STRING
	NUMBER
	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   string
	Line      int
}

func NewToken(tokenType TokenType, lexeme string, literal string, line int) Token {
	return Token{
		tokenType,
		lexeme,
		literal,
		line,
	}
}
