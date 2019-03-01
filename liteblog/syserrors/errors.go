package syserrors

type Error interface {
	Error() string
	Code() int
	ReasonError() error
}

func NewError(msg string, err2 error) UnknowError {
	err := UnknowError{}
	err.msg = msg
	err.reason = err2
	return err
}
