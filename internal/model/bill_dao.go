package model

import (
	"github.com/riicarus/loveshop/internal/sql"
	"github.com/riicarus/loveshop/pkg/logic"
	"gorm.io/gorm"
)

type DefaultBillModel struct {
	logic.DBModel
}

func (m *DefaultBillModel) Conn(db *gorm.DB) BillModel {
	m.DB = db

	return m
}

func (m *DefaultBillModel) Add(bill *Bill) error {
	err := m.DB.Create(bill).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultBillModel) FindById(id string) (*Bill, error) {
	bill := &Bill{}
	if err := m.DB.Where("id = ?", id).Find(bill).Error; err != nil {
		return nil, err
	}

	return bill, nil
}

func (m *DefaultBillModel) FindPageOrderByTime(desc bool, num, size int) ([]*Bill, error) {
	billSlice := make([]*Bill, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.BillFindPageOrderByTimeDesc
	} else {
		sqlUse = sql.BillFindPageOrderByTimeAsc
	}
	if err := m.DB.Raw(sqlUse, num, size).Scan(&billSlice).Error; err != nil {
		return nil, err
	}

	return billSlice, nil
}

func (m *DefaultBillModel) FindPageByOrderTypeOrderByTime(orderType string, desc bool, num, size int) ([]*Bill, error) {
	billSlice := make([]*Bill, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.BillFindPageByOrderTypeOrderByTimeDesc
	} else {
		sqlUse = sql.BillFindPageByOrderTypeOrderByTimeAsc
	}
	if err := m.DB.Raw(sqlUse, orderType, num, size).Scan(&billSlice).Error; err != nil {
		return nil, err
	}

	return billSlice, nil
}
