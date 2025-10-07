package db

import (
	"myapp-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB)*GormUserRepository{
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// type UserRepository struct{
// 	Create(user *models.User) error
// }

// func userRepository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB)*userRepository{
// 	return &userRepository{db: db}
// }

// func (r *userRepository) Create(user *models.User) error{
// 	return r.db.Create(user).Error
// }
