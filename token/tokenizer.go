package token

import (
	"strconv"
)

type Tokenizer struct {
	tokens       []Token
	depth        int
	bufferTokens []Token
}

func (e *Tokenizer) GetTokens() []Token {
	return e.tokens
}

func (e *Tokenizer) PushBool(s string) {
	b, _ := strconv.ParseBool(s)
	token := NewBoolToken(b)
	e.push(token)
}

func (e *Tokenizer) PushInt(s string) {
	i, _ := strconv.ParseInt(s, 10, 64)
	token := NewIntToken(i)
	e.push(token)
}

func (e *Tokenizer) PushStr(s string) {
	token := NewStrToken(s)
	e.push(token)
}

func (e *Tokenizer) PushNil() {
	token := NewNilToken()
	e.push(token)
}

func (e *Tokenizer) PushSymbol(s string) {
	token := NewSymbolToken(s)
	e.push(token)
}

func (e *Tokenizer) Begin() {
	e.depth++
	token := NewListToken()
	e.bufferTokens = append(e.bufferTokens, token)
}

func (e *Tokenizer) End() {
	e.depth--
	token := e.bufferTokens[e.depth]
	e.bufferTokens = e.bufferTokens[:e.depth]
	e.push(token)
}

func (e *Tokenizer) push(token Token) {
	d := e.depth
	if d < 1 {
		e.tokens = append(e.tokens, token)
		return
	}
	d--
	e.bufferTokens[d].ValueList = append(e.bufferTokens[d].ValueList, token)
}
