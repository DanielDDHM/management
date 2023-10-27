package repositories

import (
	"fmt"

	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
	"github.com/DanielDDHM/management/pkg/utils"
)

type DepositRepository struct{}

func (s *DepositRepository) GetAll(userId int, curr_type string, page int, size int) (utils.PaginateResponse, error) {
	db := database.GetDatabase()

	var deposits []*models.Deposit
	var wallets []models.Wallet
	var currencies []models.Currency

	if curr_type == "" {
		curr_type = "crypto"
	}

	err := db.Find(&currencies, "type = ?", curr_type).Error

	fmt.Println("Currencies", currencies)
	if err != nil {
		return utils.PaginateResponse{}, err
	}

	var currenciesIds []int

	for _, c := range currencies {
		currenciesIds = append(currenciesIds, c.ID)
	}

	err = db.Find(&wallets, `owner = ? AND wallet_currency IN ?`, userId, currenciesIds).Error

	if err != nil {
		return utils.PaginateResponse{}, err
	}
	var walletIds []int
	for _, w := range wallets {
		walletIds = append(walletIds, w.ID)
	}

	db.Preload("WalletData").Limit(size).Offset((page-1)*size).Find(&deposits, "wallet IN ?", walletIds)

	var count int64

	db.Model(&models.Deposit{}).Where("wallet IN ?", walletIds).Count(&count)

	return utils.PaginateResponse{
		TotalCount: int(count),
		TotalPages: ((int(count)) / size) + 1,
		Data:       deposits,
	}, nil
}
