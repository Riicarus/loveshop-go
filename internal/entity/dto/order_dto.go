package dto

import "github.com/riicarus/loveshop/internal/model/inner"

type OrderAddParam struct {
	UserId      string                      `json:"userId"`
	AdminId     string                      `json:"adminId"`
	Commodities inner.CommodityInOrderSlice `json:"commodities" binding:"required"`
	Type        string                      `json:"type" binding:"required"`
}

type OrderDetailAdminView struct {
	Id          string                      `json:"id"`
	UserId      string                      `json:"userId"`
	Username    string                      `json:"username"`
	AdminId     string                      `json:"adminId"`
	Adminname   string                      `json:"adminname"`
	Time        string                      `json:"time"`
	Timestamp   int64                       `json:"timestamp"`
	Commodities inner.CommodityInOrderSlice `json:"commodities"`
	Payment     float64                     `json:"payment"`
	Status      string                      `json:"status"`
	Type        string                      `json:"type"`
}

type OrderDetailUserView struct {
	Id          string                      `json:"id"`
	UserId      string                      `json:"userId"`
	Time        string                      `json:"time"`
	Commodities inner.CommodityInOrderSlice `json:"commodities"`
	Payment     float64                     `json:"payment"`
}
