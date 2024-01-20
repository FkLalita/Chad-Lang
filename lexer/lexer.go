package lexer

import (
	"chad/token"
)

type Lexer struct {
	input        string // the source code that we are scanning.
	position     int    // current position in input (points to current char).
	readPosition int    //current reading position in input (after current char).
	ch           byte   //current char under examination
}

func New(input) *Lexer {
	return &Lexer{
		input: input,
	}
}
