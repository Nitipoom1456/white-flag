package dto

type CreateAppRequest struct {
	Name string `json:"name" binding:"required"`
}

type AppResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
