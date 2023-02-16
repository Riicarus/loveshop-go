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

		commodityGroupV1.GET("/view/detail/:id", handler.CommodityFindDetailViewById(svcctx))
		commodityGroupV1.GET("/view/detail/isbn/:isbn", handler.CommodityFindDetailViewByIsbn(svcctx))

		commodityGroupV1.GET("/view/simple/:num/:size", handler.CommodityFindSimpleViewPage(svcctx))
		commodityGroupV1.GET("/view/simple/type/:type/:num/:size", handler.CommodityFindSimpleViewPageByType(svcctx))
		commodityGroupV1.GET("/view/simple/name/:name/:num/:size", handler.CommodityFindSimpleViewPageByFuzzyName(svcctx))
		commodityGroupV1.GET("/view/simple/name/:name/type/:type/:num/:size", handler.CommodityFindSimpleViewPageByFuzzyNameAndType(svcctx))
	}

	orderGroupV1 := router.Group("/api/order/v1")
	{
		orderGroupV1.POST("", handler.OrderAdd(svcctx))

		orderGroupV1.PUT("/cancel/:id", handler.OrderCancel(svcctx))
		orderGroupV1.PUT("/pay/:id", handler.OrderPay(svcctx))
		orderGroupV1.PUT("/finish/:id", handler.OrderFinish(svcctx))

		orderGroupV1.GET("/view/detail/admin/:desc/:num/:size", handler.OrderFindDetailAdminViewPageOrderByTime(svcctx))
		orderGroupV1.GET("/view/detail/admin/status/:status/:desc/:num/:size", handler.OrderFindDetailAdminViewPageByStatusOrderByTime(svcctx))
		orderGroupV1.GET("/view/detail/user/:uid/:desc/:num/:size", handler.OrderFindDetailUserViewPageByUidOrderByTime(svcctx))
		orderGroupV1.GET("/view/detail/user/:uid/status/:status/:desc/:num/:size", handler.OrderFindDetailUserViewPageByUidAndStatusOrderByTime(svcctx))
	}

	billGroupV1 := router.Group("/api/bill/v1")
	{
		billGroupV1.GET("/view/detail/admin/id/:id", handler.BillFindDetailAdminViewById(svcctx))
		billGroupV1.GET("/view/detail/admin/:desc/:num/:size", handler.BillFindDetailAdminViewPageOrderByTime(svcctx))
		billGroupV1.GET("/view/detail/admin/type/:orderType/:desc/:num/:size", handler.BillFindDetailAdminViewPageByOrderTypeOrderByTime(svcctx))
		billGroupV1.GET("/analyze", handler.BillAnalyze(svcctx))
	}
}
