package error_enum

type errCode int
type errMsg string

type ErrAble interface {
	Code() errCode
	Msg() errMsg
	Err() string
}

type errAble struct {
	code errCode
	msg  errMsg
	err  error
}

func (err *errAble) Code() errCode {
	return err.code
}

func (err *errAble) Msg() errMsg {
	return err.msg
}

func (err *errAble) Err() string {
	return err.err.Error()
}
