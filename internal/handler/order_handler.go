package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/service"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/resp"
)

func OrderAdd(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := &dto.OrderAddParam{}
		err := ctx.Bind(param)
		if err != nil {
			fmt.Println("handler OrderAdd(), binding err: ", err)
		}

		orderService := service.NewOrderService(svcctx)
		err2 := orderService.Add(ctx, param)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func OrderCancle(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		orderService := service.NewOrderService(svcctx)
		err := orderService.CancleOrder(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func OrderPay(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		orderService := service.NewOrderService(svcctx)
		err := orderService.PayOrder(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func OrderFinish(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		orderService := service.NewOrderService(svcctx)
		err := orderService.FinishOrder(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}
