package context

import (
	"github.com/riicarus/loveshop/conf"
	"github.com/riicarus/loveshop/internal/model"
)

type ServiceContext struct {
	Config *conf.ServiceConfig

	// model
	AdminModel     model.AdminModel
	UserModel      model.UserModel
	CommodityModel model.CommodityModel
	OrderModel     model.OrderModel
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		Config: conf.ServiceConf,

		AdminModel:     &model.DefaultAdminModel{},
		UserModel:      &model.DefaultUserModel{},
		CommodityModel: &model.DefaultCommodityModel{},
		OrderModel:     &model.DefaultOrderModel{},
	}
}
