package problem

import (
	"encoding/json"
	"net/http"
)

// Details tuân thủ tiêu chuẩn RFC 7807 (Problem Details for HTTP APIs)
type Details struct {
	Type      string            `json:"type"`
	Title     string            `json:"title"`
	Status    int               `json:"status"`
	Detail    string            `json:"detail"`
	Instance  string            `json:"instance,omitempty"`
	Code      string            `json:"code,omitempty"`
	RequestID string            `json:"request_id,omitempty"`
	Errors    map[string]string `json:"errors,omitempty"`
}

func New(status int, title, detail, code string) *Details {
	return &Details{
		Type:   "about:blank",
		Title:  title,
		Status: status,
		Detail: detail,
		Code:   code,
	}
}

func NotFound(detail string) *Details {
	return New(http.StatusNotFound, "Tài nguyên không tìm thấy", detail, "NOT_FOUND")
}

func BadRequest(detail string, fieldErrors map[string]string) *Details {
	d := New(http.StatusBadRequest, "Yêu cầu không hợp lệ", detail, "BAD_REQUEST")
	d.Errors = fieldErrors
	return d
}

func Unauthorized(detail string) *Details {
	return New(http.StatusUnauthorized, "Chưa xác thực danh tính", detail, "UNAUTHORIZED")
}

func Forbidden(detail string) *Details {
	return New(http.StatusForbidden, "Không có quyền truy cập", detail, "FORBIDDEN")
}

func InternalServerError(detail string) *Details {
	return New(http.StatusInternalServerError, "Lỗi máy chủ nội bộ", detail, "INTERNAL_SERVER_ERROR")
}

func (d *Details) Write(w http.ResponseWriter, requestID, instance string) {
	d.RequestID = requestID
	d.Instance = instance

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(d.Status)
	_ = json.NewEncoder(w).Encode(d)
}
