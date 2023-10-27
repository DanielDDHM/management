package models

import "time"

type UsersSession struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Active     int        `json:"active,omitempty"`
	SessionId  int        `json:"sessid,omitempty"`
	User       int        `json:"user,omitempty"`
	Ua         string     `json:"ua,omitempty"`
	Ip         string     `json:"ip,omitempty"`
	Device     string     `json:"device,omitempty"`
	Previleges string     `json:"privileges,omitempty"`
	Type       string     `json:"type,omitempty"`
	ApiApp     int        `json:"api_app,omitempty"`
	UserData   *User      `json:"user_data,omitempty" gorm:"foreignKey:User"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
