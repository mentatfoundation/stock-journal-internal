package logger

type Mock struct{}

var InfoMock func(operator string, message string)

func (as Mock) Info(operator string, message string) {
	InfoMock(operator, message)
}
