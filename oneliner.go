//go:build !debug

package main

import "fmt"

var RunOneliner = runOneliner

func runOneliner(program string) error {
	fmt.Println("Use excel file.")
	return nil
}
