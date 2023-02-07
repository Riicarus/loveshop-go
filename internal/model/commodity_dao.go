package model

import (
	"github.com/riicarus/loveshop/internal/sql"
	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/logic"
	"gorm.io/gorm"
)

type DefaultCommodityModel struct{
	logic.DBModel
}

func (m *DefaultCommodityModel) Conn(txctx *connection.TxContext) CommodityModel {
	m.Txctx = txctx
	
	return m
}

func (m *DefaultCommodityModel) Add(commodity *Commodity) error {
	err := m.Txctx.Tx.Create(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Update(commodity *Commodity) error {
	// the sturct in Updates() must be the database mapping model
	err := m.Txctx.Tx.Model(commodity).Updates(commodity).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) UpdateAmount(id string, number int) error {
	commodity := &Commodity{
		Id: id,
	}

	err := m.Txctx.Tx.Model(commodity).Where("amount + ? >= 0", number).Update("amount", gorm.Expr("amount + ?", number)).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Delete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := m.Txctx.Tx.Model(commodity).Update("deleted", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) Undelete(id string) error {
	commodity := &Commodity{
		Id: id,
	}
	err := m.Txctx.Tx.Model(commodity).Update("deleted", false).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultCommodityModel) FindById(id string) (*Commodity, error) {
	commodity := &Commodity{}
	err := m.Txctx.Tx.Where("id = ?", id).Find(commodity).Error
	if err != nil {
		return nil, err
	}

	return commodity, nil
}

func (m *DefaultCommodityModel) FindByIsbn(isbn string) (*Commodity, error) {
	commodity := &Commodity{}
	err := m.Txctx.Tx.Raw(sql.CommodityFindDetailByIsbn, isbn).Scan(commodity).Error
	if err != nil {
		return nil, err
	}

	return commodity, nil
}

func (m *DefaultCommodityModel) FindPage(num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.Txctx.Tx.Raw(sql.CommodityFindPage, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByType(t string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.Txctx.Tx.Raw(sql.CommodityFindPageByType, t, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByFuzzyName(name string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.Txctx.Tx.Raw(sql.CommodityFindPageByFuzzyName, name, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}

func (m *DefaultCommodityModel) FindPageByFuzzyNameAndType(name, t string, num, size int) ([]*Commodity, error) {
	commoditySlice := make([]*Commodity, 0, size)
	err := m.Txctx.Tx.Raw(sql.CommodityFindPageByFuzzyNameAndType, t, name, num, size).Scan(&commoditySlice).Error
	if err != nil {
		return nil, err
	}

	return commoditySlice, nil
}
