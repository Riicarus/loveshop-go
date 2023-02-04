package model

import "github.com/riicarus/loveshop/pkg/util"

type User struct {
	Id        string               `json:"id"`
	Name      string               `json:"name"`
	StudentId string               `json:"studentId"`
	Password  string               `json:"-"`
	Email     string               `json:"email"`
	Salt      string               `json:"-"`
	RoleIds   util.JSONStringSlice `gorm:"TYPE:json" json:"roleIds"`
	Enabled   bool                 `json:"enabled"`
}

func (User) TableName() string {
	return "user"
}

type UserModel interface {
	FindById(id string) (*User, error)
	FindByStudentId(studentId string) (*User, error)

	// Register() error
	Unable(id string) error
}