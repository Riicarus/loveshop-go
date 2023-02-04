package model

import "github.com/riicarus/loveshop/pkg/util"

type Role struct {
	Id            string               `json:"id"`
	Role          string               `json:"role"`
	PermissionIds util.JSONStringSlice `gorm:"TYPE:json" json:"permissionIds"`
}

func (Role) TableName() string {
	return "role"
}