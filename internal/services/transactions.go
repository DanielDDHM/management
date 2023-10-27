package services

import (
	"time"

	"github.com/DanielDDHM/management/internal/interfaces/repositories"
	"github.com/DanielDDHM/management/pkg/utils"
)

type TransactionService struct {
	repository repositories.ITransactionRepository
}

func (r *TransactionService) GetAll(user_id int, page int, size int,
	sortOrder string, sortOrderName string, startDate time.Time, endDate time.Time) (utils.PaginateResponse, error) {

	return r.repository.GetAll(user_id, page, size, sortOrder, sortOrderName, startDate, endDate)
}

func NewTransactionService(r repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{
		repository: r,
	}
}
