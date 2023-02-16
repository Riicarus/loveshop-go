package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/service"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/resp"
	"github.com/riicarus/loveshop/pkg/util"
)

func OrderAdd(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := &dto.OrderAddParam{}
		err := ctx.Bind(param)
		if err != nil {
			fmt.Println("handler OrderAdd(), binding err: ", err)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_ERR.Msg, e.VALIDATE_ERR.Code))
			return
		}

		if err := param.Check(); err != nil {
			fmt.Println("handler OrderAdd(), validate err: ", err)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_ERR.Msg, e.VALIDATE_ERR.Code))
			return
		}

		orderService := service.NewOrderService(svcctx)
		err2 := orderService.Create(ctx, param)
		switch err2 {
		case e.STOCK_ERR:
			ctx.JSON(http.StatusOK, resp.Fail[string](e.STOCK_ERR.Msg, e.STOCK_ERR.Code))
		case nil:
			ctx.JSON(http.StatusOK, resp.OK(""))
		default:
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		}
	}
}

func OrderCancel(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		orderService := service.NewOrderService(svcctx)
		err := orderService.CancelOrder(ctx, id)
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

func OrderFindDetailAdminViewPageOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.OrderDetailAdminView](num, size)

		orderService := service.NewOrderService(svcctx)
		err4 := orderService.FindDetailAdminViewPageOrderByTime(ctx, desc, page)
		if err4 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func OrderFindDetailAdminViewPageByStatusOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Param("status")
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.OrderDetailAdminView](num, size)

		orderService := service.NewOrderService(svcctx)
		err4 := orderService.FindDetailAdminViewPageByStatusOrderByTime(ctx, status, desc, page)
		if err4 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func OrderFindDetailUserViewPageByUidOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.OrderDetailUserView](num, size)

		orderService := service.NewOrderService(svcctx)
		err4 := orderService.FindDetailUserViewPageByUidOrderByTime(ctx, uid, desc, page)
		if err4 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func OrderFindDetailUserViewPageByUidAndStatusOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		status := ctx.Param("status")
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.OrderDetailUserView](num, size)

		orderService := service.NewOrderService(svcctx)
		err4 := orderService.FindDetailUserViewPageByUidAndStatusOrderByTime(ctx, uid, status, desc, page)
		if err4 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}
