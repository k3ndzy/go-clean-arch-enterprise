package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

var sensitivePattern = regexp.MustCompile(`(?i)(password|passwd|secret|token|authorization|auth|apiKey|private_key)`)

// RedactionHandler tự động che các trường nhạy cảm trước khi ghi log
type RedactionHandler struct {
	slog.Handler
}

func NewRedactionHandler(h slog.Handler) *RedactionHandler {
	return &RedactionHandler{Handler: h}
}

func (h *RedactionHandler) Handle(ctx context.Context, r slog.Record) error {
	var cleanAttrs []slog.Attr

	r.Attrs(func(a slog.Attr) bool {
		cleanAttrs = append(cleanAttrs, sanitizeAttr(a))
		return true
	})

	newRecord := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	newRecord.AddAttrs(cleanAttrs...)

	return h.Handler.Handle(ctx, newRecord)
}

func sanitizeAttr(a slog.Attr) slog.Attr {
	if sensitivePattern.MatchString(a.Key) {
		return slog.String(a.Key, "[REDACTED]")
	}

	if a.Value.Kind() == slog.KindGroup {
		var subAttrs []slog.Attr
		for _, sub := range a.Value.Group() {
			subAttrs = append(subAttrs, sanitizeAttr(sub))
		}
		return slog.Group(a.Key, convertToAny(subAttrs)...)
	}

	// Nếu giá trị chuỗi chứa token hoặc bearer
	if a.Value.Kind() == slog.KindString {
		strVal := a.Value.String()
		if strings.HasPrefix(strings.ToLower(strVal), "bearer ") {
			return slog.String(a.Key, "Bearer [REDACTED]")
		}
	}

	return a
}

func convertToAny(attrs []slog.Attr) []any {
	result := make([]any, len(attrs))
	for i, v := range attrs {
		result[i] = v
	}
	return result
}

// New khởi tạo Structured Logger JSON với tính năng bảo mật che giấu bí mật
func New(w io.Writer, level slog.Level) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}
	jsonHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(NewRedactionHandler(jsonHandler))
}
