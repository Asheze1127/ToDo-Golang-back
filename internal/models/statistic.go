package models 

import (
	"time"
	"github.com/google/uuid"
)

type Statistic struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserId    uuid.UUID `gorm:"not null" json:"user_id"`
	TotalTasks int       `gorm:"not null" json:"total_tasks"`
	CompletedTasks int       `gorm:"not null" json:"completed_tasks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}