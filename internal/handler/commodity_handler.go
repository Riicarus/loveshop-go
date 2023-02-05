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

func CommodityAdd(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addParam := &dto.CommodityAddParam{}
		err := ctx.Bind(addParam)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_ERR.Msg, e.VALIDATE_ERR.Code))
			return
		}

		commodityService := service.NewCommodityService(svcctx)
		err2 := commodityService.Add(ctx, addParam)
		if err2 != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func CommodityUpdate(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		updateParam := &dto.CommodityUpdateParam{}
		err := ctx.Bind(updateParam)
		if err != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_ERR.Msg, e.VALIDATE_ERR.Code))
			return
		}

		commodityService := service.NewCommodityService(svcctx)
		err2 := commodityService.Update(ctx, updateParam)
		if err2 != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func CommodityUpdateAmount(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		number, err := strconv.Atoi(ctx.Param("number"))
		if err != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_ERR.Msg, e.VALIDATE_ERR.Code))
			return
		}

		commodityService := service.NewCommodityService(svcctx)
		err1 := commodityService.UpdateAmount(ctx, id, number)
		if err1 != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func CommodityDelete(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		commodityService := service.NewCommodityService(svcctx)
		err1 := commodityService.Delete(ctx, id)
		if err1 != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func CommodityUndelete(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		commodityService := service.NewCommodityService(svcctx)
		err1 := commodityService.Undelete(ctx, id)
		if err1 != nil {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func CommodityFindDetailInfoById(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		commoditySerivce := service.NewCommodityService(svcctx)
		detailInfo, err := commoditySerivce.FindDetailById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(detailInfo))
		}
	}
}

func CommodityFindDetailInfoByIsbn(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isbn := ctx.Param("isbn")
		// TODO use regex to check whether isbn is right-formatted

		commoditySerivce := service.NewCommodityService(svcctx)
		detailInfo, err := commoditySerivce.FindDetailByIsbn(ctx, isbn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(detailInfo))
		}
	}
}

func CommodityFindInfoPage(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		num, err1 := strconv.Atoi(ctx.Param("num"))
		size, err2 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.CommoditySimpleInfo](num, size)

		commodityService := service.NewCommodityService(svcctx)
		err := commodityService.FindInfoPage(ctx, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func CommodityFindInfoPageByType(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.Param("type")
		num, err1 := strconv.Atoi(ctx.Param("num"))
		size, err2 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.CommoditySimpleInfo](num, size)

		commodityService := service.NewCommodityService(svcctx)
		err := commodityService.FindInfoPageByType(ctx, t, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func CommodityFindInfoPageByFuzzyName(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		num, err1 := strconv.Atoi(ctx.Param("num"))
		size, err2 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.CommoditySimpleInfo](num, size)

		commodityService := service.NewCommodityService(svcctx)
		err := commodityService.FindInfoPageByFuzzyName(ctx, name, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}

func CommodityFindInfoPageByFuzzyNameAndType(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		t := ctx.Param("type")
		num, err1 := strconv.Atoi(ctx.Param("num"))
		size, err2 := strconv.Atoi(ctx.Param("size"))
		if err1 != nil || err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		page := util.NewPage[*dto.CommoditySimpleInfo](num, size)

		commodityService := service.NewCommodityService(svcctx)
		err := commodityService.FindInfoPageByFuzzyNameAndType(ctx, name, t, page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(page))
		}
	}
}