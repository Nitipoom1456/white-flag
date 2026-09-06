package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FeatureFlag struct {
	ID          uuid.UUID                   `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()" json:"id"`
	AppID       uuid.UUID                   `gorm:"column:app_id;type:uuid;not null" json:"app_id"`
	Name        string                      `gorm:"column:name;type:varchar(256);not null" json:"name"`
	Description *string                     `gorm:"column:description;type:text" json:"description"`
	Tags        datatypes.JSONSlice[string] `gorm:"column:tags;type:jsonb" json:"tags"`
	CreatedAt   time.Time                   `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time                   `gorm:"column:updated_at;type:timestamptz;not null;default:now()" json:"updated_at"`
	DeletedAt   gorm.DeletedAt              `gorm:"column:deleted_at;type:timestamptz;index" json:"-"`

	App *App `gorm:"foreignKey:AppID;references:ID" json:"app,omitempty"`
}

func (FeatureFlag) TableName() string {
	return "feature_flag"
}
