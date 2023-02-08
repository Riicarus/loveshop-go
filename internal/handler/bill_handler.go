package handler

import (
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

func BillFindDetailAdminViewById(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		billService := service.NewBillService(svcctx)
		detailView, err := billService.FindDetailAdminViewById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(detailView))
		}
	}
}

func BillFindDetailAdminViewPageOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.BillDetailAdminView](num, size)

		billService := service.NewBillService(svcctx)
		err := billService.FindDetailAdminViewPageOrderByTime(ctx, desc, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func BillFindDetailAdminViewPageByOrderTypeOrderByTime(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderType := ctx.Param("orderType")
		desc, err1 := strconv.ParseBool(ctx.Param("desc"))
		num, err2 := strconv.Atoi(ctx.Param("num"))
		size, err3 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil || err3 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.BillDetailAdminView](num, size)

		billService := service.NewBillService(svcctx)
		err := billService.FindDetailAdminViewPageByOrderTypeOrderByTime(ctx, orderType, desc, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func BillAnalyze(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		billService := service.NewBillService(svcctx)
		analyzeInfo, err := billService.Analyze(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(analyzeInfo))
		}
	}
}