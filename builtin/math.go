package builtin

import (
	"math"

	"github.com/jiro4989/xlsxlang/token"
)

var (
	BuiltinMathFunctions map[string]func(a, b token.Token) token.Token = map[string]func(a, b token.Token) token.Token{
		"+":  Plus,
		"-":  Minus,
		"*":  Mul,
		"/":  Div,
		"%":  Mod,
		"**": Power,
	}
)

func Plus(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(a.ValueInt + b.ValueInt)
}

func Minus(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(a.ValueInt - b.ValueInt)
}

func Mul(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(a.ValueInt * b.ValueInt)
}

func Div(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(a.ValueInt / b.ValueInt)
}

func Mod(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(a.ValueInt % b.ValueInt)
}

func Power(a, b token.Token) token.Token {
	validateInt2(a, b)
	return token.NewIntToken(int64(math.Pow(float64(a.ValueInt), float64(b.ValueInt))))
}

func validateInt(t token.Token) {
	if t.Kind != token.KindInt {
		panic("token must be int: kind = " + t.Kind.String())
	}
}

func validateInt2(t, t2 token.Token) {
	validateInt(t)
	validateInt(t2)
}
