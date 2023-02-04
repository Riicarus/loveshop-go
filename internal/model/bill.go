package model

type Bill struct {
	Id        string `json:"id"`
	Time      int    `json:"time"`
	OrderId   string `json:"orderId"`
	AdminId   string `json:"adminId"`
	OrderType string `json:"orderType"`
}

func (Bill) TableName() string {
	return "bill"
}