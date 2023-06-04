package store

import (
	"context"
	"time"

	"github.com/LouisHatton/insight-wave/internal/events"
)

type Reader interface {
	Get(ctx context.Context, from time.Time, to time.Time) ([]events.Event, error)
}

type Writer interface {
	Add(ctx context.Context, event events.Event) error
	DeleteSource(ctx context.Context, connectionId string) error
}

type Manager interface {
	Reader
	Writer
}
