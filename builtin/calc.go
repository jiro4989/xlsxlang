package builtin

import (
	"math"

	"github.com/jiro4989/xlsxlang/token"
)

func Plus(a, b int64) token.Token {
	return token.NewIntToken(a + b)
}

func Minus(a, b int64) token.Token {
	return token.NewIntToken(a - b)
}

func Mul(a, b int64) token.Token {
	return token.NewIntToken(a * b)
}

func Div(a, b int64) token.Token {
	return token.NewIntToken(a / b)
}

func Mod(a, b int64) token.Token {
	return token.NewIntToken(a % b)
}

func Power(a, b int64) token.Token {
	return token.NewIntToken(int64(math.Pow(float64(a), float64(b))))
}
