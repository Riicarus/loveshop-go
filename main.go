package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/conf"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/route"
	"github.com/riicarus/loveshop/pkg/connection"
)

func main() {
	conf.InitConfig()

	connection.InitRedisConn()

	router := gin.Default()

	svctx := context.NewServiceContext()
	route.RegisterHandlers(router, svctx)

	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.ServiceConf.Server.Port), router)
	if err != nil {
		panic(err)
	}
}
