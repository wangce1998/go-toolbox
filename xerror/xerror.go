package xerror

type XError interface {
	Code() int
	Message() string
}

type defaultXError struct {
	code    int
	message string
}

func (x defaultXError) Code() int {
	return x.code
}

func (x defaultXError) Message() string {
	return x.message
}

// New 实例xerror
func New(code int, message string) XError {
	return &defaultXError{
		code:    code,
		message: message,
	}
}

// Wrap 将error包装成xerror
func Wrap(err error) XError {
	if err == nil {
		return nil
	}

	return New(10000, err.Error())
}
