package services

import (
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"github.com/vkhangstack/dlt/internal/core/ports"
)

type DailyService struct {
	repo ports.DailyRepository
}

func NewDailyService(repo ports.DailyRepository) *DailyService {
	return &DailyService{repo: repo}
}

func (d *DailyService) CreateTask(payload *dto.CreateDailyDto, userId uint64) (*model.DailyTask, error) {
	return d.repo.CreateTask(payload, userId)
}

func (d *DailyService) UpdateTask(payload *dto.UpdateDailyDto, userId uint64) error {
	return d.repo.UpdateTask(payload, userId)
}
func (d *DailyService) DeleteTask(id uint64, userId uint64) error {
	return d.repo.DeleteTask(id, userId)
}
func (d *DailyService) ListTasks(userId uint64) ([]*model.DailyTask, error) {
	return d.repo.ListTasks(userId)
}
