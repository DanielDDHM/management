package repositories

import "github.com/DanielDDHM/management/internal/models"

type IWalletRepository interface {
	GetById(id int) (models.Wallet, error)
	GetWallets(user_id int) ([]models.Wallet, error)
	GetWalletBalance(wallet_id int) (string, error)
}
