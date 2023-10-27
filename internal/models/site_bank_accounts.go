package models

import "time"

type SiteBankAccount struct {
	ID             int        `json:"id" gorm:"primaryKey"`
	BankName       string     `json:"bank_name,omitempty"`
	BankData       string     `json:"bank_data,omitempty"`
	BankDataRaw    string     `json:"bank_data_raw,omitempty"`
	DepositMethods string     `json:"deposit_methods,omitempty"`
	Currencies     string     `json:"currencies,omitempty"`
	Enabled        int        `json:"enabled,omitempty"`
	Deposits       []*Deposit `json:"deposits,omitempty" gorm:"foreignkey:Currency"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
