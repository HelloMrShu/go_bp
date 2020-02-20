package errors

type errorString struct {
	s string
}

func New(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.s
}