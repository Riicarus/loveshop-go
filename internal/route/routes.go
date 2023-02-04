package route

import (
	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/handler"
)

func RegisterHandlers(router *gin.Engine, svcctx *context.ServiceContext) {
	adminGroupV1 := router.Group("/api/admin/v1")
	{
		adminGroupV1.POST("/login", handler.AdminLogin(svcctx))
		adminGroupV1.PUT("/register", handler.AdminRegister(svcctx))
		adminGroupV1.PUT("/unable/:id", handler.UnableAdmin(svcctx))
	}

	commodityGroupV1 := router.Group("/api/commodity/v1")
	{
		commodityGroupV1.POST("", handler.CommodityAdd(svcctx))

		commodityGroupV1.PUT("", handler.CommodityUpdate(svcctx))
		commodityGroupV1.PUT("/amount/:id/:number", handler.CommodityUpdateAmount(svcctx))

		commodityGroupV1.DELETE("/:id", handler.CommodityDelete(svcctx))
		commodityGroupV1.PUT("/active/:id", handler.CommodityUndelete(svcctx))

		commodityGroupV1.GET("/info/:num/:size", handler.CommodityFindInfoPage(svcctx))
		commodityGroupV1.GET("/info/detail/:id", handler.CommodityFindDetailInfo(svcctx))
	}
}