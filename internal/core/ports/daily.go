package ports

import (
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
)

type DailyService interface {
	CreateTask(payload *dto.CreateTaskDto, userId uint64) (*model.Task, error)
	UpdateTask(payload *dto.UpdateTaskDto, userId uint64) error
	DeleteTask(id uint64, userId uint64) error
	ListTasks(userId uint64) ([]*model.Task, error)
	DailyTask(userId uint64) string
}

type DailyRepository interface {
	CreateTask(payload *dto.CreateTaskDto, userId uint64) (*model.Task, error)
	UpdateTask(payload *dto.UpdateTaskDto, userId uint64) error
	DeleteTask(id uint64, userId uint64) error
	ListTasks(userId uint64) ([]*model.Task, error)
	CheckDaily(userId uint64) bool
	CreateDaily(userId uint64) error
}
