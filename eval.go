package main

import (
	"fmt"
)

type TokenKind int

type Eval struct {
	Token
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
	fmt.Println(s)
}

func (e *Eval) PushInt(s string) {
	fmt.Println(s)
}

func (e *Eval) PushStr(s string) {
	fmt.Println(s)
}

func (e *Eval) PushSymbol(s string) {
	fmt.Println(s)
}

func (e *Eval) PushList() {
	fmt.Println("push list")
}

func (e *Eval) Evaluate() {
	fmt.Println("eval")
}
