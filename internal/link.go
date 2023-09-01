package internal

import "time"

type Link struct {
	ID        int64      `json:"-"`
	URL       string     `json:"url" validate:"required,url"`
	Slug      string     `json:"slug" gorm:"-"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
