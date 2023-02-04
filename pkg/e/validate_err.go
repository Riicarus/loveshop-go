package e

type ValidateErr struct {
	Code RespCode
	Msg  RespMsg
}

var VALIDATE_ERR *ValidateErr = &ValidateErr{
	Code: VALIDATE_FAILED_CODE,
	Msg:  VALIDATE_FAILED_MSG,
}

func (e *ValidateErr) Error() string {
	return string(e.Msg)
}
