package myerr

import (
	"bala/app/protocol/errcode"
	"bala/app/util"
	"fmt"

	"github.com/go-errors/errors"
)

type Error struct {
	code errcode.ErrorCode
	err  *errors.Error
}

func New(code errcode.ErrorCode, err error) error {
	if err == nil {
		return nil
	}
	if v, ok := err.(*Error); ok {
		return v
	}
	return &Error{code: code, err: errors.Wrap(err, 1)}
}

func NewFmt(code errcode.ErrorCode, format string, args ...interface{}) error {
	return &Error{code: code, err: errors.Wrap(fmt.Errorf(format, args...), 1)}
}

func (b *Error) Code() errcode.ErrorCode {
	return b.code
}

func (b *Error) Error() string {
	return b.err.Error()
}

func (b *Error) Stack() string {
	return fmt.Sprintf("%s,%s\n%s", b.code.String(), b.err.Error(), util.BytesToString(b.err.Stack()))
}
