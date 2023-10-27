package models

import "time"

type Wallet struct {
	ID                int            `json:"id" gorm:"primaryKey"`
	Owner             int            `json:"owner,omitempty"`
	WalletName        string         `json:"wallet_name,omitempty"`
	WalletCurrency    int            `json:"wallet_currency,omitempty"`
	IsPublic          int            `json:"is_public,omitempty"`
	IsOwnerInfoPublic int            `json:"is_owner_info_public,omitempty"`
	Enabled           int            `json:"enabled,omitempty"`
	Ordering          int            `json:"ordering,omitempty"`
	WalletAddress     string         `json:"wallet_address,omitempty"`
	IsFavorite        int            `json:"is_favorite,omitempty"`
	Balance           string         `json:"balance,omitempty"`
	CreatedAt         *time.Time     `json:"created_at,omitempty"`
	UpdatedAt         *time.Time     `json:"updated_at,omitempty"`
	Transactions      []*Transaction `json:"transactions,omitempty" gorm:"foreignkey:Wallet"`
	Deposits          []*Deposit     `json:"deposits,omitempty" gorm:"foreignkey:Wallet"`
}
