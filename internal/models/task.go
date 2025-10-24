package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityNormal TaskPriority = "normal"
	TaskPriorityHigh   TaskPriority = "high"
)

type Task struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"` // 自動採番されるタスク識別子
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`                  // このタスクを所有するユーザーID
	User   User      `gorm:"foreignKey:UserID" json:"user,omitempty"`                  // タスク所有者の詳細情報（Preload 用）

	Title       string       `gorm:"size:140;not null" json:"title"`                      // タスクのタイトル
	Description string       `gorm:"type:text" json:"description"`                        // タスクの詳細説明
	Status      TaskStatus   `gorm:"type:task_status;default:'todo'" json:"status"`       // 現在の状態（todo / in_progress / done）
	Priority    TaskPriority `gorm:"type:task_priority;default:'normal'" json:"priority"` // 優先度（low / normal / high）
	ExtraInfo   string       `gorm:"size:255" json:"extra_info"`                          // 補足情報やラベルなどの自由入力欄

	PlannedStartAt *time.Time `json:"planned_start_at"` // 計画開始日時（未設定可）
	PlannedDueAt   *time.Time `json:"planned_due_at"`   // 期限となる予定日時（未設定可）
	CompletedAt    *time.Time `json:"completed_at"`     // 完了した日時（未完了なら null）
	PinnedAt       *time.Time `json:"pinned_at"`        // ピン留めした日時（未設定なら null）

	CreatedAt time.Time      `json:"created_at"`     // レコード作成日時
	UpdatedAt time.Time      `json:"updated_at"`     // レコード更新日時
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // ソフトデリート用の削除日時
}
