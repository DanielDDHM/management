package models

import "time"

type Deposit struct {
	ID              int              `json:"id" gorm:"primaryKey"`
	Ref             string           `json:"ref,omitempty"`
	Currency        int              `json:"currency,omitempty"`
	Wallet          int              `json:"wallet,omitempty"`
	Status          string           `json:"status,omitempty"`
	ByUser          int              `json:"by_user,omitempty"`
	ByAdmin         int              `json:"by_admin,omitempty"`
	OriginData      string           `json:"origin_data,omitempty"`
	BankAccount     int              `json:"bank_account,omitempty"`
	Value           string           `json:"value,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	CurrencyData    *Currency        `json:"currency_data,omitempty" gorm:"foreignKey:Currency"`
	WalletData      *Wallet          `json:"wallet_data,omitempty" gorm:"foreignKey:Wallet"`
	BankAccountData *SiteBankAccount `json:"bank_account_data,omitempty" gorm:"foreignKey:BankAccount"`
}
