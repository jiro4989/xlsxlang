package builtin

import "math"

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Mod(a, b int) int {
	return a % b
}

func Power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
