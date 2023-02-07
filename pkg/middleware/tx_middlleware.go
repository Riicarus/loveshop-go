package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/resp"
)

func TxMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		txctx, err := connection.NewTxContext()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, resp.Fail[string](e.INTERNAL_ERROR_MSG, e.INTERNAL_ERROR_CODE))
			ctx.Abort()
		}
		ctx.Set("txctx", txctx)
		ctx.Next()
	}
}