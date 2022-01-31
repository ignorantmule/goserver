package errno

import "fmt"

// Errno defines a new error type used by goserver.
type Errno struct {
	Code    int
	Message string
}

// Error implement the `Error` method in error interface.
func (err Errno) Error() string {
	return err.Message
}

// Err represents an error.
type Err struct {
	Code    int
	Message string
	Err     error
}

// New create a new `Err` error.
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// Add add message to `Err` error.
func (err *Err) Add(message string) error {
	err.Message += " " + message

	return err
}

// Addf add a formated message to `Err` error.
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)

	return err
}

// Error return error message in string format.
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// IsErrUserNotFound return true if the `err` is a `ErrUserNotFound` type error.
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)

	return code == ErrUserNotFound.Code
}

// DecodeErr decode an err message.
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	//nolint: errorlint
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
