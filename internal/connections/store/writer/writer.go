package writer

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/LouisHatton/insight-wave/internal/connections"
	connectionStore "github.com/LouisHatton/insight-wave/internal/connections/store"
	"go.uber.org/zap"
)

var _ connectionStore.Writer = (*Writer)(nil)

type Writer struct {
	l          *zap.Logger
	collection string
	db         *firestore.CollectionRef
}

func New(logger *zap.Logger, collection string, client *firestore.Client) (*Writer, error) {
	r := Writer{
		l:          logger,
		collection: collection,
	}
	r.db = client.Collection(collection)
	return &r, nil
}

func (r *Writer) Set(id string, connection *connections.Connection) error {
	logger := r.l.With(zap.String("connectionId", id))

	logger.Debug("setting connection doc")

	_, err := r.db.Doc(id).Create(context.TODO(), connection)
	if err != nil {
		return fmt.Errorf("error getting connection: %w", err)
	}
	logger.Debug("connection doc set")

	return nil
}
