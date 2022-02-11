package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	Program string
	Args    []string
}

func ParseArgs() (*CmdArgs, error) {
	opts := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.StringVar(&opts.Program, "e", "", "one line of script")
	flag.Parse()
	opts.Args = flag.Args()

	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &opts, nil
}

func flagHelpMessage() {
	cmd := os.Args[0]
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS] [files...]", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s sample.xlsx", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")

	flag.PrintDefaults()
}

func (c *CmdArgs) Validate() error {
	if len(c.Args) < 1 && c.Program == "" {
		return fmt.Errorf("Must need args or '-e' option")
	}

	return nil
}
