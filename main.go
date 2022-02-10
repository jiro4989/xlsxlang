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

	parser, err := parse(opts.Eval)
	if err != nil {
		Err(err)
		return exitStatusParseErr
	}
	parser.Evaluate()

	return exitStatusOK
}

func parse(s string) (*Parser, error) {
	parser := &Parser{Buffer: s}
	parser.Init()
	if err := parser.Parse(); err != nil {
		return nil, err
	}

	parser.Execute()
	return parser, nil
}
