package logger

import (
	"fmt"
)

// make interface for other loggers
type Logger interface {
	Info(operator string, message string)
}

type logger struct{}

func New(env string) *logger {
	return &logger{}
}

func (l *logger) Info(operator string, message string) {
	fmt.Println(operator + "::" + message)
}
