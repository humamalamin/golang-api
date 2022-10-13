package jsonHelpers

import paginationHelper "test-fullstack/helpers/pagination"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type meta struct {
	Status     bool                   `json:"-"`
	Code       string                 `json:"code"`
	Message    interface{}            `json:"message"`
	Pagination *paginationHelper.Page `json:"pagination,omitempty"`
}
