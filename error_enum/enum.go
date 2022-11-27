package error_enum

const (
	_COMMON_ENUM errCode = iota * 100
	_UN_AUTH
)

const (
	UnknownError = _COMMON_ENUM + iota + 1
	MissingParameter
)

// 通用异常
func setCommonEnum() {
	setErrEnum(UnknownError, "UnknownError")
	setErrEnum(MissingParameter, "MissingParameter")
}

func init() {
	setCommonEnum()
}
