package e

type StockErr struct {
	Code RespCode
	Msg  RespMsg
}

var STOCK_ERR *StockErr = &StockErr{
	Code: STOCK_NOT_ENOUGH_CODE,
	Msg:  STOCK_NOT_ENOUGH_MSG,
}

func (e *StockErr) Error() string {
	return string(e.Msg)
}
