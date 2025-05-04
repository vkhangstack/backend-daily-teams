package model

import (
	"time"
)

type DailyTask struct {
	Title     string     `json:"title" db:"title"`
	Content   string     `json:"content" db:"content"`
	Published *time.Time `json:"published" db:"published"`
	UserId    uint64     `json:"userId" db:"user_id"`
	IsDaily   bool       `json:"isDaily" db:"is_daily"`
	*SqlModel
}

func (DailyTask) TableName() string {
	return "daily_tasks"
}
