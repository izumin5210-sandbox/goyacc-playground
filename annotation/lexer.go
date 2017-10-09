package main

import (
	"errors"
	"fmt"
	"io"
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result []Annotation
	state  state
	errors []error
}

var symbolMap = map[rune]int{
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	':': COLON,
	',': COMMA,
	'=': EQUAL,
}

func (l *Lexer) Init(src io.Reader) *scanner.Scanner {
	l.Scanner.Init(src)
	l.state = stateWaitAnnotation
	l.Whitespace = l.state.Whitespace()
	l.Scanner.Error = func(s *scanner.Scanner, msg string) {
		l.errors = append(l.errors, errors.New(msg))
	}
	return &l.Scanner
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := int(l.Scan())
	s := l.TokenText()

	if t == scanner.EOF {
		return t
	}
	if sym, ok := symbolMap[rune(t)]; ok {
		t = sym
	}
	if t == scanner.Ident {
		t = IDENT
	}
	if t == scanner.String {
		t = STRING
		s = s[1 : len(s)-1]
	}
	if t == scanner.Int {
		t = INT
	}
	if t == scanner.Float {
		t = FLOAT
	}
	if s == "true" {
		t = TRUE
	}
	if s == "false" {
		t = FALSE
	}
	l.accept(t, s)

	lval.token = Token{Token: t, Literal: s}

	return t
}

func (l *Lexer) Error(e string) {
	l.Scanner.ErrorCount++
	l.Scanner.Error(&(l.Scanner), e)
}

func (l *Lexer) accept(input int, str string) {
	nextState := l.state.NextState(input)
	if nextState == stateUnknown {
		l.Error(fmt.Sprintf("unexpected token %v", input))
	}
	l.state = nextState
	l.Whitespace = nextState.Whitespace()
}
