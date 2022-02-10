package main

import (
	"fmt"
	"strconv"
)

type TokenKind int

type Eval struct {
	Token
	BufferDepth int
	BufferToken []Token
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
	e.BufferToken = append(e.BufferToken, token)
	e.Token = token
}

func (e *Eval) PushInt(s string) {
	i, _ := strconv.ParseInt(s, 10, 64)
	token := Token{
		Kind:     kindInt,
		ValueInt: i,
	}
	e.BufferToken = append(e.BufferToken, token)
	e.Token = token
}

func (e *Eval) PushStr(s string) {
	token := Token{
		Kind:     kindStr,
		ValueStr: s,
	}
	e.BufferToken = append(e.BufferToken, token)
	e.Token = token
}

func (e *Eval) PushSymbol(s string) {
	token := Token{
		Kind:        kindSymbol,
		ValueSymbol: s,
	}
	e.BufferToken = append(e.BufferToken, token)
	e.Token = token
}

func (e *Eval) PushList() {
	token := Token{
		Kind:      kindList,
		ValueList: e.BufferToken,
	}
	e.Token = token
}

func (e *Eval) Evaluate() {
	fmt.Println("eval")
}
