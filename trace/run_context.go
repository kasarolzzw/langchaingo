package trace

import (
	"context"

	"github.com/google/uuid"
)

var runContextKey = struct{}{}

type RunContext struct {
	TraceID     string
	RunID       string
	RunName     string
	ParentRunID string
}

func NewRunContext(ctx context.Context) context.Context {
	return newRunContextWithName(ctx, "")
}

func NewRunContextWithName(ctx context.Context, runName string) context.Context {
	return newRunContextWithName(ctx, runName)
}

func GetRunContext(ctx context.Context) RunContext {
	v := ctx.Value(runContextKey)
	if v != nil {
		if runContext, ok := v.(RunContext); ok {
			return runContext
		}
	}
	return RunContext{}
}

func newRunContextWithName(ctx context.Context, runName string) context.Context {
	v := ctx.Value(runContextKey)
	if v != nil {
		if parentRunContext, ok := v.(RunContext); ok &&
			parentRunContext.TraceID != "" &&
			parentRunContext.RunID != "" {
			return context.WithValue(ctx, runContextKey, newChildRunContext(parentRunContext, runName))
		}
	}
	return context.WithValue(ctx, runContextKey, newRootRunContext(runName))
}

func newRootRunContext(runName string) RunContext {
	return RunContext{
		TraceID: uuid.NewString(),
		RunID:   uuid.NewString(),
		RunName: runName,
	}
}

func newChildRunContext(parentRunContext RunContext, runName string) RunContext {
	return RunContext{
		TraceID:     parentRunContext.TraceID,
		RunID:       uuid.NewString(),
		RunName:     runName,
		ParentRunID: parentRunContext.RunID,
	}
}
