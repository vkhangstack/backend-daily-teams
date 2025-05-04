package dto

type CreateDailyDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	IsDaily bool   `json:"isDaily"`
}

type UpdateDailyDto struct {
	ID      uint64 `json:"id"`
	Content string `json:"content,omitempty"`
	Title   string `json:"title,omitempty"`
}
