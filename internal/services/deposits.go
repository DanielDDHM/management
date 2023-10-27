package services

import (
	"github.com/DanielDDHM/management/internal/interfaces/repositories"
	"github.com/DanielDDHM/management/pkg/utils"
)

type DepositService struct {
	repository repositories.IDepositRepository
}

func (s *DepositService) ListDeposits(userId int, curr_type string, page int, size int) (utils.PaginateResponse, error) {
	return s.repository.GetAll(userId, curr_type, page, size)
}

func NewDepositService(r repositories.IDepositRepository) *DepositService {
	return &DepositService{
		repository: r,
	}
}
