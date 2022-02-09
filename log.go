package main

import (
	"log"
	"os"
)

var (
	stdoutLogger = log.New(os.Stdout, appName, log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	stderrLogger = log.New(os.Stderr, appName, log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
)

func Info(s string) {
	stdoutLogger.Output(1, "[INFO] "+s)
}

func Err(err error) {
	stderrLogger.Output(1, "[ERR ] "+err.Error())
}
