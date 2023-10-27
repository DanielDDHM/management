package repositories

import (
	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
)

type WalletRepository struct{}

func (s *WalletRepository) GetById(id int) (models.Wallet, error) {
	db := database.GetDatabase()

	var wallet models.Wallet
	err := db.First(&wallet, `id = ?`, id).Error

	if err != nil {
		return models.Wallet{}, err
	}

	return wallet, nil
}

func (s *WalletRepository) GetWallets(user_id int) ([]models.Wallet, error) {
	db := database.GetDatabase()

	var wallets []models.Wallet

	err := db.Find(&wallets, `owner = ?`, user_id).Error

	if err != nil {
		return []models.Wallet{}, err
	}

	return wallets, nil
}

func (s *WalletRepository) GetWalletBalance(wallet_id int) (string, error) {
	db := database.GetDatabase()

	var balance string
	var transaction models.Transaction

	err := db.Order(`id desc`).First(&transaction, `wallet = ?`, wallet_id).Error

	if err != nil {
		balance = "0"
	} else {
		balance = transaction.Balance
	}

	return balance, nil
}
