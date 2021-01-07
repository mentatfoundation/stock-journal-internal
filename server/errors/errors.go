package errors

type Error struct {
	Operation string
	Message   string
	Code      int
}

func Create(op string, msg string, code int) *Error {
	return &Error{Operation: op, Message: msg, Code: code}
}
