package eventstore

import (
	"time"

	"github.com/LouisHatton/user-audit-saas/internal/events"
)

type Storer interface {
	AddEvent(*events.Event) error
	GetMany(from time.Time, to time.Time) ([]events.Event, error)
}
