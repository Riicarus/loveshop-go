package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Order struct {
	Id          string                `json:"id"`
	UserId      string                `json:"userId"`
	AdminId     string                `json:"adminId"`
	Time        int                   `json:"time"`
	Commodities CommodityInOrderSlice `gorm:"TYPE:json" json:"commodities"`
	Payment     float32               `json:"payment"`
	Status      string                `json:"status"`
	Type        string                `json:"type"`
}

func (Order) TableName() string {
	return "order"
}

type CommodityInOrderSlice []CommodityInOrder

func (s CommodityInOrderSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

func (s *CommodityInOrderSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}

type CommodityInOrder struct {
	CommodityId string  `json:"commodityId"`
	Amount      int     `json:"amount"`
	Discount    float32 `json:"discount"`
}

func (c CommodityInOrder) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *CommodityInOrder) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
