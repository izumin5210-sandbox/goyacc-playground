package main

import (
	"text/scanner"
)

type Token struct {
	token   int
	literal string
}

type Tag struct {
	key   string
	value Expr
}

type Expr interface {
}

type Lexer struct {
	scanner.Scanner
	result []Tag
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := int(l.Scan())
	s := l.TokenText()
	if t == scanner.Int {
		t = INT
	}
	if t == scanner.Float {
		t = FLOAT
	}
	if t == scanner.String {
		t = STRING
		s = s[1 : len(s)-1]
	}
	if t == scanner.Ident {
		t = KEY
	}
	if s == "," {
		t = COMMA
	}
	if s == ":" {
		t = COLON
	}
	if s == "true" {
		t = TRUE
	}
	if s == "false" {
		t = FALSE
	}

	lval.token = Token{token: t, literal: s}

	return t
}

func (l *Lexer) Error(e string) {
	panic(e)
}
