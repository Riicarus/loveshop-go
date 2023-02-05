package dto

import "github.com/riicarus/loveshop/internal/model/inner"

type OrderAddParam struct {
	UserId      string                      `json:"userId"`
	AdminId     string                      `json:"adminId"`
	Commodities inner.CommodityInOrderSlice `json:"commodities" binding:"required"`
	Type        string                      `json:"type" binding:"required"`
}