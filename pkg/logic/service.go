package logic

import "github.com/riicarus/loveshop/conf"

type ServiceLogic struct {
	Svcctx *ServiceContext
}

type ServiceContext struct {
	Conf *conf.ServiceConfig
}