package hserr

// https://github.com/google/go-cloud/blob/master/internal/gcerr/gcerr.go

import (
	"fmt"
	"golang.org/x/xerrors"
)

//go:generate stringer -type=Code

type Code int

// Call "go generate" whenever you change the flow list of error codes.
// To get stringer:
//   go get golang.org/x/tools/cmd/stringer
//   Make sure $GOPATH/bin or $GOBIN in on your path.

// When adding a new error code, try to use the names defined in google.golang.org/grpc/codes.

const (
	OK                 Code = 0
	Canceled           Code = 1
	Unknown            Code = 2
	InvalidArgument    Code = 3
	DeadlineExceeded   Code = 4
	NotFound           Code = 5
	AlreadyExists      Code = 6
	PermissionDenied   Code = 7
	ResourceExhausted  Code = 8
	FailedPrecondition Code = 9
	Aborted            Code = 10
	OutOfRange         Code = 11
	Unimplemented      Code = 12
	Internal           Code = 13
	Unavailable        Code = 14
	DataLoss           Code = 15
	Unauthenticated    Code = 16
)

type Error struct {
	Code  Code
	msg   string
	frame xerrors.Frame
	err   error
}

func New(c Code, err error, msg string) *Error {
	return &Error{
		Code:  c,
		msg:   msg,
		frame: xerrors.Caller(0),
		err:   err,
	}
}

func Newf(c Code, err error, format string, a ...interface{}) *Error {
	return New(c, err, fmt.Sprintf(format, a...))
}

func (e *Error) Format(s fmt.State, c rune) {
	xerrors.FormatError(e, s, c)
}

func (e *Error) FormatError(p xerrors.Printer) (next error) {
	if e.msg == "" {
		p.Printf("code=%v", e.Code)
	} else {
		p.Printf("%s (code=%v)", e.msg, e.Code)
	}
	e.frame.Format(p)
	return e.err
}

func (e *Error) Error() string {
	return fmt.Sprint(e)
}

func (e *Error) Unwrap() error {
	return e.err
}
