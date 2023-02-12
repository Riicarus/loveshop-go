package dto

import (
	"strings"

	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/util"
)

type AdminLoginParam struct {
	StudentId string `json:"studentId" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AdminLoginView struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StudentId string `json:"studentId"`
	Type      string `json:"type"`
	Token     string `json:"token"`
}

type AdminRegisterParam struct {
	Name      string               `json:"name" binding:"required"`
	StudentId string               `json:"studentId" binding:"required"`
	Password  string               `json:"password" binding:"required"`
	Email     string               `json:"email" binding:"required"`
	Group     string               `json:"group" binding:"required"`
	Roles     util.JSONStringSlice `json:"roles" binding:"required"`
}

func (a *AdminRegisterParam) Check() error {
	if strings.TrimSpace(a.Name) == "" {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(a.StudentId) == "" {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(a.Password) == "" {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(a.Email) == "" {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(a.Group) == "" {
		return e.VALIDATE_ERR
	}

	return nil
}
