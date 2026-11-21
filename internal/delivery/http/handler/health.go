package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct {
	startTime time.Time
	version   string
}

func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		version:   version,
	}
}

// RegisterRoutes gán các route liveness & readiness
func (h *HealthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/livez", h.Livez)
	mux.HandleFunc("/readyz", h.Readyz)
	mux.HandleFunc("/health", h.Livez)
}

// Livez báo cáo ứng dụng đang sống (Liveness Probe của Kubernetes)
func (h *HealthHandler) Livez(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "UP",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Readyz báo cáo ứng dụng sẵn sàng nhận traffic (Readiness Probe của Kubernetes)
func (h *HealthHandler) Readyz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "READY",
		"uptime":    time.Since(h.startTime).String(),
		"version":   h.version,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
