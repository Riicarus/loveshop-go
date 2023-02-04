package dto

import "github.com/riicarus/loveshop/pkg/util"

type AdminLoginParam struct {
	StudentId string `json:"studentId" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AdminRegisterParam struct {
	Name      string               `json:"name" binding:"required"`
	StudentId string               `json:"studentId" binding:"required"`
	Password  string               `json:"password" binding:"required"`
	Email     string               `json:"email" binding:"required"`
	Group     string               `json:"group" binding:"required"`
	Roles     util.JSONStringSlice `json:"roles" binding:"required"`
}
