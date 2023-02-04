package model

import (
	"fmt"

	"github.com/riicarus/loveshop/pkg/connection"
)

type DefaultAdminModel struct {}

func (m *DefaultAdminModel) FindById(id string) (*Admin, error) {
	admin := &Admin{}
	err := connection.SqlConn.Where("id = ?", id).Find(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.FindById() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) FindByStudentId(studentId string) (*Admin, error) {
	admin := &Admin{}
	err := connection.SqlConn.Where("student_id = ?", studentId).Find(admin).Error

	if err != nil {
		fmt.Println("DefaultAdminModel.FindByStudentId() err: ", err)
		return nil, err
	}

	return admin, nil
}

func (m *DefaultAdminModel) Unable(id string) error {
	err := connection.SqlConn.Model(&Admin{}).Where("id = ?", id).Update("enabled", false).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Unable() err: ", err)
		return err
	}

	return nil
}

func (m *DefaultAdminModel) Register(admin *Admin) error {
	err := connection.SqlConn.Save(admin).Error
	if err != nil {
		fmt.Println("DefaultAdminModel.Register() err: ", err)
		return err
	}
	
	return nil
}