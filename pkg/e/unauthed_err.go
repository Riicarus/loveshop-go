package e

type UnauthedErr struct {
	Code RespCode
	Msg  RespMsg
}

var UNAUTHED_ERR *UnauthedErr = &UnauthedErr{
	Code: UNATHENRIZED_CODE,
	Msg:  UNATHENRIZED_MSG,
}

func (e *UnauthedErr) Error() string {
	return string(e.Msg)
}
