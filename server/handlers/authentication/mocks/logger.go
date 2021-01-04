package mocks

type LoggerMock struct{}

var InfoMock func(operator string, message string)

func (as LoggerMock) Info(operator string, message string) {
	InfoMock(operator, message)
}
