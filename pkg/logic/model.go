package logic

import "github.com/riicarus/loveshop/pkg/connection"

type IDBModel[T interface{}] interface {
	Conn(txctx *connection.TxContext) T
}

type DBModel struct {
	Txctx *connection.TxContext
}