package main

import (
	"strings"
	"testing"
)

func Test_yyParse(t *testing.T) {
	cases := []struct {
		input  string
		errCnt int
	}{
		{
			input:  `foo:"bar"`,
			errCnt: 0,
		},
		{
			input:  `foo:"bar`,
			errCnt: 1,
		},
		{
			input:  `foo: "bar"`,
			errCnt: 1,
		},
		{
			input:  `foo :"bar"`,
			errCnt: 1,
		},
	}

	for _, c := range cases {
		l := new(Lexer)
		l.Init(strings.NewReader(c.input))
		yyParse(l)
		if got, want := l.ErrorCount, c.errCnt; got != want {
			t.Errorf("Caused error count when parsing %q was %d, want %d", c.input, got, want)
		}
	}
}
