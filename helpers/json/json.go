package jsonHelpers

import (
	"encoding/json"
	"net/http"

	loggingHelper "test-fullstack/helpers/log"
	paginationHelper "test-fullstack/helpers/pagination"
)

// Return JSON
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// Return JSON Success with HTTP Code
func SuccessResponse(w http.ResponseWriter, r *http.Request, status bool, httpStatus int, code string, data interface{}, message string, pagination *paginationHelper.Page) {
	meta := &meta{
		Status:     status,
		Message:    message,
		Code:       code,
		Pagination: pagination,
	}

	res := &response{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "SUCCESS", data)

	WriteJSON(w, httpStatus, res)
}

// Return JSON Error
func ErrorResponse(w http.ResponseWriter, r *http.Request, status bool, httpCode int, message string, data interface{}) {

	meta := &meta{
		Code:    string(httpCode),
		Status:  status,
		Message: message,
	}

	res := &response{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "ERROR", res)

	WriteJSON(w, httpCode, res)
}
