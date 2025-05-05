package dto

type CreateDailyDto struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	Start           string `json:"start"`
	End             string `json:"end"`
	TextColor       string `json:"textColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	IsDaily         bool   `json:"isDaily"`
}

type UpdateDailyDto struct {
	ID      uint64 `json:"id"`
	Content string `json:"content,omitempty"`
	Title   string `json:"title,omitempty"`
}
