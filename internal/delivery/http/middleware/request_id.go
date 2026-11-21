package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type ctxKey string

const RequestIDKey ctxKey = "request_id"
const HeaderRequestID = "X-Request-ID"

// RequestID tạo hoặc truyền tiếp Request ID duy nhất cho mỗi HTTP request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(HeaderRequestID)
		if reqID == "" {
			bytes := make([]byte, 16)
			_, _ = rand.Read(bytes)
			reqID = hex.EncodeToString(bytes)
		}

		w.Header().Set(HeaderRequestID, reqID)
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value(RequestIDKey).(string); ok {
		return val
	}
	return "unknown"
}
