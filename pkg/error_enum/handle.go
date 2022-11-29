package error_enum

import (
	"fmt"
)

var (
	codeEnumMap = map[errCode]errMsg{}
)

func setErrEnum(code errCode, msg errMsg) {
	if _, has := codeEnumMap[code]; has {
		panic(fmt.Sprintf("duplicate error code:%d", code))
	}
	codeEnumMap[code] = msg
}

func UndefinedError(err any) *errAble {
	e := &errAble{
		code: -1,
		msg:  "UndefinedError",
	}
	if err != nil {
		e.err = fmt.Errorf("%v", err)
	}
	return e
}

func ErrPanic(err error, code errCode, args ...interface{}) {
	if err == nil {
		return
	}
	var e *errAble
	msg, has := codeEnumMap[code]
	if !has {
		e = UndefinedError(err)
	} else {
		e = &errAble{
			code: code,
			msg:  msg,
			err:  err,
		}
	}
	if len(args) > 0 {
		e.msg = errMsg(fmt.Sprintf(string(e.msg), args...))
	}
	panic(e)
}

func FalsePanic(ok bool, code errCode, args ...interface{}) {
	if ok {
		return
	}
	handlePanic(code, args...)
}

func handlePanic(code errCode, args ...interface{}) {
	var e *errAble
	msg, has := codeEnumMap[code]
	if !has {
		e = UndefinedError(nil)
	} else {
		e = &errAble{
			code: code,
			msg:  msg,
		}
	}
	if len(args) > 0 {
		e.msg = errMsg(fmt.Sprintf(string(e.msg), args...))
	}
	panic(e)
}
