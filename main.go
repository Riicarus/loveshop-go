package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/conf"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/route"
	"github.com/riicarus/loveshop/internal/service"
	"github.com/riicarus/loveshop/pkg/connection"
)

func main() {
	conf.InitConfig()

	connection.InitRedisConn()

	router := gin.Default()

	svctx := context.NewServiceContext()
	route.RegisterHandlers(router, svctx)

	// kafka consumer
	orderService := service.NewOrderService(svctx)
	orderService.ConsumeFromKafka(&gin.Context{})

	// cache commodity to redis
	commodityService := service.NewCommodityService(svctx)
	if err := commodityService.CacheStock(&gin.Context{}); err != nil {
		panic(err)
	}

	go func ()  {
		for {
			time.Sleep(10 * time.Second)

			if err := commodityService.StoreStock(&gin.Context{}); err != nil {
				fmt.Println("commodity stock store failed, err: ", err)
			}
		}
	}()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.ServiceConf.Server.Port), router); err != nil {
		panic(err)
	}
}
