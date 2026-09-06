package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Environment struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(64);not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index" json:"-"`
	AppID     uuid.UUID      `gorm:"column:app_id;type:uuid;not null" json:"app_id"`

	App *App `gorm:"foreignKey:AppID;references:ID" json:"app,omitempty"`
}

func (Environment) TableName() string {
	return "environment"
}
