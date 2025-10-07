package service

import (
	"github.com/google/uuid"
	"myapp-backend/internal/infra/db"
	"myapp-backend/internal/models"
	"myapp-backend/internal/util"
)

type AuthService struct {
	UserRepo db.UserRepository
}

func NewAuthService(repo db.UserRepository)*AuthService{
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) SignUp(username, password string) (string, error){
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &models.User{
		ID: uuid.New(),
		Username: username,
		Password: hashedPassword,
	}
	if err := s.UserRepo.Create(user); err != nil{
		return "", err
	}

	token, err := util.GenerateJWT(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
