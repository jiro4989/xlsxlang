package logger

import (
	"io"
	"log"
)

const (
	logFlag = log.LstdFlags | log.Lshortfile | log.Lmsgprefix | log.Lmicroseconds
)

type Logger struct {
	stdout *log.Logger
	stderr *log.Logger
}

func New(appName string, stdout, stderr io.Writer) Logger {
	l := Logger{
		stdout: log.New(stdout, appName, logFlag),
		stderr: log.New(stderr, appName, logFlag),
	}
	return l
}

func (l *Logger) Info(s string) {
	l.stdout.Output(1, " [INFO] "+s)
}

func (l *Logger) Err(err error) {
	l.stderr.Output(1, " [ERR ] "+err.Error())
}
