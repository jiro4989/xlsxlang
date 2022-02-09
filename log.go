package main

import (
	"log"
	"os"
)

const (
	logFlag = log.LstdFlags | log.Lshortfile | log.Lmsgprefix | log.Lmicroseconds
)

var (
	stdoutLogger = log.New(os.Stdout, appName, logFlag)
	stderrLogger = log.New(os.Stderr, appName, logFlag)
)

func Info(s string) {
	stdoutLogger.Output(1, " [INFO] "+s)
}

func Err(err error) {
	stderrLogger.Output(1, " [ERR ] "+err.Error())
}
