package response

type Error struct {
	Code int
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewError(code int, err error) error {
	return &Error{code, err}
}
