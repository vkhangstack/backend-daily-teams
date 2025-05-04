package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   int         `json:"error"`
}

func HandleSuccess(ctx *gin.Context, result any) {

	res := &ResultResponse{
		Error:   0,
		Message: "Success",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, res)
}
