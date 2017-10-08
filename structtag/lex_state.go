package main

import (
	"text/scanner"
)

type state int

const (
	stateDefault state = iota
	stateKeyDetected
	stateColonDetected
	stateValueDetected
	stateCommaDetected
)

func (s state) whitespace() uint64 {
	switch s {
	case stateKeyDetected:
		return 1
	case stateColonDetected:
		return 1
	}
	return scanner.GoWhitespace
}
