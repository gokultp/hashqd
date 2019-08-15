package commands

import "fmt"

var (
	errInvalidCommand = newErr(101, "invalid cmd")
	errUnknown        = newErr(102, "something went wrong")
	errBadRequest     = newErr(103, "bad request")
	errPut            = newErr(104, "could not put")
	errReserve        = newErr(104, "could not reserve")
	errMutate         = newErr(105, "could not mutate")
)

type Error struct {
	Message string
	Code    int
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (e *Error) SetMessage(msg string) *Error {
	fmt.Println(e.Message, msg)
	e.Message += " " + msg
	return e
}

func (e *Error) Bytes() []byte {
	return []byte(fmt.Sprintf("%d : %s.\n", e.Code, e.Message))
}

func newErr(code int, msg string) func(string) *Error {
	return func(spmsg string) *Error {
		return NewError(code, msg).SetMessage(spmsg)
	}
}
