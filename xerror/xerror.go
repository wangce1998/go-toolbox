package xerror

type XError interface {
	Code() int
	Message() string
}

type xError struct {
	code    int
	message string
}

func (x xError) Code() int {
	return x.code
}

func (x xError) Message() string {
	return x.message
}

// NewXError 实例xerror
func NewXError(code int, message string) XError {
	return &xError{
		code:    code,
		message: message,
	}
}

// Wrap 将error包装成xerror
func Wrap(err error) XError {
	return NewXError(10000, err.Error())
}
