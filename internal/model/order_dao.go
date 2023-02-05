package model

import (
	"github.com/riicarus/loveshop/internal/constant"
	"github.com/riicarus/loveshop/pkg/connection"
)

type DefaultOrderModel struct{}

func (m *DefaultOrderModel) Add(order *Order) error {
	err := connection.SqlConn.Create(order).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (m *DefaultOrderModel) CancleOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := connection.SqlConn.Model(order).Where("status IN('CREATED', 'PAYED')").Update("status", constant.ORDER_STATUS_CANCLED).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultOrderModel) PayOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := connection.SqlConn.Model(order).Where("status = 'CREATED'").Update("status", constant.ORDER_STATUS_PAYED).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *DefaultOrderModel) FinishOrder(id string) error {
	order := &Order{
		Id: id,
	}
	err := connection.SqlConn.Model(order).Where("status = 'PAYED'").Update("status", constant.ORDER_STATUS_FINISHED).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *DefaultOrderModel) FindById(id string) (*Order, error) {
	order := &Order{}
	err := connection.SqlConn.Where("id = ?").Find(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}