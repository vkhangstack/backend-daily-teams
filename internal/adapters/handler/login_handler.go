package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
)

func (h *UserHandler) LoginWithKey(ctx *gin.Context) {
	var payload *dto.LoginDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}
	response, err := h.svc.Login(payload.Username, payload.Password)
	if err != nil {
		HandleError(ctx, enum.UserWrongPassword, nil)
		return
	}

	data := map[string]interface{}{
		"user": map[string]interface{}{
			"id": response.ID,
		},
		"accessToken":  response.AccessToken,
		"refreshToken": response.RefreshToken,
	}

	HandleSuccess(ctx, data)
}

//func (h *UserHandler) LoginUserSocial(ctx *gin.Context) {
//	var payload *model.UserLoginSocial
//	if err := ctx.ShouldBindJSON(&payload); err != nil {
//		HandleError(ctx, enum.BadRequest, err)
//		return
//	}
//	response, err := h.svc.LoginSocial(payload)
//	if err != nil {
//		if strings.Contains(err.Error(), "invalid email") {
//			HandleError(ctx, enum.InvalidEmail, err)
//			return
//		}
//		HandleError(ctx, enum.InternalServerError, nil)
//		return
//	}
//
//	HandleSuccess(ctx, response)
//}
