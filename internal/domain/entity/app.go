package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type App struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()"`
	Name      string         `gorm:"column:name;type:varchar(120);not null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`
}

func (App) TableName() string {
	return "app"
}
