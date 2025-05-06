package dto

type CreateDailyDto struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	Start           string `json:"start"`
	End             string `json:"end"`
	TextColor       string `json:"textColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	IsDaily         bool   `json:"isDaily,omitempty"`
	AllDay          bool   `json:"allDay" default:"false"`
}

type UpdateDailyDto struct {
	ID              uint64 `json:"id"`
	Content         string `json:"content,omitempty"`
	Title           string `json:"title,omitempty"`
	Start           string `json:"start,omitempty"`
	End             string `json:"end,omitempty"`
	TextColor       string `json:"textColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	IsDaily         bool   `json:"isDaily,omitempty"`
	AllDay          bool   `json:"allDay,omitempty"`
}
