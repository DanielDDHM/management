package services

import (
	"github.com/DanielDDHM/management/internal/interfaces/repositories"
	"github.com/DanielDDHM/management/internal/models"
)

type WalletService struct {
	repository repositories.IWalletRepository
}

func (s *WalletService) GetWalletById(id int) (models.Wallet, error) {
	return s.repository.GetById(id)
}

func (s *WalletService) GetWallets(user_id int) ([]models.Wallet, error) {
	wallets, err := s.repository.GetWallets(user_id)

	if err != nil {
		return []models.Wallet{}, nil
	}

	for i, w := range wallets {
		balance, _ := s.repository.GetWalletBalance(w.ID)
		wallets[i].Balance = balance
	}

	return wallets, nil
}

func (s *WalletService) getWalletBallance(wallet_id int) (string, error) {

	_, err := s.GetWalletById(wallet_id)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	// if as_currency {
	// 	var currency models.Currency

	// 	db.First(&currency, )
	// }
	return "", nil
}

func NewWalletService(r repositories.IWalletRepository) *WalletService {
	return &WalletService{
		repository: r,
	}
}
