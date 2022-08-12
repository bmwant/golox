package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var keywords = map[string]int{
	"and":    AND,
	"class":  AND,
	"else":   AND,
	"false":  AND,
	"for":    AND,
	"fun":    AND,
	"if":     AND,
	"nil":    AND,
	"or":     AND,
	"print":  AND,
	"return": AND,
	"super":  AND,
	"this":   AND,
	"true":   AND,
	"var":    AND,
	"while":  AND,
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
	return ' '
}

func (s *Scanner) identifier() {
	for unicode.IsLetter(rune(s.peek())) {
		s.advance()
	}
	s.addToken(IDENTIFIER)
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
	// todo: update addToken definition to accept f
	_ = f
	s.addToken(NUMBER)
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
	// value = s.Source[]
	s.addToken(STRING)

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
func (s *Scanner) addToken(tokenType int) {
	text := "The token text"
	s.Tokens = append(s.Tokens, NewToken(tokenType, text, "", s.Line))
}
