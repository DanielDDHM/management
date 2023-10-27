package services

import (
	"github.com/DanielDDHM/management/internal/interfaces/repositories"
	"github.com/DanielDDHM/management/internal/models"
)

type SessionService struct {
	repository repositories.ISessionRepository
}

func (s *SessionService) ListSessionsByUser(userId int) ([]models.UsersSession, error) {
	return s.repository.GetAllByUser(userId)
}

func NewSessionService(r repositories.ISessionRepository) *SessionService {
	return &SessionService{
		repository: r,
	}
}
