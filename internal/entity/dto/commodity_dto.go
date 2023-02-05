package dto

import (
	"github.com/riicarus/loveshop/pkg/util"
)

type CommodityAddParam struct {
	Type      string       `json:"type" binding:"required"`
	Numbering string       `json:"numbering" binding:"required"`
	Name      string       `json:"name" binding:"required"`
	Amount    int          `json:"amount" binding:"required"`
	Price     float64      `json:"price" binding:"required"`
	Extension util.JSONMap `json:"extension"`
}

type CommodityUpdateParam struct {
	Id        string       `json:"id" binding:"required"`
	Type      string       `json:"type"`
	Numbering string       `json:"numbering"`
	Name      string       `json:"name"`
	Price     float64      `json:"price"`
	Extension util.JSONMap `json:"extension"`
}

type CommoditySimpleView struct {
	Id        string       `json:"id"`
	Type      string       `json:"type"`
	Numbering string       `json:"numbering"`
	Name      string       `json:"name"`
	Amount    int          `json:"amount"`
	Price     float64      `json:"price"`
	Extension util.JSONMap `json:"extension"`
}

type CommodityDetailView struct {
	Id        string       `json:"id"`
	Type      string       `json:"type"`
	Numbering string       `json:"numbering"`
	Name      string       `json:"name"`
	Amount    int          `json:"amount"`
	Price     float64      `json:"price"`
	Extension util.JSONMap `json:"extension"`
	Deleted   bool         `json:"deleted"`
}
