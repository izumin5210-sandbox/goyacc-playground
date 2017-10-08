package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_yyParse(t *testing.T) {
	cases := []struct {
		input  string
		output []Annotation
	}{
		{
			input: "test:foo",
			output: []Annotation{
				{
					Namespace: Token{Token: IDENT, Literal: "test"},
					Name:      Token{Token: IDENT, Literal: "foo"},
				},
			},
		},
		{
			input: `test:foo(bar = "baz")`,
			output: []Annotation{
				{
					Namespace: Token{Token: IDENT, Literal: "test"},
					Name:      Token{Token: IDENT, Literal: "foo"},
					Fields: []Field{
						{
							Name: Token{Token: IDENT, Literal: "bar"},
							Expr: Token{Token: STRING, Literal: "baz"},
						},
					},
				},
			},
		},
		{
			input: `test:foo(bar = "baz", qux = true)`,
			output: []Annotation{
				{
					Namespace: Token{Token: IDENT, Literal: "test"},
					Name:      Token{Token: IDENT, Literal: "foo"},
					Fields: []Field{
						{
							Name: Token{Token: IDENT, Literal: "bar"},
							Expr: Token{Token: STRING, Literal: "baz"},
						},
						{
							Name: Token{Token: IDENT, Literal: "qux"},
							Expr: Token{Token: TRUE, Literal: "true"},
						},
					},
				},
			},
		},
		{
			input: `test:foo(bar = {"baz", "qux"})`,
			output: []Annotation{
				{
					Namespace: Token{Token: IDENT, Literal: "test"},
					Name:      Token{Token: IDENT, Literal: "foo"},
					Fields: []Field{
						{
							Name: Token{Token: IDENT, Literal: "bar"},
							Expr: []Expr{
								Token{Token: STRING, Literal: "baz"},
								Token{Token: STRING, Literal: "qux"},
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		l := new(Lexer)
		l.Init(strings.NewReader(c.input))
		yyParse(l)
		if got, want := len(l.result), len(c.output); got != want {
			t.Errorf("yyParse(%q).result includes %v, want %d tags", c.input, got, want)
		} else {
			for i, wantAnn := range c.output {
				gotAnn := l.result[i]
				if got, want := gotAnn.Namespace, wantAnn.Namespace; !reflect.DeepEqual(got, want) {
					t.Errorf("yyParse(%q).result[%d].Namespace is %v, want %v", c.input, i, got, want)
				}
				if got, want := gotAnn.Name, wantAnn.Name; !reflect.DeepEqual(got, want) {
					t.Errorf("yyParse(%q).result[%d].Name is %v, want %v", c.input, i, got, want)
				}
				if got, want := len(gotAnn.Fields), len(wantAnn.Fields); got != want {
					t.Errorf("yyParse(%q).result[%d].Fiels has %d items, want %d items", c.input, i, got, want)
				} else {
					for j, wantField := range wantAnn.Fields {
						gotField := gotAnn.Fields[j]
						if got, want := gotField.Name, wantField.Name; !reflect.DeepEqual(got, want) {
							t.Errorf("yyParse(%q).result[%d].Fields[%d].Name is %v, want %v", c.input, i, j, got, want)
						}
						if wantExprs, ok := wantField.Expr.([]Expr); ok {
							if gotExprs, ok := gotField.Expr.([]Expr); ok {
								for k, wantExpr := range wantExprs {
									gotExpr := gotExprs[k]
									if got, want := gotExpr, wantExpr; !reflect.DeepEqual(got, want) {
										t.Errorf("yyParse(%q).result[%d].Fields[%d].Exprs[%d] is %v, want %v", c.input, i, j, k, got, want)
									}
								}
							} else {
								t.Errorf("yyParse(%q).result[%d].Fields[%d].Expr is %v, want %v", c.input, i, j, got, want)
							}
						} else {
							if got, want := gotField.Expr, wantField.Expr; !reflect.DeepEqual(got, want) {
								t.Errorf("yyParse(%q).result[%d].Fields[%d].Expr is %v, want %v", c.input, i, j, got, want)
							}
						}
					}
				}
			}
		}
	}
}
