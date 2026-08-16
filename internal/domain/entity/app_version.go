package entity

import (
	"time"

	"github.com/google/uuid"
)

type AppVersion struct {
	ID         uuid.UUID  `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()" json:"id"`
	AppID      uuid.UUID  `gorm:"column:app_id;type:uuid;not null;index" json:"app_id"`
	Version    string     `gorm:"column:version;type:varchar(12);not null" json:"version"`
	ReleasedAt *time.Time `gorm:"column:released_at;type:timestamptz" json:"released_at"`
	CreatedAt  time.Time  `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;type:timestamptz;not null;default:now()" json:"updated_at"`

	App *App `gorm:"foreignKey:AppID;references:ID" json:"app,omitempty"`
}

func (AppVersion) TableName() string {
	return "app_version"
}
