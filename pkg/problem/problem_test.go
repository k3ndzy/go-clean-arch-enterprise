package problem

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProblemDetails_Write(t *testing.T) {
	rec := httptest.NewRecorder()
	prob := BadRequest("Email không hợp lệ", map[string]string{"email": "sai định dạng"})

	prob.Write(rec, "req-12345", "/api/v1/users")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("kỳ vọng HTTP status 400 nhưng nhận: %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/problem+json" {
		t.Errorf("kỳ vọng Content-Type là application/problem+json nhưng nhận: %s", rec.Header().Get("Content-Type"))
	}

	var res Details
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("lỗi parse json: %v", err)
	}

	if res.RequestID != "req-12345" || res.Instance != "/api/v1/users" || res.Code != "BAD_REQUEST" {
		t.Errorf("dữ liệu phản hồi sai: %+v", res)
	}
}
