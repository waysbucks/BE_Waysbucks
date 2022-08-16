package models

import "time"

type Transaction struct {
	ID        int64       `json:"id"`
	UserID    int         `json:"user_id"`
	User      UserProfile `json:"user"`
	Status    string      `json:"status"`
	Total     int         `json:"total"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}
