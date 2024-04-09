package handler

import (
	"context"
)

type Handler[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}
