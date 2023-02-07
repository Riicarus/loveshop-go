package model

import (
	"fmt"

	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/logic"
)

// TODO add Txctx check
type DefaultAdminModel struct {
	logic.DBModel
}

func (m *DefaultAdminModel) Conn(txctx *connection.TxContext) AdminModel {
	m.Txctx = txctx

	return m
}

func (m *DefaultAdminModel) FindById(id string) (*Admin, error) {
	admin := &Admin{}
	err := m.Txctx.DB().Where("id = ?", id).Find(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.FindById() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) FindByStudentId(studentId string) (*Admin, error) {
	admin := &Admin{}
	err := m.Txctx.DB().Where("student_id = ?", studentId).Find(admin).Error

	if err != nil {
		fmt.Println("DefaultAdminModel.FindByStudentId() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) Unable(id string) error {
	err := m.Txctx.DB().Model(&Admin{}).Where("id = ?", id).Update("enabled", false).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Unable() err: ", err)
		return err
	}

	return nil
}

func (m *DefaultAdminModel) Register(admin *Admin) error {
	err := m.Txctx.DB().Save(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Register() err: ", err)
		return err
	}

	return nil
}
