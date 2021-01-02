package logger

import (
	"fmt"
)

// make interface for other loggers
type Logger interface {
	Info(message string)
}

type logger struct{}

func New(env string) *logger {
	fmt.Println(env)
	return &logger{}
}

func (l *logger) Info(message string) {
	fmt.Println(message)
}
