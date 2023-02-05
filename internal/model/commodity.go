package model

import (
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
	Add(commoidty *Commodity) error

	Update(commodity *Commodity) error
	UpdateAmount(id string, number int) error

	Delete(id string) error
	Undelete(id string) error

	FindById(id string) (*Commodity, error)
	FindByIsbn(isbn string) (*Commodity, error)

	FindPage(num, size int) ([]*Commodity, error)
	FindPageByType(t string, num, size int) ([]*Commodity, error)
	FindPageByFuzzyName(name string, num, size int) ([]*Commodity, error)
	FindPageByFuzzyNameAndType(name, t string, num, size int) ([]*Commodity, error)
}