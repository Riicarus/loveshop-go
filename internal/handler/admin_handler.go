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

func AdminLogin(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginParam dto.AdminLoginParam
		err1 := ctx.Bind(&loginParam)
		if err1 != nil {
			fmt.Println("handler AdminLogin(), databind err: ", err1)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		adminService := service.NewAdminService(svcctx)
		token, err2 := adminService.LoginWithPass(ctx, &loginParam)
		if err2 == e.UNAUTHED_ERR {
			ctx.JSON(http.StatusOK, resp.Fail[string](e.UNAUTHED_ERR.Msg, e.UNAUTHED_ERR.Code))
		} else if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.UNATHENRIZED_MSG, e.UNATHENRIZED_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(token))
		}
	}
}

func UnableAdmin(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		adminService := service.NewAdminService(svcctx)
		err := adminService.Unable(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}

func AdminRegister(svcctx *context.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerParam dto.AdminRegisterParam
		err1 := ctx.Bind(&registerParam)
		if err1 != nil {
			fmt.Println("handler AdminRegister(), databind err: ", err1)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		if err := registerParam.Check(); err != nil {
			fmt.Println("handler AdminRegister(), validate err: ", err1)
			ctx.JSON(http.StatusOK, resp.Fail[string](e.VALIDATE_FAILED_MSG, e.VALIDATE_FAILED_CODE))
			return
		}

		adminService := service.NewAdminService(svcctx)
		err2 := adminService.Register(ctx, &registerParam)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
		} else {
			ctx.JSON(http.StatusOK, resp.OK(""))
		}
	}
}
