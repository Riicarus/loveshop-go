package model

import (
	"github.com/riicarus/loveshop/internal/constant"
	"github.com/riicarus/loveshop/internal/sql"
	"github.com/riicarus/loveshop/pkg/logic"
	"gorm.io/gorm"
)

type DefaultOrderModel struct {
	logic.DBModel
}

func (m *DefaultOrderModel) Conn(db *gorm.DB) OrderModel {
	m.DB = db

	return m
}

func (m *DefaultOrderModel) Add(order *Order) error {
	err := m.DB.Debug().Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultOrderModel) CancleOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := m.DB.Model(order).Where("status IN('CREATED', 'PAYED')").Update("status", constant.ORDER_STATUS_CANCLED).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultOrderModel) PayOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := m.DB.Model(order).Where("status = 'CREATED'").Update("status", constant.ORDER_STATUS_PAYED).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *DefaultOrderModel) FinishOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := m.DB.Model(order).Where("status = 'PAYED'").Update("status", constant.ORDER_STATUS_FINISHED).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultOrderModel) FindById(id string) (*Order, error) {
	order := &Order{}
	err := m.DB.Where("id = ?").Find(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (m *DefaultOrderModel) FindDetailById(id string) (*OrderDetail, error) {
	orderDetail := &OrderDetail{}

	if err := m.DB.Raw(sql.OrderFindDetailAdminViewById, id).Scan(orderDetail).Error; err != nil {
		return nil, err
	}

	return orderDetail, nil
}

func (m *DefaultOrderModel) FindPageOrderByTime(desc bool, num, size int) ([]*OrderDetail, error) {
	orderSlice := make([]*OrderDetail, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.OrderFindPageOrderByTimeDesc
	} else {
		sqlUse = sql.OrderFindPageOrderByTimeAsc
	}

	err := m.DB.Raw(sqlUse, num, size).Scan(&orderSlice).Error
	if err != nil {
		return nil, err
	}
	return orderSlice, nil
}

func (m *DefaultOrderModel) FindPageByStatusOrderByTime(status string, desc bool, num, size int) ([]*OrderDetail, error) {
	orderSlice := make([]*OrderDetail, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.OrderFindPageByStatusOrderByTimeDesc
	} else {
		sqlUse = sql.OrderFindPageByStatusOrderByTimeAsc
	}

	err := m.DB.Raw(sqlUse, status, num, size).Scan(&orderSlice).Error
	if err != nil {
		return nil, err
	}
	return orderSlice, nil
}

func (m *DefaultOrderModel) FindUserViewPageByUidOrderByTime(uid string, desc bool, num, size int) ([]*Order, error) {
	orderSlice := make([]*Order, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.OrderFindUserViewPageByUidOrderByTimeDesc
	} else {
		sqlUse = sql.OrderFindUserViewPageByUidOrderByTimeAsc
	}

	err := m.DB.Raw(sqlUse, uid, num, size).Scan(&orderSlice).Error
	if err != nil {
		return nil, err
	}
	return orderSlice, nil
}

func (m *DefaultOrderModel) FindUserViewPageByUidAndStatusOrderByTime(uid, status string, desc bool, num, size int) ([]*Order, error) {
	orderSlice := make([]*Order, 0)

	var sqlUse string
	if desc {
		sqlUse = sql.OrderFindUserViewPageByUidAndStatusOrderByTimeDesc
	} else {
		sqlUse = sql.OrderFindUserViewPageByUidAndStatusOrderByTimeAsc
	}

	err := m.DB.Raw(sqlUse, uid, status, num, size).Scan(&orderSlice).Error
	if err != nil {
		return nil, err
	}
	return orderSlice, nil
}
