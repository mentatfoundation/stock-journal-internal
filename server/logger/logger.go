package logger

import (
	"fmt"
	"mentatfoundation/stock-journal/server/config"
)

// make interface for other loggers

type Logger struct{}

func New(c config.ConfigurationSettings) *Logger {
	fmt.Println(c.Port)
	return &Logger{}
}

func (l *Logger) Info(message string) {
	fmt.Println(message)
}
