package token

import (
	"strconv"
)

type Tokenizer struct {
	tokens       []Token
	depth        int
	bufferTokens []Token
}

type Token struct {
	Kind        TokenKind
	ValueBool   bool
	ValueInt    int64
	ValueStr    string
	ValueNil    bool
	ValueSymbol string
	ValueList   []Token
}

func (e *Tokenizer) GetTokens() []Token {
	return e.tokens
}

func (e *Tokenizer) PushBool(s string) {
	b, _ := strconv.ParseBool(s)
	token := Token{
		Kind:      KindBool,
		ValueBool: b,
	}
	e.push(token)
}

func (e *Tokenizer) PushInt(s string) {
	i, _ := strconv.ParseInt(s, 10, 64)
	token := Token{
		Kind:     KindInt,
		ValueInt: i,
	}
	e.push(token)
}

func (e *Tokenizer) PushStr(s string) {
	token := Token{
		Kind:     KindStr,
		ValueStr: s,
	}
	e.push(token)
}

func (e *Tokenizer) PushNil() {
	token := Token{
		Kind:     KindNil,
		ValueNil: true,
	}
	e.push(token)
}

func (e *Tokenizer) PushSymbol(s string) {
	token := Token{
		Kind:        KindSymbol,
		ValueSymbol: s,
	}
	e.push(token)
}

func (e *Tokenizer) Begin() {
	e.depth++
	token := Token{
		Kind: KindList,
	}
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
