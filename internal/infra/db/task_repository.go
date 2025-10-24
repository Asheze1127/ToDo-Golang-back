package db

import (
	"time"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"myapp-backend/internal/models"
)

type TaskRepository interface {
	Create(task *models.Task) error
	FindRecent() ([]models.Task, error)
	FindByUserID(userId uuid.UUID) ([]models.Task, error)
	DeleteByID(id uuid.UUID, userId uuid.UUID) error
	UpdateByID(id uuid.UUID, userId uuid.UUID, task models.Task) error
}

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository{
	return &GormTaskRepository{db:db}
}

//POST /tasks
func (r *GormTaskRepository) Create(task *models.Task) error{
	return r.db.Create(task).Error
}

//GET /tasks
func (r *GormTaskRepository) FindRecent() ([]models.Task, error){//直近の50件のタスクを取得
	var tasks []models.Task
	err := r.db.Where("created_at <= ?", time.Now().Add(-24 * time.Hour)).Find(&tasks).Error
	if err != nil{
		return nil, err
	}
	return tasks, nil
}

//GET /tasks/{userId}
func (r *GormTaskRepository) FindByUserID(userId uuid.UUID) ([]models.Task, error){//ユーザーごとのタスクを取得
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&tasks).Error
	if err != nil{
		return nil, err
	}
	return tasks, nil
}

//DELETE /tasks/{id}
func (r *GormTaskRepository) DeleteByID(id uuid.UUID, userId uuid.UUID) error{
	err := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Task{}).Error
	if err != nil{
		return err
	}
	return nil
}

//PUT /tasks/{id}
func (r *GormTaskRepository) UpdateByID(id uuid.UUID, userId uuid.UUID, task models.Task) error{
	err := r.db.Model(&models.Task{}).Where("id = ? AND user_id = ?", id, userId).Updates(task).Error
	if err != nil{
		return err
	}
	return nil	
}
