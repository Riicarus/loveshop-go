package dto

import (
	"strings"

	"github.com/riicarus/loveshop/internal/consts"
	"github.com/riicarus/loveshop/internal/model/inner"
	"github.com/riicarus/loveshop/pkg/e"
)

type OrderAddParam struct {
	UserId      string                      `json:"userId"`
	AdminId     string                      `json:"adminId"`
	Commodities inner.CommodityInOrderSlice `json:"commodities" binding:"required"`
	Type        string                      `json:"type" binding:"required"`
}

func (o *OrderAddParam) Check() error {
	if strings.TrimSpace(o.UserId) == "" && strings.TrimSpace(o.AdminId) == "" {
		return e.VALIDATE_ERR
	}

	if len(o.Commodities) == 0 {
		return e.VALIDATE_ERR
	}

	switch o.Type {
	case consts.OFFLINE: 
	case consts.ONLINE:
	default:
		return e.VALIDATE_ERR
	}

	return nil
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
