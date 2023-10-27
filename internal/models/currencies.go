package models

import "time"

type Currency struct {
	ID                     int        `json:"id" gorm:"primaryKey"`
	Type                   string     `json:"type,omitempty"`
	Precision              int        `json:"precision,omitempty"`
	Name                   string     `json:"name,omitempty"`
	IsoCode                string     `json:"iso_code,omitempty"`
	Symbol                 string     `json:"symbol,omitempty"`
	Separator              string     `json:"separator,omitempty"`
	Decimal                string     `json:"decimal,omitempty"`
	IsPivot                int        `json:"is_pivot,omitempty"`
	PriceInPivot           string     `json:"price_in_pivot,omitempty"`
	BuyRatio               float64    `json:"buy_ratio,omitempty"`
	SellRatio              float64    `json:"sell_ratio,omitempty"`
	Reservation            string     `json:"reservation,omitempty"`
	SysWallet              int        `json:"sys_wallet,omitempty"`
	Enabled                int        `json:"enabled,omitempty"`
	SiteBankAccount        int        `json:"site_bank_account,omitempty"`
	SpredFactor            float64    `json:"spread_factor,omitempty"`
	DefaultInterest        float64    `json:"default_interest,omitempty"`
	DefaultWeekendInterest float64    `json:"default_weekend_interest,omitempty"`
	Deposits               []*Deposit `json:"deposits,omitempty" gorm:"foreignkey:BankAccount"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}
