package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"net/http"
)

func HandleError(ctx *gin.Context, errorCode enum.ErrorCode, err error) {

	msgErr := ""
	if err == nil {
		msgErr = enum.MsgErr(errorCode)
	} else {
		msgErr = err.Error()
	}

	res := &ResultResponse{
		Error:   int(errorCode),
		Message: msgErr,
		Data:    map[string]interface{}{},
	}

	ctx.JSON(http.StatusOK, res)
}
