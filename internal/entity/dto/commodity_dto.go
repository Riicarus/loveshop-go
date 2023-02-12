package dto

import (
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/riicarus/loveshop/internal/consts"
	"github.com/riicarus/loveshop/pkg/e"
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

func (c *CommodityAddParam) Check() error {
	switch c.Type {
	case consts.BOOK_TYPE:
		// check isbn, author, publisher
	case consts.CULTURAL_CREATIVITY_TYPE:
	case consts.BOARD_GAME_TYPE:
	case consts.SPORTS_GOODS_TYPE:
	case consts.DAILY_NECESSITY_TYPE:
	default:
		return e.VALIDATE_ERR
	}

	// check numbering

	if strings.TrimSpace(c.Name) == "" {
		return e.VALIDATE_ERR
	}

	if c.Amount < 0 {
		return e.VALIDATE_ERR
	}

	if c.Price < 0 {
		return e.VALIDATE_ERR
	}

	return nil
}

type CommodityUpdateParam struct {
	Id        string       `json:"id" binding:"required"`
	Type      string       `json:"type"`
	Numbering string       `json:"numbering"`
	Name      string       `json:"name"`
	Price     float64      `json:"price"`
	Extension util.JSONMap `json:"extension"`
}

func (c *CommodityUpdateParam) Check() error {
	if strings.TrimSpace(c.Id) == "" {
		return e.VALIDATE_ERR
	}

	switch c.Type {
	case consts.BOOK_TYPE:
		// check isbn, author, publisher
		isbn, _ := c.Extension["ISBN"].(string)
		r := regexp2.MustCompile(consts.ISBN_REG, 0)
		if ok, _ := r.MatchString(isbn); !ok {
			return e.VALIDATE_ERR
		}

		if _, ok := c.Extension["author"]; !ok {
			return e.VALIDATE_ERR
		}
		if _, ok := c.Extension["publisher"]; !ok {
			return e.VALIDATE_ERR
		}
	case consts.CULTURAL_CREATIVITY_TYPE:
	case consts.BOARD_GAME_TYPE:
	case consts.SPORTS_GOODS_TYPE:
	case consts.DAILY_NECESSITY_TYPE:
	default:
		return e.VALIDATE_ERR
	}

	// check numbering
	r := regexp2.MustCompile(consts.POSITION_REG, 0)
	if ok, _ := r.MatchString(c.Numbering); !ok {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(c.Name) == "" {
		return e.VALIDATE_ERR
	}

	if c.Price < 0 {
		return e.VALIDATE_ERR
	}

	return nil
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
