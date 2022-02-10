package builtin

import "math"

func Plus(a, b int64) int64 {
	return a + b
}

func Minus(a, b int64) int64 {
	return a - b
}

func Mul(a, b int64) int64 {
	return a * b
}

func Div(a, b int64) int64 {
	return a / b
}

func Mod(a, b int64) int64 {
	return a % b
}

func Power(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
