package main

import (
	"github.com/jiro4989/xlsxlang/builtin"
	"github.com/jiro4989/xlsxlang/token"
)

type Eval struct {
	tokens []token.Token
}

func Evaluate(tokens []token.Token) token.Token {
	for 0 < len(tokens) {
		var t token.Token
		t, tokens = dequeue(tokens)

		switch t.Kind {
		case token.KindBool:
			return t
		case token.KindInt:
			return t
		case token.KindStr:
			return t
		case token.KindNil:
			return t
		case token.KindSymbol:
			// mathの関数はいずれも引数が2つだけ
			f, ok := isBuiltinMathFunction(t)
			if ok {
				var a, b token.Token
				a, tokens = dequeue(tokens)
				b, tokens = dequeue(tokens)
				a = Evaluate([]token.Token{a})
				b = Evaluate([]token.Token{b})
				return f(a, b)
			}
		case token.KindList:
			return Evaluate(t.ValueList)
		}
	}
	return token.NewNilToken()
}

func dequeue(tokens []token.Token) (token.Token, []token.Token) {
	t := tokens[0]
	tokens = tokens[1:]
	return t, tokens
}

func isBuiltinMathFunction(t token.Token) (func(a, b token.Token) token.Token, bool) {
	if t.Kind != token.KindSymbol {
		return nil, false
	}

	sym := t.ValueSymbol
	f, ok := builtin.BuiltinMathFunctions[sym]
	if !ok {
		return nil, false
	}

	return f, ok
}
