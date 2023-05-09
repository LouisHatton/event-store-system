package events

import "time"

type Event struct {
	Timestamp *time.Time   `json:"timestamp"`
	ProjectId string       `json:"projectId"`
	Version   EventVersion `json:"version"`
	Payload   string       `json:"payload"`
}

type EventVersion string

var (
	V1 EventVersion = "1"
)
