package dto

type CreateEnvironmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type EnvironmentResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	AppID string `json:"app_id"`
}
