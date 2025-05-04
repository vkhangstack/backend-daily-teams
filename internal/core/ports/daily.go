package ports

import (
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
)

type DailyService interface {
	CreateTask(payload *dto.CreateDailyDto, userId uint64) (*model.DailyTask, error)
	UpdateTask(payload *dto.UpdateDailyDto, userId uint64) error
	DeleteTask(id uint64, userId uint64) error
	ListTasks(userId uint64) ([]*model.DailyTask, error)
}

type DailyRepository interface {
	CreateTask(payload *dto.CreateDailyDto, userId uint64) (*model.DailyTask, error)
	UpdateTask(payload *dto.UpdateDailyDto, userId uint64) error
	DeleteTask(id uint64, userId uint64) error
	ListTasks(userId uint64) ([]*model.DailyTask, error)
}
