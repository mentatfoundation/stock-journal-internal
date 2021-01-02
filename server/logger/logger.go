package logger

import (
	"fmt"
)

type Logger interface {
	Info(operator string, message string)
}

type localLogger struct{}

func New(env string) Logger {
	switch env {
	case "dev":
		return &localLogger{}
	}
	return &localLogger{}
}

func (l *localLogger) Info(operator string, message string) {
	fmt.Println(operator + "::" + message)
}
