package context

import (
	"github.com/riicarus/loveshop/conf"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/connection"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Conf *conf.ServiceConfig

	// gorm
	DB *gorm.DB

	// model
	AdminModel     model.AdminModel
	UserModel      model.UserModel
	CommodityModel model.CommodityModel
	OrderModel     model.OrderModel
	BillModel      model.BillModel
}

func NewServiceContext() *ServiceContext {
	db, err := connection.NewSqlConn()
	if err != nil {
		panic("Could not connect to database")
	}

	return &ServiceContext{
		Conf: conf.ServiceConf,

		DB: db,

		AdminModel:     &model.DefaultAdminModel{},
		UserModel:      &model.DefaultUserModel{},
		CommodityModel: &model.DefaultCommodityModel{},
		OrderModel:     &model.DefaultOrderModel{},
		BillModel:      &model.DefaultBillModel{},
	}
}
