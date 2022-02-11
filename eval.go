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
			// if は引数が3つ
			if t.ValueSymbol == "if" {
				var a, b, c token.Token
				a, tokens = dequeue(tokens)
				b, tokens = dequeue(tokens)
				c, tokens = dequeue(tokens)

				a = Evaluate([]token.Token{a})
				var arg token.Token
				// trueのときはbだけ評価、そうでなければcだけ評価
				if a.IsTrue() {
					arg = b
				} else {
					arg = c
				}
				return Evaluate([]token.Token{arg})
			}
			// mathの関数はいずれも引数が2つだけ
			if f, ok := isBuiltinMathFunction(t); ok {
				var a, b token.Token
				a, tokens = dequeue(tokens)
				b, tokens = dequeue(tokens)
				a = Evaluate([]token.Token{a})
				b = Evaluate([]token.Token{b})
				return f(a, b)
			}
			// print関数はいずれも引数が1つだけ
			if f, ok := isBuiltinPrintFunction(t); ok {
				var a token.Token
				a, tokens = dequeue(tokens)
				a = Evaluate([]token.Token{a})
				return f(a)
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

func isBuiltinMathFunction(t token.Token) (builtin.MathFunction, bool) {
	sym := t.ValueSymbol
	f, ok := builtin.MathFunctions[sym]
	if !ok {
		return nil, false
	}

	return f, ok
}

func isBuiltinPrintFunction(t token.Token) (builtin.PrintFunction, bool) {
	sym := t.ValueSymbol
	f, ok := builtin.PrintFunctions[sym]
	if !ok {
		return nil, false
	}

	return f, ok
}
