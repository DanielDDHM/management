package repositories

import "github.com/DanielDDHM/management/pkg/utils"

type IDepositRepository interface {
	GetAll(userId int, curr_type string, page int, size int) (utils.PaginateResponse, error)
}
