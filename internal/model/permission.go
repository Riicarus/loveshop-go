package model

type Permission struct {
	Id         string `json:"id"`
	Permission string `json:"permission"`
}

func (Permission) TableName() string {
	return "permission"
}