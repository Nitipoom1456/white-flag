package dto

type CreateFeatureFlagRequest struct {
	AppID         string   `json:"app_id" binding:"required,uuid"`
	Name          string   `json:"name" binding:"required"`
	Description   *string  `json:"description"`
	Active        bool     `json:"active"`
	Tags          []string `json:"tags"`
	MinAppVersion *string  `json:"min_app_version" binding:"omitempty,uuid"`
	MaxAppVersion *string  `json:"max_app_version" binding:"omitempty,uuid"`
}

type UpdateFeatureFlagRequest struct {
	Name          string   `json:"name"`
	Description   *string  `json:"description"`
	Active        bool     `json:"active"`
	Tags          []string `json:"tags"`
	MinAppVersion *string  `json:"min_app_version" binding:"omitempty,uuid"`
	MaxAppVersion *string  `json:"max_app_version" binding:"omitempty,uuid"`
}

type FeatureFlagResponse struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   *string  `json:"description"`
	Active        bool     `json:"active"`
	Tags          []string `json:"tags"`
	MinAppVersion *string  `json:"min_app_version" binding:"omitempty"`
	MaxAppVersion *string  `json:"max_app_version" binding:"omitempty"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type GetAppFeatureFlagByEnvironmentIDQuery struct {
	Search string `form:"search"`
}
