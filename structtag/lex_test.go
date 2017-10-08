package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_yyParse(t *testing.T) {
	cases := []struct {
		input    string
		exitCode int
		output   []Tag
	}{
		{
			input:    `foo:"bar"`,
			exitCode: 0,
			output: []Tag{
				{key: "foo", value: Token{token: STRING, literal: "bar"}},
			},
		},
		{
			input:    `foo:"bar",baz:13,qux:true`,
			exitCode: 0,
			output: []Tag{
				{key: "foo", value: Token{token: STRING, literal: "bar"}},
				{key: "baz", value: Token{token: INT, literal: "13"}},
				{key: "qux", value: Token{token: TRUE, literal: "true"}},
			},
		},
		{
			input:    `foo:"bar`,
			exitCode: 1,
			output:   []Tag{},
		},
		{
			input:    `foo: "bar"`,
			exitCode: 1,
			output:   []Tag{},
		},
		{
			input:    `foo :"bar"`,
			exitCode: 1,
			output:   []Tag{},
		},
	}

	for _, c := range cases {
		l := new(Lexer)
		l.Init(strings.NewReader(c.input))
		if got, want := yyParse(l), c.exitCode; got != want {
			t.Errorf("yyParse(%q) was %d, want %d", c.input, got, want)
		}
		if got, want := len(l.result), len(c.output); got != want {
			t.Errorf("yyParse(%q).result includes %v, want %d tags", c.input, got, want)
		} else {
			for i, want := range c.output {
				if got := l.result[i]; !reflect.DeepEqual(got, want) {
					t.Errorf("yyParse(%q).result[%d] is %v, want %v", c.input, i, got, want)
				}
			}
		}
	}
}
