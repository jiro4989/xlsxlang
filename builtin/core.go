package builtin

import (
	"fmt"

	"github.com/jiro4989/xlsxlang/token"
)

type PrintFunction func(a token.Token) token.Token

var (
	PrintFunctions map[string]PrintFunction = map[string]PrintFunction{
		"println": Println,
	}
)

func Println(a token.Token) token.Token {
	s := a.StringResult()
	fmt.Println(s)
	return token.NewNilToken()
}
