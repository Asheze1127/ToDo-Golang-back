package service

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"myapp-backend/internal/infrastructure/db"
	"myapp-backend/internal/models"
)

type AuthService struct {
	UserRepo db.UserRepository
}

func NewAuthService(repo db.UserRepository)*AuthService{
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) SignUp(username, password string) (*models.User, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		ID: uuid.New(),
		Username: username,
		Password: string(hashedPassword),
	}
	if err := s.UserRepo.CreateUser(user); err != nil{
		return nil, err
	}

	return user, nil
}
