package model

import (
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/pkg/util"
)

type Commodity struct {
	Id        string       `json:"id"`
	Type      string       `json:"type"`
	Numbering string       `json:"numbering"`
	Name      string       `json:"name"`
	Amount    int          `json:"amount"`
	Price     float64      `json:"price"`
	Extension util.JSONMap `gorm:"TYPE:json" json:"extension"`
	Deleted   bool         `json:"deleted"`
}

func (Commodity) TableName() string {
	return "commodity"
}

type CommodityModel interface {
	Add(param *dto.CommodityAddParam) error

	Update(param *dto.CommodityUpdateParam) error
	UpdateAmount(id string, number int) error

	Delete(id string) error
	Undelete(id string) error


	FindById(id string) (*Commodity, error)
	FindInfoPage(num, size int) ([]*Commodity, error)
	FindDetailById(id string) (*dto.CommodityDetailInfo, error)
}