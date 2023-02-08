package model

import (
	"github.com/riicarus/loveshop/internal/sql"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/logic"
	"gorm.io/gorm"
)

type DefaultCommodityModel struct {
	logic.DBModel
}

func (m *DefaultCommodityModel) Conn(db *gorm.DB) CommodityModel {
	m.DB = db

	return m
}

func (m *DefaultCommodityModel) Add(commodity *Commodity) error {
	err := m.DB.Create(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Update(commodity *Commodity) error {
	// the sturct in Updates() must be the database mapping model
	err := m.DB.Model(commodity).Updates(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) UpdateAmount(id string, number int) error {
	commodity := &Commodity{
		Id: id,
	}

	db := m.DB.Model(commodity).Where("amount + ? >= 0", number).Update("amount", gorm.Expr("amount + ?", number))
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected != 1 {
		return e.STOCK_ERR
	}

	return nil
}

func (m *DefaultCommodityModel) Delete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := m.DB.Model(commodity).Update("deleted", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Undelete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := m.DB.Model(commodity).Update("deleted", false).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) FindById(id string) (*Commodity, error) {
	commodity := &Commodity{}
	err := m.DB.Where("id = ?", id).Find(commodity).Error
	if err != nil {
		return nil, err
	}

	return commodity, nil
}

func (m *DefaultCommodityModel) FindByIsbn(isbn string) (*Commodity, error) {
	commodity := &Commodity{}
	err := m.DB.Raw(sql.CommodityFindDetailByIsbn, isbn).Scan(commodity).Error
	if err != nil {
		return nil, err
	}

	return commodity, nil
}

func (m *DefaultCommodityModel) FindPage(num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.DB.Raw(sql.CommodityFindPage, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByType(t string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.DB.Raw(sql.CommodityFindPageByType, t, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByFuzzyName(name string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.DB.Raw(sql.CommodityFindPageByFuzzyName, name, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByFuzzyNameAndType(name, t string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.DB.Raw(sql.CommodityFindPageByFuzzyNameAndType, t, name, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}
