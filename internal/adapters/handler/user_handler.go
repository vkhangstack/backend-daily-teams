package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"github.com/vkhangstack/dlt/internal/core/services"
	"net/http"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(UserService services.UserService) *UserHandler {
	return &UserHandler{
		svc: UserService,
	}
}

//	func (h *UserHandler) CreateUser(ctx *gin.Context) {
//		var user model.User
//		if err := ctx.ShouldBindJSON(&user); err != nil {
//			HandleError(ctx, http.StatusBadRequest, err)
//			return
//		}
//		//
//		for i := 1; i < 100000; i++ {
//			_, err := h.svc.CreateUser(uuid.NewString()+"@gmail", utils.GenerateID())
//
//			if err != nil {
//				HandleError(ctx, http.StatusBadRequest, err)
//				return
//			}
//		}
//
//		//ctx.JSON(http.StatusCreated, gin.H{
//		//	"message": "New user created successfully",
//		//})
//		result := map[string]interface{}{
//			"message": "New user created successfully",
//		}
//		HandleSuccess(ctx, result)
//	}

func (h *UserHandler) Register(ctx *gin.Context) {

	var payload *dto.RegisterDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	err := h.svc.Register(payload)
	if err != nil {
		HandleError(ctx, enum.UseAlreadyExits, err)
		return
	}
	HandleSuccess(ctx, nil)

}
func (h *UserHandler) ProfileMe(ctx *gin.Context) {
	userId, err := utils.GetUserId(ctx)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	userId64, err := utils.TransformStringToUInt64(userId)
	if err != nil {
		HandleError(ctx, enum.BadRequest, fmt.Errorf("malformed request data"))
		return
	}

	user, err := h.svc.ProfileMe(userId64)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}
	data := map[string]interface{}{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"avatarUrl": user.AvatarURL,
	}
	HandleSuccess(ctx, data)
}

func (h *UserHandler) GetAccessToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	users, err := h.svc.GetAccessToken(token)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	HandleSuccess(ctx, users)
}
