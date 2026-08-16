package response

type SuccessResponse struct {
	Message string `json:"message"`
}

type SuccessDataResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PaginationResponse[T any] struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
	Data     []T `json:"data"`
}
