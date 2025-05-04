package model

import "time"

type SqlModel struct {
	ID        uint64     `gorm:"primary_key" json:"id" db:"id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" db:"deleted_at"`
}

type UpdateModel struct {
	UpdatedBy uint64     `json:"userId" db:"user_id"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateModel struct {
	CreatedAt *time.Time `json:"createdAt,omitempty" db:"created_at"`
	CreatedBy uint64     `json:"createdBy,omitempty" db:"created_by"`
}
