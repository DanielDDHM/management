package repositories

import (
	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/internal/models"
)

type UserRepository struct{}

func (s *UserRepository) GetById(id int) (models.User, error) {

	db := database.GetDatabase()

	var user models.User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	var datas models.UsersData

	err = db.Find(&datas, "user = ?", user.ID).Error

	return models.User{}, nil
}
