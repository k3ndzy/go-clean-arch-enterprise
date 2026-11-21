package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func TestLogger_SecretRedaction(t *testing.T) {
	buf := new(bytes.Buffer)
	log := New(buf, slog.LevelInfo)

	log.InfoContext(context.Background(), "user registration attempt",
		slog.String("email", "dev@example.com"),
		slog.String("password", "super-secret-password-123"),
		slog.String("apiKey", "api-key-998877"),
		slog.String("Authorization", "Bearer eyJhbGciOi..."),
	)

	var logData map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logData); err != nil {
		t.Fatalf("lỗi parse json log: %v", err)
	}

	if logData["password"] != "[REDACTED]" {
		t.Errorf("kỳ vọng password bị redact nhưng nhận: %v", logData["password"])
	}

	if logData["apiKey"] != "[REDACTED]" {
		t.Errorf("kỳ vọng apiKey bị redact nhưng nhận: %v", logData["apiKey"])
	}

	if logData["email"] != "dev@example.com" {
		t.Errorf("email không nên bị redact: %v", logData["email"])
	}

	rawLog := buf.String()
	if strings.Contains(rawLog, "super-secret-password-123") {
		t.Error("chuỗi mật khẩu thật xuất hiện trong log!")
	}
}
