package models

import "time"

type User struct {
	ID                  uint            `json:"id" gorm:"primaryKey"`
	Password            string          `json:"password,omitempty"`
	Pin                 string          `json:"pin,omitempty"`
	Type                string          `json:"type,omitempty"`
	Enabled             uint            `json:"enabled,omitempty"`
	OtpSecret           string          `json:"otp_secret,omitempty"`
	Status              uint            `json:"status,omitempty"`
	NextAccountBirthday string          `json:"next_account_birthday"`
	Sessions            []*UsersSession `json:"sessions,omitempty" gorm:"foreignkey:User"`
	CreatedAt           *time.Time      `json:"created_at"`
	UpdatedAt           *time.Time      `json:"updated_at,omitempty"`
}
