package dto

type CreateAppVersionRequest struct {
	AppID   string `json:"app_id"`
	Version string `json:"version"`
}
