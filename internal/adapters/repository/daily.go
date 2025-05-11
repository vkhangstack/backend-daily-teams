package repository

import (
	"fmt"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"time"
)

func (u *DB) CreateTask(dto *dto.CreateTaskDto, userId uint64) (*model.Task, error) {
	task := &model.Task{
		UserId:   userId,
		SqlModel: &model.SqlModel{ID: utils.GenerateID()},
	}

	task.Title = dto.Title
	task.Content = dto.Content
	task.Start = dto.Start
	task.End = dto.End
	task.AllDay = dto.AllDay

	if dto.BackgroundColor != "" {
		task.BackgroundColor = dto.BackgroundColor
	}
	if dto.TextColor != "" {
		task.TextColor = dto.TextColor
	}

	req := u.db.Create(&task)

	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("create daily task: %w", req.Error)
	}

	return task, nil
}

func (u *DB) UpdateTask(payload *dto.UpdateTaskDto, userId uint64) error {
	task := &model.Task{}
	id64, _ := utils.TransformStringToUInt64(payload.ID)
	req := u.db.First(&task, "id = ? and user_id = ?", id64, userId)
	if req.RowsAffected == 0 {
		return fmt.Errorf("task not found: %w", req.Error)
	}

	if payload.Content != "" {
		task.Content = payload.Content
	}
	if payload.Title != "" {
		task.Title = payload.Title
	}
	if payload.BackgroundColor != "" {
		task.BackgroundColor = payload.BackgroundColor
	}
	if payload.TextColor != "" {
		task.TextColor = payload.TextColor
	}
	if payload.Start != "" {
		task.Start = payload.Start
	}
	if payload.End != "" {
		task.End = payload.End
	}
	if payload.AllDay != nil {
		task.AllDay = payload.AllDay
	}

	task.UpdatedAt = time.Now()

	req = u.db.Model(&task).Where("id = ?", task.ID).Where("user_id = ?", userId).Updates(&task)
	if req.RowsAffected == 0 {
		return fmt.Errorf("update daily task error: %w", req.Error)
	}

	return nil
}

func (u *DB) DeleteTask(id uint64, userId uint64) error {
	task := &model.Task{}
	req := u.db.Where("id = ?", id).Where("user_id = ?", userId).Delete(&task)
	if req.RowsAffected == 0 {
		return fmt.Errorf("delete daily task error: %w", req.Error)
	}
	return nil
}

func (u *DB) ListTasks(userId uint64) ([]*model.Task, error) {

	var tasks []*model.Task

	req := u.db.Find(&tasks, "user_id = ?", userId)
	if req.Error != nil {
		return nil, fmt.Errorf("list daily task error: %w", req.Error)
	}
	return tasks, nil
}
func (u *DB) CheckDaily(userId uint64) bool {
	daily := &model.Daily{}
	req := u.db.Where("user_id = ?", userId).Where("CAST(created_at AS date) = ?", time.Now().UTC().Format("2006-01-02")).Select("id").First(&daily)

	if req.RowsAffected == 0 {
		return false
	}

	return true
}

func (u *DB) CreateDaily(userId uint64) error {
	daily := &model.Daily{}

	daily.ID = utils.GenerateID()
	daily.UserID = userId
	daily.CreatedAt = time.Now()
	daily.UpdatedAt = time.Now()
	req := u.db.Create(&daily)

	if req.RowsAffected == 0 {
		return fmt.Errorf("create daily daily error: %w", req.Error)
	}
	return nil
}
