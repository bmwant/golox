package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	Source  string
	Start   int
	Current int
	Line    int
	Tokens  []Token
}

func NewScanner(source string) Scanner {
	return Scanner{
		Source:  source,
		Start:   0,
		Current: 0,
		Line:    1,
		Tokens:  make([]Token, 0),
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) scanTokens() {
	// tokens := make([]Token)

	for !s.isAtEnd() {
		s.Start = s.Current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, NewToken(EOF, "", "", s.Line))
}

func (s *Scanner) ScanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ':
		break
	case '\r':
		break
	case '\t':
		break
	case '\n':
		s.Line++
	case '"':
		s.string()
	default:
		if unicode.IsDigit(rune(c)) {
			s.number()
		} else if unicode.IsLetter(rune(c)) {
			s.identifier()
		} else {
			fmt.Println("Unexpected character at line", s.Line)
		}
	}
}

func (s *Scanner) advance() byte {
	ch := s.Source[s.Current]
	s.Current++
	return ch
}

func (s *Scanner) identifier() {
	for unicode.IsLetter(rune(s.peek())) {
		s.advance()
	}
	text := s.Source[s.Start:s.Current]
	value, ok := keywords[text]
	if !ok {
		value = IDENTIFIER
	}
	s.addToken(value)
}

func (s *Scanner) number() {
	for unicode.IsDigit(rune(s.peek())) {
		s.advance()
	}
	if s.peek() == '.' && unicode.IsDigit(rune(s.peekNext())) {
		s.advance()

		for unicode.IsDigit(rune(s.peek())) {
			s.advance()
		}
	}

	f, _ := strconv.ParseFloat(s.Source[s.Start:s.Current], 64)
	s.addToken(NUMBER, f)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		fmt.Println("Unterminated string at", s.Line)
		return
	}
	s.advance()
	value := s.Source[s.Start+1 : s.Current-1]
	s.addToken(STRING, value)

}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		// Thats how you do this in go
		return '\000'
	}
	return s.Source[s.Current]
}

func (s *Scanner) peekNext() byte {
	if s.Current+1 >= len(s.Source) {
		return '\000'
	}
	return s.Source[s.Current+1]
}

func (s *Scanner) match(char byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Source[s.Current] != char {
		return false
	}
	s.Current++
	return true
}

func (s *Scanner) addToken(tokenType TokenType, extra ...interface{}) {
	var literal interface{}
	if len(extra) > 0 {
		literal = extra[0]
	}
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, NewToken(tokenType, text, literal, s.Line))
}
