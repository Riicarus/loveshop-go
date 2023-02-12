package dto

import (
	"strings"

	"github.com/riicarus/loveshop/internal/consts"
	"github.com/riicarus/loveshop/pkg/e"
)

type BillAddParam struct {
	Time      int64  `json:"time"`
	OrderId   string `json:"orderId"`
	AdminId   string `json:"adminId"`
	OrderType string `json:"orderType"`
}

func (b *BillAddParam) Check() error {
	if b.Time <= 0 {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(b.OrderId) == "" {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(b.AdminId) == "" {
		return e.VALIDATE_ERR
	}

	switch b.OrderType {
	case consts.OFFLINE:
	case consts.ONLINE:
	default:
		return e.VALIDATE_ERR
	}

	return nil
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
