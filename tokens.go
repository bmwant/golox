package main

const (
	LEFT_PAREN = iota
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
	TokenType int
	Lexeme    string
	Literal   string
	Line      int
}

func NewToken(tokenType int, lexeme string, literal string, line int) Token {
	return Token{
		tokenType,
		lexeme,
		literal,
		line,
	}
}
