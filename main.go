package main

import (
	"fmt"
)

func main() {
	parser := &Parser{Buffer: "1+2*3-(4+5)"}
	parser.Init()
	err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	} else {
		parser.Execute()
		parser.Compute()
	}
}
