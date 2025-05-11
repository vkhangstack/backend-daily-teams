package model

type Daily struct {
	UserID uint64 `json:"userId" db:"user_id"`
	*SqlModel
}

func (Daily) TableName() string {
	return "daily"
}
