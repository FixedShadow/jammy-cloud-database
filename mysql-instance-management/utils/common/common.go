package common

import (
	"context"
	"github.com/google/uuid"
	"strings"
)

func GenerateRandomStringLess32(length int) string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")[:length]
}

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, "trace-id", traceId)
}

func TraceId(ctx context.Context) string {
	v := ctx.Value("trace-id")
	traceId, _ := v.(string)
	return traceId
}

func ProjectId(ctx context.Context) string {
	v := ctx.Value("projectId")
	projectId, _ := v.(string)
	return projectId
}

func WithProjectId(ctx context.Context, projectId string) context.Context {
	return context.WithValue(ctx, "projectId", projectId)
}
