package model

import (
	"github.com/riicarus/loveshop/pkg/logic"
)

type Bill struct {
	Id        string `json:"id"`
	Time      int64  `json:"time"`
	OrderId   string `json:"orderId"`
	AdminId   string `json:"adminId"`
	OrderType string `json:"orderType"`
}

func (Bill) TableName() string {
	return "bill"
}

type BillModel interface {
	logic.IDBModel[BillModel]

	Add(bill *Bill) error

	FindById(id string) (*Bill, error)
	FindAll() ([]*Bill, error)
	FindPageOrderByTime(desc bool, num, size int) ([]*Bill, error)
	FindPageByOrderTypeOrderByTime(orderType string, desc bool, num, size int) ([]*Bill, error)
}
