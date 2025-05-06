package model

import (
	"time"
)

type DailyTask struct {
	Title           string     `json:"title" db:"title"`
	Content         string     `json:"content" db:"content"`
	Start           string     `json:"start" db:"start"`
	End             string     `json:"end" db:"end"`
	TextColor       string     `json:"textColor" db:"text_color"`
	BackgroundColor string     `json:"backgroundColor" db:"background_color"`
	Published       *time.Time `json:"published" db:"published"`
	UserId          uint64     `json:"userId" db:"user_id"`
	IsDaily         bool       `json:"isDaily" db:"is_daily"`
	AllDay          bool       `json:"allDay" db:"all_day"`
	*SqlModel
}

func (DailyTask) TableName() string {
	return "daily_tasks"
}
