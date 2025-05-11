package dto

type CreateTaskDto struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	Start           string `json:"start"`
	End             string `json:"end"`
	TextColor       string `json:"textColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	AllDay          *bool  `json:"allDay" default:"false"`
}

type UpdateTaskDto struct {
	ID              string `json:"id"`
	Content         string `json:"content,omitempty"`
	Title           string `json:"title,omitempty"`
	Start           string `json:"start,omitempty"`
	End             string `json:"end,omitempty"`
	TextColor       string `json:"textColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	AllDay          *bool  `json:"allDay,omitempty"`
}

type DailyTaskDto struct {
	AccessToken string `json:"accessToken"`
	Content     string `json:"content"`
}
