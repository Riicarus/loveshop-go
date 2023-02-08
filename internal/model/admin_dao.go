package model

import (
	"fmt"

	"github.com/riicarus/loveshop/pkg/logic"
	"gorm.io/gorm"
)

type DefaultAdminModel struct {
	logic.DBModel
}

func (m *DefaultAdminModel) Conn(db *gorm.DB) AdminModel {
	m.DB = db

	return m
}

func (m *DefaultAdminModel) FindById(id string) (*Admin, error) {
	admin := &Admin{}
	err := m.DB.Where("id = ?", id).Find(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.FindById() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) FindByStudentId(studentId string) (*Admin, error) {
	admin := &Admin{}
	err := m.DB.Where("student_id = ?", studentId).Find(admin).Error

	if err != nil {
		fmt.Println("DefaultAdminModel.FindByStudentId() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) Unable(id string) error {
	err := m.DB.Model(&Admin{}).Where("id = ?", id).Update("enabled", false).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Unable() err: ", err)
		return err
	}

	return nil
}

func (m *DefaultAdminModel) Register(admin *Admin) error {
	err := m.DB.Save(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Register() err: ", err)
		return err
	}

	return nil
}
