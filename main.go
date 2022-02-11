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
	exitStatusReadXlsxErr
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

	if opts.Program != "" {
		if err := RunOneliner(opts.Program); err != nil {
			log.Err(err)
			return exitStatusParseErr
		}
		return exitStatusOK
	}

	for _, filePath := range opts.Args {
		program, err := ReadXlsx(filePath)
		if err != nil {
			log.Err(err)
			return exitStatusReadXlsxErr
		}
		if err := eval(program); err != nil {
			log.Err(err)
			return exitStatusParseErr
		}
	}

	return exitStatusOK
}

func eval(program string) error {
	parser, err := parse(program)
	if err != nil {
		return err
	}
	result := Evaluate(parser.GetTokens())
	fmt.Println(result.StringResult())

	return nil
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
