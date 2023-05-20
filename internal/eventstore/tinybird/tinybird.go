package tinybird

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/LouisHatton/insight-wave/internal/events"
	"github.com/LouisHatton/insight-wave/internal/eventstore"
	"github.com/caarlos0/env/v8"
	"go.uber.org/zap"
)

var _ eventstore.Storer = (*TinyBird)(nil)

type TinyBird struct {
	log       *zap.Logger
	Url       string `env:"TINYBIRD_API_URL"`
	Token     string `env:"TINYBIRD_API_TOKEN"`
	EventName string `env:"TINYBIRD_EVENT_NAME"`
}

func New(log zap.Logger) *TinyBird {
	cfg := TinyBird{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("There was an error parsing the env for TinyBird: ", zap.Error(err))
	}
	cfg.log = &log

	return &cfg
}

func (t *TinyBird) AddEvent(event *events.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		t.log.Error("Error marshaling the event: ", zap.Error(err))
		return err
	}

	request, err := http.NewRequest(http.MethodPost, t.Url+t.EventName, bytes.NewBuffer(data))
	if err != nil {
		t.log.Error("Error creating the request: ", zap.Error(err))
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+t.Token)
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		t.log.Error("Error creating sending the request: ", zap.Error(err))
		return err
	}

	return nil
}

func (tinybird *TinyBird) GetMany(from time.Time, to time.Time) ([]events.Event, error) {
	return nil, nil
}
