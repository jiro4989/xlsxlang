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
		c.pushDigit(b + a)
	case "-":
		c.pushDigit(b - a)
	case "*":
		c.pushDigit(b * a)
	case "/":
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
	fmt.Println(c.popDigit())
}
