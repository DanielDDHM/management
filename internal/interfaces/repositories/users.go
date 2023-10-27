package repositories

import "github.com/DanielDDHM/management/internal/models"

type IUserRepository interface {
	GetById(id int) (models.User, error)
}
