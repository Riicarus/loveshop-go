package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/sql"
	"github.com/riicarus/loveshop/pkg/connection"
	"gorm.io/gorm"
)

type DefaultCommodityModel struct{}

func (m *DefaultCommodityModel) Add(param *dto.CommodityAddParam) error {
	commodity := &Commodity{
		Id: uuid.New().String(),
		Type: param.Type,
		Numbering: param.Numbering,
		Name: param.Name,
		Amount: param.Amount,
		Price: param.Price,
		Extension: param.Extension,
		Deleted: false,
	}
	err := connection.SqlConn.Create(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Update(param *dto.CommodityUpdateParam) error {
	commodity := &Commodity{
		Id: param.Id,
		Numbering: param.Numbering,
		Name: param.Name,
		Type: param.Type,
		Price: param.Price,
		Extension: param.Extension,
	}

	// the sturct in Updates() must be the database mapping model
	err := connection.SqlConn.Model(commodity).Updates(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) UpdateAmount(id string, number int) error {
	commodity := &Commodity{
		Id: id,
	}

	err := connection.SqlConn.Model(commodity).Where("amount + ? >= 0", number).Update("amount", gorm.Expr("amount + ?", number)).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Delete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := connection.SqlConn.Model(commodity).Update("deleted", true).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (m *DefaultCommodityModel) Undelete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := connection.SqlConn.Model(commodity).Update("deleted", false).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (m *DefaultCommodityModel) FindById(id string) (*Commodity, error) {
	commodity := &Commodity{}
	err := connection.SqlConn.Where("id = ?", id).Find(commodity).Error
	if err != nil {
		return nil, err
	}

	return commodity, nil
}

func (m *DefaultCommodityModel) FindInfoPage(num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0)
	err := connection.SqlConn.Raw(sql.CommodityFindInfoPage, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindDetailById(id string) (*dto.CommodityDetailInfo, error) {
	c, err := m.FindById(id)
	if err != nil {
		fmt.Println("DefaultCommodityModel.FindDetailById err: ", err)
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	return &dto.CommodityDetailInfo{
		Id:        c.Id,
		Type:      c.Type,
		Numbering: c.Numbering,
		Name:      c.Name,
		Amount:    c.Amount,
		Price:     c.Price,
		Extension: c.Extension,
		Deleted:   c.Deleted,
	}, nil
}
