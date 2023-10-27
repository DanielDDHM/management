package repositories

import (
	"time"

	"github.com/DanielDDHM/management/pkg/utils"
)

type ITransactionRepository interface {
	GetAll(wallet_id int, page int, size int, sortOrder string, sortOrderName string,
		startDate time.Time, endDate time.Time) (utils.PaginateResponse, error)
	GetBalance(wallet_id int) (string, error)
}
