package main

import (
	"github.com/jiro4989/xlsxlang/builtin"
	"github.com/jiro4989/xlsxlang/token"
)

func Evaluate(tokens []token.Token) {
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
