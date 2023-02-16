package model

import (
	"github.com/riicarus/loveshop/internal/model/inner"
	"github.com/riicarus/loveshop/pkg/logic"
)

type Order struct {
	Id          string                      `json:"id"`
	UserId      string                      `json:"userId"`
	AdminId     string                      `json:"adminId"`
	Time        int64                       `json:"time"`
	Commodities inner.CommodityInOrderSlice `gorm:"TYPE:json" json:"commodities"`
	Payment     float64                     `json:"payment"`
	Status      string                      `json:"status"`
	Type        string                      `json:"type"`
}

type OrderDetail struct {
	Id          string                      `json:"id"`
	UserId      string                      `json:"userId"`
	Username    string                      `json:"username"`
	AdminId     string                      `json:"adminId"`
	Adminname   string                      `json:"adminname"`
	Time        int64                       `json:"time"`
	Commodities inner.CommodityInOrderSlice `gorm:"TYPE:json" json:"commodities"`
	Payment     float64                     `json:"payment"`
	Status      string                      `json:"status"`
	Type        string                      `json:"type"`
}

func (Order) TableName() string {
	return "order"
}

type OrderModel interface {
	logic.IDBModel[OrderModel]

	Add(order *Order) error
	CancelOrder(id string) error
	PayOrder(id string) error
	FinishOrder(id string) error

	FindById(id string) (*Order, error)
	FindDetailById(id string) (*OrderDetail, error)

	FindPageOrderByTime(desc bool, num, size int) ([]*OrderDetail, error)
	FindPageByStatusOrderByTime(status string, desc bool, num, size int) ([]*OrderDetail, error)
	FindUserViewPageByUidOrderByTime(uid string, desc bool, num, size int) ([]*Order, error)
	FindUserViewPageByUidAndStatusOrderByTime(uid, status string, desc bool, num, size int) ([]*Order, error)
}
