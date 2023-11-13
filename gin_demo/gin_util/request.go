package ginUtil

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

const (
	CtxRequestID = "X-Request-Id"
)
func UUID() string {
	uid, _ := uuid.NewRandom()
	return strings.ReplaceAll(uid.String(), "-", "")
}

func WithRequestID(ctx context.Context) context.Context {
	if _, ok := ctx.Value(CtxRequestID).(string); ok {
		return ctx
	}
	return context.WithValue(ctx, CtxRequestID, UUID())
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(CtxRequestID).(string)
	if ok {
		return requestID
	}
	return ""
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, CtxRequestID, requestID)
}

func GenerateRequestID() string {
	return UUID()
}