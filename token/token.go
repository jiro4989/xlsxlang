package token

import (
	"fmt"
	"strings"
)

type Token struct {
	Kind        TokenKind
	ValueBool   bool
	ValueInt    int64
	ValueStr    string
	ValueNil    bool
	ValueSymbol string
	ValueList   []Token
}

func NewBoolToken(b bool) Token {
	return Token{
		Kind:      KindBool,
		ValueBool: b,
	}
}

func NewIntToken(i int64) Token {
	return Token{
		Kind:     KindInt,
		ValueInt: i,
	}
}

func NewStrToken(s string) Token {
	return Token{
		Kind:     KindStr,
		ValueStr: s,
	}
}

func NewNilToken() Token {
	return Token{
		Kind:     KindNil,
		ValueNil: true,
	}
}

func NewSymbolToken(s string) Token {
	return Token{
		Kind:        KindSymbol,
		ValueSymbol: s,
	}
}

func NewListToken() Token {
	return Token{
		Kind: KindList,
	}
}

func (t *Token) IsTrue() bool {
	return t.Kind == KindBool && t.ValueBool
}

func (t *Token) StringResult() string {
	switch t.Kind {
	case KindBool:
		return fmt.Sprintf("%v", t.ValueBool)
	case KindInt:
		return fmt.Sprintf("%d", t.ValueInt)
	case KindStr:
		return t.ValueStr
	case KindNil:
		return "nil"
	case KindSymbol:
		return fmt.Sprintf("Symbol:%v", t.ValueSymbol)
	case KindList:
		var arr []string
		for _, v := range t.ValueList {
			s := v.StringResult()
			arr = append(arr, s)
		}
		return strings.Join(arr, ",")
	}
	return "Undefined"
}
