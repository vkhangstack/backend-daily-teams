package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"github.com/vkhangstack/dlt/internal/core/services"
)

type DailyHandler struct {
	svc services.DailyService
}

func NewDailyHandler(DailyService services.DailyService) *DailyHandler {
	return &DailyHandler{
		svc: DailyService,
	}
}

func (h *DailyHandler) CreateTask(ctx *gin.Context) {
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

	var payload *dto.CreateTaskDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	task, err := h.svc.CreateTask(payload, userId64)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}
	data := map[string]interface{}{
		"id":              utils.TransformUInt64ToString(task.ID),
		"title":           task.Title,
		"content":         task.Content,
		"start":           task.Start,
		"end":             task.End,
		"textColor":       task.TextColor,
		"backgroundColor": task.BackgroundColor,
		"createdAt":       task.CreatedAt,
		"updatedAt":       task.UpdatedAt,
		"userId":          utils.TransformUInt64ToString(task.UserId),
		"allDay":          task.AllDay,
	}

	HandleSuccess(ctx, data)
}

func (h *DailyHandler) UpdateTask(ctx *gin.Context) {

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

	var payload *dto.UpdateTaskDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	err = h.svc.UpdateTask(payload, userId64)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	HandleSuccess(ctx, map[string]interface{}{})
}

func (h *DailyHandler) DeleteTask(ctx *gin.Context) {

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

	idParam := ctx.Param("id")

	if idParam == "" {
		HandleError(ctx, enum.BadRequest, fmt.Errorf("param not found"))
		return
	}

	id, err := utils.TransformStringToUInt64(idParam)
	if err != nil {
		HandleError(ctx, enum.BadRequest, fmt.Errorf("param not found"))
		return
	}

	err = h.svc.DeleteTask(id, userId64)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	HandleSuccess(ctx, map[string]interface{}{})
}

func (h *DailyHandler) ListTasks(ctx *gin.Context) {
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

	tasks, err := h.svc.ListTasks(userId64)

	if err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}
	result := make([]map[string]interface{}, 0)

	for _, task := range tasks {
		data := map[string]interface{}{
			"id":              utils.TransformUInt64ToString(task.ID),
			"title":           task.Title,
			"content":         task.Content,
			"start":           task.Start,
			"end":             task.End,
			"color":           task.BackgroundColor,
			"allDay":          task.AllDay,
			"backgroundColor": task.BackgroundColor,
			"textColor":       task.TextColor,
		}
		result = append(result, data)
	}

	HandleSuccess(ctx, result)
}

func (h *DailyHandler) DailyTask(ctx *gin.Context) {
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

	var payload *dto.DailyTaskDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandleError(ctx, enum.BadRequest, err)
		return
	}

	c := h.svc.DailyTask(userId64, payload)

	if c != "" {
		HandleError(ctx, enum.InternalServerError, fmt.Errorf(c))
		return
	}
	HandleSuccess(ctx, c)
}
