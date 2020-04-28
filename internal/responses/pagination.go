package responses

import (
	"reflect"
)

type paginationMeta struct {
	Page       int `json:"current_page"`
	PerPage    int `json:"per_page"`
	Count      int `json:"count"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type pagination struct {
	Data interface{}     `json:"data"`
	Meta *paginationMeta `json:"meta"`
}

func NewPagination(data interface{}) *pagination {
	return &pagination{
		Data: data,
		Meta: &paginationMeta{
			Count: reflect.ValueOf(data).Len(),
			Page:  1,
		},
	}
}
