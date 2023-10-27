package utils

type PaginateResponse struct {
	Data       interface{} `json:"data,omitempty"`
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
}
