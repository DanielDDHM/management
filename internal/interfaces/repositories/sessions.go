package repositories

import "github.com/DanielDDHM/management/internal/models"

type ISessionRepository interface {
	GetAllByUser(userId int) ([]models.UsersSession, error)
}
