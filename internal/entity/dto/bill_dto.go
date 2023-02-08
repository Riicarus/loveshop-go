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

type BillAnalyzeInfo struct {
	Day                    float64    `json:"day"`
	Week                   float64    `json:"week"`
	Month                  float64    `json:"month"`
	All                    float64    `json:"all"`
	Book                   float64    `json:"book"`
	CulturalCreativity     float64    `json:"culturalCreativity"`
	DailyNecessity         float64    `json:"dailyNecessity"`
	SportsGoods            float64    `json:"sportsGoods"`
	BoardGame              float64    `json:"boardGames"`
	BillCount              int        `json:"billCount"`
	CulturalCreativityInfo [][]string `json:"culturalCreativityInfo"`
}
