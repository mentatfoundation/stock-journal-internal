package logger

import (
	"fmt"
	"mentatfoundation/stock-journal/server/config"
)

type Logger interface {
	Info(operator string, message string)
}

type localLogger struct{}

type logger struct{}

type testLogger struct{}

func New(c config.ConfigurationSettings) Logger {
	switch c.Env {
	case "dev":
		return &localLogger{}
	case "test":
		return &testLogger{}
	}

	return &logger{}
}

func (l *localLogger) Info(operator string, message string) {
	fmt.Println(operator + "::" + message)
}

func (l *logger) Info(operator string, message string) {
	fmt.Println(operator + "::" + message)
}

func (l *testLogger) Info(operator string, message string) {}
