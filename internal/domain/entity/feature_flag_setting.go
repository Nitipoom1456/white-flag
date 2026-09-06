package entity

import (
	"time"

	"github.com/google/uuid"
)

type FeatureFlagSetting struct {
	ID              uuid.UUID  `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()" json:"id"`
	FeatureFlagID   uuid.UUID  `gorm:"column:feature_flag_id;type:uuid;not null" json:"feature_flag_id"`
	EnvironmentID   uuid.UUID  `gorm:"column:environment_id;type:uuid;not null" json:"environment_id"`
	Active          bool       `gorm:"column:active;type:boolean;not null;default:false" json:"active"`
	MinAppVersionID *uuid.UUID `gorm:"column:min_app_version_id;type:uuid" json:"min_app_version_id"`
	MaxAppVersionID *uuid.UUID `gorm:"column:max_app_version_id;type:uuid" json:"max_app_version_id"`
	CreatedAt       time.Time  `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at;type:timestamptz;not null;default:now()" json:"updated_at"`

	FeatureFlag   *FeatureFlag `gorm:"foreignKey:FeatureFlagID;references:ID" json:"feature_flag,omitempty"`
	Environment   *Environment `gorm:"foreignKey:EnvironmentID;references:ID" json:"environment,omitempty"`
	MinAppVersion *AppVersion  `gorm:"foreignKey:MinAppVersionID;references:ID" json:"min_app_version,omitempty"`
	MaxAppVersion *AppVersion  `gorm:"foreignKey:MaxAppVersionID;references:ID" json:"max_app_version,omitempty"`
}

func (FeatureFlagSetting) TableName() string {
	return "feature_flag_setting"
}
