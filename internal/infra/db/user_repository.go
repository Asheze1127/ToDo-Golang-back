package db

import (
	"myapp-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct{
	Create(user *models.User) error
}

func userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB)*userRepository{
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error{
	return r.db.Create(user).Error
}
