package model

import (
	"github.com/riicarus/loveshop/pkg/logic"
	"github.com/riicarus/loveshop/pkg/util"
)

type Admin struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	StudentId   string               `json:"studentId"`
	Password    string               `json:"-"`
	Email       string               `json:"email"`
	Salt        string               `json:"-"`
	Group       string               `json:"group"`
	Integration float32              `json:"integration"`
	RoleIds     util.JSONStringSlice `gorm:"TYPE:json" json:"roleIds"`
	Enabled     bool                 `json:"enabled"`
}

func (Admin) TableName() string {
	return "admin"
}

type AdminModel interface {
	logic.IDBModel[AdminModel]

	FindById(id string) (*Admin, error)
	FindByStudentId(studentId string) (*Admin, error)

	Register(admin *Admin) error
	Unable(id string) error
}