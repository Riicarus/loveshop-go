package e

type NotExistErr struct {
	Code RespCode
	Msg  RespMsg
}

var NOT_EXIST_ERR *NotExistErr = &NotExistErr{
	Code: NOT_EXIST_CODE,
	Msg:  NOT_EXIST_MSG,
}

func (e *NotExistErr) Error() string {
	return string(e.Msg)
}
