package models

import "time"

type Profile struct {
	ID         int       `json:"id"`
	Phone      string    `json:"phone" gorm:"type: varchar(255)"`
	Image      string    `json:"image" gorm:"type: varchar(255)"`
	Address    string    `json:"address" gorm:"type: varchar(255)"`
	City       string    `json:"city" gorm:"type: varchar(255)"`
	PostalCode int       `json:"postal_code" gorm:"type: int"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
