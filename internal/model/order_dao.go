package model

import (
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
	return nil
}

func (m *DefaultOrderModel) PayOrder(id string) error {
	return nil
}

func (m *DefaultOrderModel) FinishOrder(id string) error {
	return nil
}

func (m *DefaultOrderModel) FindById(id string) (*Order, error) {
	return nil, nil
}