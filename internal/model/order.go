package model

import (
	"github.com/riicarus/loveshop/internal/model/inner"
)

type Order struct {
	Id          string                      `json:"id"`
	UserId      string                      `json:"userId"`
	AdminId     string                      `json:"adminId"`
	Time        int64                         `json:"time"`
	Commodities inner.CommodityInOrderSlice `gorm:"TYPE:json" json:"commodities"`
	Payment     float64                     `json:"payment"`
	Status      string                      `json:"status"`
	Type        string                      `json:"type"`
}

func (Order) TableName() string {
	return "order"
}

type OrderModel interface {
	Add(order *Order) error
	CancleOrder(id string) error
	PayOrder(id string) error
	FinishOrder(id string) error

	FindById(id string) (*Order, error)
}
