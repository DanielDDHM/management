package models

import "time"

type Transaction struct {
	ID            int        `json:"id" gorm:"primaryKey"`
	Owner         int        `json:"movement,omitempty"`
	Wallet        int        `json:"wallet,omitempty"`
	FromWallet    int        `json:"from_wallet,omitempty"`
	ToWallet      int        `json:"to_wallet,omitempty"`
	History       string     `json:"history,omitempty"`
	Args          string     `json:"args,omitempty"`
	AmountOrigin  string     `json:"amount_origin,omitempty"`
	Fees          string     `json:"fees,omitempty"`
	RelatedType   string     `json:"related_type,omitempty"`
	RelatedAction string     `json:"related_action,omitempty"`
	RelatedId     int        `json:"related_id,omitempty"`
	Comments      string     `json:"comments,omitempty"`
	Balance       string     `json:"balance,omitempty"`
	Schedule      int        `json:"schedule,omitempty"`
	Txid          string     `json:"txid,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	WalletData    *Wallet    `json:"wallet_data" gorm:"foreignKey:Wallet"`
}
