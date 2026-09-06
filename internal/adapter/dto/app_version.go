package dto

type CreateAppVersionRequest struct {
	Version string `json:"version"`
}

type AppVersionResponse struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	AppID   string `json:"app_id"`
}
