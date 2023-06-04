package writer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LouisHatton/insight-wave/internal/events"
	"github.com/LouisHatton/insight-wave/internal/events/store"
	"go.uber.org/zap"
)

var _ store.Writer = (*TinyBird)(nil)

type TinyBird struct {
	logger      zap.Logger
	CreateToken *string
	DeleteToken *string
}

const (
	deleteUrl   = "https://api.tinybird.co/v0/datasources/"
	newEventUrl = "https://api.tinybird.co/v0/events?name="
)

func New(logger zap.Logger, createToken *string, deleteToken *string) *TinyBird {
	svc := TinyBird{
		logger:      logger,
		CreateToken: createToken,
		DeleteToken: deleteToken,
	}
	return &svc
}

func (t *TinyBird) Add(ctx context.Context, event events.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		t.logger.Error("Error marshaling the event: ", zap.Error(err))
		return err
	}

	request, err := http.NewRequest(http.MethodPost, newEventUrl+event.ConnectionId, bytes.NewBuffer(data))
	if err != nil {
		t.logger.Error("Error creating the request: ", zap.Error(err))
		return err
	}

	var token string
	if t.CreateToken == nil {
		return fmt.Errorf("create token is empty in config")
	}
	token = *t.CreateToken

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		t.logger.Error("Error creating sending the request: ", zap.Error(err))
		return err
	}

	return nil
}

func (t *TinyBird) DeleteSource(ctx context.Context, connectionId string) error {
	request, err := http.NewRequest(http.MethodDelete, deleteUrl+connectionId, nil)
	if err != nil {
		return fmt.Errorf("error creating delete http request: %w", err)
	}

	var token string
	if t.CreateToken == nil {
		return fmt.Errorf("create token is empty in config")
	}
	token = *t.DeleteToken

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		t.logger.Error("Error creating sending the request: ", zap.Error(err))
		return err
	}

	return nil
}
