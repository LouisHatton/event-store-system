package eventstore

import (
	"time"

	"github.com/LouisHatton/insight-wave/internal/events"
)

type Storer interface {
	AddEvent(*events.Event) error
	GetMany(from time.Time, to time.Time) ([]events.Event, error)
}
