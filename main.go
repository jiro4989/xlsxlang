package main

import (
	"fmt"
	"strconv"
)

type Calc struct {
	OpeStack   []string
	DigitQueue []int
}

func (c *Calc) PushOpe(ope string) {
	a := c.popDigit()
	b := c.popDigit()
	switch ope {
	case "+":
		fmt.Printf("%d+%d\n", b, a)
		c.pushDigit(b + a)
	case "-":
		fmt.Printf("%d-%d\n", b, a)
		c.pushDigit(b - a)
	case "*":
		fmt.Printf("%d*%d\n", b, a)
		c.pushDigit(b * a)
	case "/":
		fmt.Printf("%d/%d\n", b, a)
		c.pushDigit(b / a)
	}
}

func (c *Calc) PushDigit(digit string) {
	n, _ := strconv.Atoi(digit)
	c.pushDigit(n)
}

func (c *Calc) pushDigit(n int) {
	c.DigitQueue = append(c.DigitQueue, n)
}

func (c *Calc) popDigit() int {
	var n int
	n, c.DigitQueue = c.DigitQueue[len(c.DigitQueue)-1], c.DigitQueue[:len(c.DigitQueue)-1]
	return n
}

func (c *Calc) Compute() {
	fmt.Printf("result: %d\n", c.popDigit())
}

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
