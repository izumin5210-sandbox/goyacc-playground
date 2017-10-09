package main

// @startuml
// [*]            -> WaitAnnotation
// WaitAnnotation --> NamespaceDetected : ident
//
// ' annotation
// NamespaceDetected --> ColonDetected : colon
// ColonDetected     --> NameDetected : ident
// NameDetected      --> NamespaceDetected : ident
// NameDetected      --> WaitFieldOrParen : lparen
//
//
// ' fields
// WaitFieldOrParen   --> WaitAnnotation : rparen
// FieldValueDetected --> WaitAnnotation : rparen
// WaitFieldOrParen   --> FieldNameDetected : ident
// FieldNameDetected  --> FieldEqualDetected : equal
// FieldEqualDetected --> FieldValueDetected : literal
// FieldValueDetected --> WaitFieldOrParen : comma
// FieldEqualDetected --> WaitArrayValueOrBrace : lbrace
//
// ' Array
// WaitArrayValueOrBrace --> FieldValueDetected : rbrace
// WaitArrayValueOrBrace --> ArrayValueDetected : literal
// ArrayValueDetected    --> WaitArrayValueOrBrace : comma
// ArrayValueDetected    --> FieldValueDetected : rbrace
// @enduml

type state int

const (
	stateUnknown state = iota
	stateWaitAnnotation
	stateNamespaceDetected
	stateColonDetected
	stateNameDetected
	stateWaitFieldOrParen
	stateFieldNameDetected
	stateFieldEqualDetected
	stateFieldValueDetected
	stateWaitArrayValueOrBrace
	stateArrayValueDetected
)

var stringByState = map[state]string{
	stateUnknown:               "Unknown",
	stateWaitAnnotation:        "WaitAnnotation",
	stateNamespaceDetected:     "NamespaceDetected",
	stateColonDetected:         "ColonDetected",
	stateNameDetected:          "NameDetected",
	stateWaitFieldOrParen:      "WaitFieldOrParen",
	stateFieldNameDetected:     "FieldNameDetected",
	stateFieldEqualDetected:    "FieldEqualDetected",
	stateFieldValueDetected:    "FieldValueDetected",
	stateWaitArrayValueOrBrace: "WaitArrayValueOrBrace",
	stateArrayValueDetected:    "ArrayValueDetected",
}

var nextStateByInputByState = map[state]map[int]state{
	stateWaitAnnotation: {
		IDENT: stateNamespaceDetected,
	},
	stateNamespaceDetected: {
		COLON: stateColonDetected,
	},
	stateColonDetected: {
		IDENT: stateNameDetected,
	},
	stateNameDetected: {
		LPAREN: stateWaitFieldOrParen,
		IDENT:  stateNamespaceDetected,
	},
	stateWaitFieldOrParen: {
		IDENT:  stateFieldNameDetected,
		RPAREN: stateWaitAnnotation,
	},
	stateFieldNameDetected: {
		EQUAL: stateFieldEqualDetected,
	},
	stateFieldEqualDetected: {
		LBRACE: stateWaitArrayValueOrBrace,
		// literals
		STRING: stateFieldValueDetected,
		INT:    stateFieldValueDetected,
		FLOAT:  stateFieldValueDetected,
		TRUE:   stateFieldValueDetected,
		FALSE:  stateFieldValueDetected,
	},
	stateWaitArrayValueOrBrace: {
		RBRACE: stateFieldValueDetected,
		// literals
		STRING: stateArrayValueDetected,
		INT:    stateArrayValueDetected,
		FLOAT:  stateArrayValueDetected,
		TRUE:   stateArrayValueDetected,
		FALSE:  stateArrayValueDetected,
	},
	stateArrayValueDetected: {
		COMMA:  stateWaitArrayValueOrBrace,
		RBRACE: stateFieldValueDetected,
	},
	stateFieldValueDetected: {
		COMMA:  stateWaitFieldOrParen,
		RPAREN: stateWaitAnnotation,
	},
}

func (s state) String() string {
	str, ok := stringByState[s]
	if !ok {
		return stringByState[stateUnknown]
	}
	return str
}

func (s state) Whitespace() uint64 {
	switch s {
	case stateNamespaceDetected, stateColonDetected:
		return 1
	}
	return 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '
}

func (s state) NextState(input int) state {
	nextState, ok := nextStateByInputByState[s][input]
	if !ok {
		return stateUnknown
	}
	return nextState
}
