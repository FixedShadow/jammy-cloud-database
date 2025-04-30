package common

import (
	"context"
	"github.com/google/uuid"
	"strings"
)

type ContextKey string

func UUID() string {
	return uuid.New().String()
}

func Generate32RandomString() string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, ContextKey("trace-id"), traceId)
}

func TraceId(ctx context.Context) string {
	v := ctx.Value(ContextKey("trace-id"))
	traceId, _ := v.(string)
	return traceId
}

