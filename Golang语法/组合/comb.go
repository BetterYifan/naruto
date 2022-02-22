package 组合

import "log"

type MyLogger struct {
	cmd string
	*log.Logger
}

type Logger struct {
	Level int
}

func NewMyLogger(cmd string) *MyLogger {
	return &MyLogger{
		cmd: cmd,
	}
}

func (l *MyLogger) FuncA() {
	l.Print()
}
