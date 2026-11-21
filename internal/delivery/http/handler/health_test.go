package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Livez(t *testing.T) {
	h := NewHealthHandler("v1.0.0")
	req := httptest.NewRequest(http.MethodGet, "/livez", nil)
	rec := httptest.NewRecorder()

	h.Livez(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("kỳ vọng 200 OK nhưng nhận: %d", rec.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("lỗi parse json: %v", err)
	}

	if res["status"] != "UP" {
		t.Errorf("kỳ vọng status là UP nhưng nhận: %v", res["status"])
	}
}

func TestHealthHandler_Readyz(t *testing.T) {
	h := NewHealthHandler("v1.0.0")
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()

	h.Readyz(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("kỳ vọng 200 OK nhưng nhận: %d", rec.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("lỗi parse json: %v", err)
	}

	if res["status"] != "READY" || res["version"] != "v1.0.0" {
		t.Errorf("kết quả readiness không đúng: %+v", res)
	}
}
