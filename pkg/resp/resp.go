package resp

import (
	"time"

	"github.com/riicarus/loveshop/pkg/e"
)

type Resp[T interface{}] struct {
	Timestamp int64      `json:"timestamp"`
	Msg       e.RespMsg  `json:"msg"`
	Code      e.RespCode `json:"code"`
	Data      T          `json:"data"`
}

func OK[T interface{}](data T) *Resp[T] {
	return &Resp[T]{
		Timestamp: time.Now().Unix(),
		Msg:       "OK",
		Code:      e.NONE_CODE,
		Data:      data,
	}
}

func Fail[T interface{}](msg e.RespMsg, code e.RespCode) *Resp[T] {
	return &Resp[T]{
		Timestamp: time.Now().Unix(),
		Msg:       msg,
		Code:      code,
	}
}
