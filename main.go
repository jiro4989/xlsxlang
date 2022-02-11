package main

import (
	"fmt"
	"os"

	"github.com/jiro4989/xlsxlang/logger"
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
	log := logger.New(appName, os.Stdout, os.Stderr)
	opts, err := ParseArgs()
	if err != nil {
		log.Err(err)
		return exitStatusCommandLineOptionParseErr
	}

	parser, err := parse(opts.Eval)
	if err != nil {
		log.Err(err)
		return exitStatusParseErr
	}
	result := Evaluate(parser.GetTokens())
	fmt.Println(result.StringResult())

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
