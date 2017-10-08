package main

import (
	"io"
	"text/scanner"
)

type Token struct {
	token   int
	literal string
}

type Tag struct {
	key   string
	value Token
}

type Lexer struct {
	scanner.Scanner
	result []Tag
	state  state
}

func (l *Lexer) Init(src io.Reader) *scanner.Scanner {
	l.state = stateDefault
	return l.Scanner.Init(src)
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := int(l.Scan())
	s := l.TokenText()
	if t == scanner.Int {
		l.setState(stateValueDetected)
		t = INT
	}
	if t == scanner.Float {
		l.setState(stateValueDetected)
		t = FLOAT
	}
	if t == scanner.String {
		l.setState(stateValueDetected)
		t = STRING
		s = s[1 : len(s)-1]
	}
	if t == scanner.Ident {
		l.setState(stateKeyDetected)
		t = KEY
	}
	if s == "," {
		l.setState(stateCommaDetected)
		t = COMMA
	}
	if s == ":" {
		l.setState(stateColonDetected)
		t = COLON
	}
	if s == "true" {
		l.setState(stateValueDetected)
		t = TRUE
	}
	if s == "false" {
		l.setState(stateValueDetected)
		t = FALSE
	}

	lval.token = Token{token: t, literal: s}

	return t
}

func (l *Lexer) Error(e string) {
	// err
}

func (l *Lexer) setState(s state) {
	l.state = s
	l.Whitespace = s.whitespace()
}
