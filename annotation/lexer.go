package main

import (
	"io"
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result []Annotation
}

var symbolMap = map[rune]int{
	'(':  LPAREN,
	')':  RPAREN,
	'{':  LBRACE,
	'}':  RBRACE,
	':':  COLON,
	',':  COMMA,
	'=':  EQUAL,
	'\n': CR,
}

func (l *Lexer) Init(src io.Reader) *scanner.Scanner {
	return l.Scanner.Init(src)
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := int(l.Scan())
	s := l.TokenText()

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

	lval.token = Token{Token: t, Literal: s}

	return t
}

func (l *Lexer) Error(e string) {
	// err
}
