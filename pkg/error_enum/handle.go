package error_enum

import (
	"fmt"
)

var (
	codeEnumMap = map[errCode]*errAble{}
)

func setErrEnum(code errCode, msg errMsg) {
	if _, has := codeEnumMap[code]; has {
		panic(fmt.Sprintf("duplicate error code:%d", code))
	}
	codeEnumMap[code] = &errAble{
		code: code,
		msg:  msg,
	}
}

func UndefinedError(err any) *errAble {
	e := &errAble{
		code: -1,
		msg:  "UndefinedError",
		err:  fmt.Errorf("%v", err),
	}
	return e
}

func ErrPanic(err error, code errCode) {
	if err == nil {
		return
	}
	e, has := codeEnumMap[code]
	if !has {
		e = UndefinedError(err)
	} else {
		e.err = err
	}
	panic(e)
}
