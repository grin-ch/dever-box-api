package error_enum

const (
	_COMMON_ENUM errCode = iota * 100
	_AUTH
)

const (
	UnknownError = _COMMON_ENUM + iota + 1
	ServerBusy
	ExecSQLError
	MissingParameter
)

const (
	AuthError = _AUTH + iota + 1
)

// 通用异常
func setCommonEnum() {
	setErrEnum(UnknownError, "UnknownError")
	setErrEnum(ServerBusy, "ServerBusy")
	setErrEnum(ExecSQLError, "ExecSQLError")
	setErrEnum(MissingParameter, "MissingParameter")
}

func setAuthEnum() {
	setErrEnum(AuthError, "AuthError")
}

func init() {
	setCommonEnum()
	setAuthEnum()
}
