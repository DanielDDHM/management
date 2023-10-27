package repositories

import (
	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
)

type SessionRepository struct{}

func (s *SessionRepository) GetAllByUser(userId int) ([]models.UsersSession, error) {
	db := database.GetDatabase()

	var sessions []models.UsersSession

	err := db.Find(&sessions, "user = ?", userId).Error

	if err != nil {
		return []models.UsersSession{}, err
	}

	return sessions, nil
}
