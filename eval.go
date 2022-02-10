package main

import (
	"strconv"
)

type TokenKind int

type Eval struct {
	tokens       []Token
	depth        int
	bufferTokens []Token
}

type Token struct {
	Kind        TokenKind
	ValueBool   bool
	ValueInt    int64
	ValueStr    string
	ValueSymbol string
	ValueList   []Token
}

const (
	kindBool TokenKind = iota
	kindInt
	kindStr
	kindSymbol
	kindList
)

func (e *Eval) PushBool(s string) {
	b, _ := strconv.ParseBool(s)
	token := Token{
		Kind:      kindBool,
		ValueBool: b,
	}
	e.push(token)
}

func (e *Eval) PushInt(s string) {
	i, _ := strconv.ParseInt(s, 10, 64)
	token := Token{
		Kind:     kindInt,
		ValueInt: i,
	}
	e.push(token)
}

func (e *Eval) PushStr(s string) {
	token := Token{
		Kind:     kindStr,
		ValueStr: s,
	}
	e.push(token)
}

func (e *Eval) PushSymbol(s string) {
	token := Token{
		Kind:        kindSymbol,
		ValueSymbol: s,
	}
	e.push(token)
}

func (e *Eval) Begin() {
	e.depth++
	token := Token{
		Kind: kindList,
	}
	e.bufferTokens = append(e.bufferTokens, token)
}

func (e *Eval) End() {
	e.depth--
	token := e.bufferTokens[e.depth]
	e.bufferTokens = e.bufferTokens[:e.depth]
	e.push(token)
}

func (e *Eval) Evaluate() {
}

func (e *Eval) push(token Token) {
	d := e.depth
	if d < 1 {
		e.tokens = append(e.tokens, token)
		return
	}
	d--
	e.bufferTokens[d].ValueList = append(e.bufferTokens[d].ValueList, token)
}
