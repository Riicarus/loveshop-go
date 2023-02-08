package dto

type BillAddParam struct {
	Time      int64  `json:"time"`
	OrderId   string `json:"orderId"`
	AdminId   string `json:"adminId"`
	OrderType string `json:"orderType"`
}

type BillDetailAdminView struct {
	Id             string                `json:"id"`
	BillCreateTime string                `json:"billCreateTime"`
	OrderView      *OrderDetailAdminView `json:"orderView"`
	AdminId        string                `json:"adminId"`
	Adminname      string                `json:"adminname"`
}
