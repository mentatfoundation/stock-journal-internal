package logger

import "fmt"

type Logger struct {}


func (l *Logger) Info() {
	fmt.Println("logging")
}