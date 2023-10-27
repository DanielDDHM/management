package repositories

import (
	"errors"
	"time"

	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
	"github.com/DanielDDHM/management/pkg/utils"
)

type TransactionRepository struct{}

func (s *TransactionRepository) GetAll(user_id int, page int, size int,
	sortOrder string, sortOrderName string, startDate time.Time, endDate time.Time) (utils.PaginateResponse, error) {

	if (!startDate.IsZero() && !startDate.IsZero()) && startDate.After(endDate) {
		return utils.PaginateResponse{}, errors.New("Range Date Invalid")
	}

	db := database.GetDatabase()

	var transactions []models.Transaction

	var walletsIds []int

	var wallets []models.Wallet

	err := db.Find(&wallets, `owner = ?`, user_id).Error

	if err != nil {
		return utils.PaginateResponse{}, nil
	}

	for _, w := range wallets {
		walletsIds = append(walletsIds, w.ID)
	}

	if endDate.IsZero() {
		endDate = time.Now()
	}

	err = db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("WalletData").Order(
		sortOrderName+" "+sortOrder).Limit(size).Offset((page-1)*size).Find(&transactions,
		`wallet IN ?`, walletsIds).Error

	if err != nil {
		return utils.PaginateResponse{}, err
	}

	var count int64
	db.Model(&models.Transaction{}).Where("created_at BETWEEN ? AND ?", startDate, endDate).Where("wallet IN ?", walletsIds).Count(&count)

	return utils.PaginateResponse{
		Data:       transactions,
		TotalPages: ((int(count)) / size) + 1,
		TotalCount: int(count),
	}, nil
}

func (s *TransactionRepository) GetBalance(wallet_id int) (string, error) {
	db := database.GetDatabase()

	var transaction models.Transaction
	err := db.Order(`id desc`).First(&transaction, `id ? =`, wallet_id).Error

	if err != nil {
		return "", err
	}

	return transaction.Balance, nil
}
