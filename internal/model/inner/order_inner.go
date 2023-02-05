package inner

import (
	"database/sql/driver"
	"encoding/json"
)

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
	Price		float64 `json:"price"`
	Discount    float64 `json:"discount"`
}

func (c CommodityInOrder) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *CommodityInOrder) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}