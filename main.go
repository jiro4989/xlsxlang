package main

import (
	"os"
)

type ExitStatus int

const (
	appName = "xlsxlang"

	exitStatusOK ExitStatus = iota
	exitStatusCommandLineOptionParseErr
	exitStatusParseErr
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(args []string) ExitStatus {
	opts, err := ParseArgs()
	if err != nil {
		Err(err)
		return exitStatusCommandLineOptionParseErr
	}

	parser := &Parser{Buffer: opts.Eval}
	parser.Init()
	if err := parser.Parse(); err != nil {
		Err(err)
		return exitStatusParseErr
	}

	parser.Execute()
	parser.Compute()
	return exitStatusOK
}
