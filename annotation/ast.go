package main

type Token struct {
	Token   int
	Literal string
}

type Annotation struct {
	Namespace Token
	Name      Token
	Fields    []Field
}

type Field struct {
	Name Token
	Expr Expr
}

type Expr interface {
}
