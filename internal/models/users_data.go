package models

import "time"

type UsersData struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	User      uint       `json:"user,omitempty"`
	Key       string     `json:"key,omitempty"`
	Value     string     `json:"value,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
