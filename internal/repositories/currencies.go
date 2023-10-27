package repositories

import (
	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
)

type CurrencyRepository struct{}

func (s *CurrencyRepository) GetById(id int) (models.Wallet, error) {
	db := database.GetDatabase()

	var wallet models.Wallet
	err := db.First(&wallet, `id = ?`, id).Error

	if err != nil {
		return models.Wallet{}, err
	}

	return wallet, nil
}
