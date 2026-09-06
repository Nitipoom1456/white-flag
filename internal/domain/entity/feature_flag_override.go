package entity

import (
	"time"

	"github.com/google/uuid"
)

type FeatureFlagOverride struct {
	ID                   uuid.UUID  `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()" json:"id"`
	FeatureFlagSettingID uuid.UUID  `gorm:"column:feature_flag_setting_id;type:uuid;not null" json:"feature_flag_setting_id"`
	SubjectValue         string     `gorm:"column:subject_value;type:varchar(256);not null" json:"subject_value"`
	Note                 *string    `gorm:"column:note;type:text" json:"note"`
	ExpiresAt            *time.Time `gorm:"column:expires_at;type:timestamptz" json:"expires_at"`
	CreatedAt            time.Time  `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;type:timestamptz;not null;default:now()" json:"updated_at"`

	FeatureFlagSetting *FeatureFlagSetting `gorm:"foreignKey:FeatureFlagSettingID;references:ID" json:"feature_flag_setting,omitempty"`
}

func (FeatureFlagOverride) TableName() string {
	return "feature_flag_override"
}
