package Writer

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/LouisHatton/insight-wave/internal/projects"
	projectStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	"go.uber.org/zap"
)

var _ projectStore.Writer = (*Writer)(nil)

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

func (r *Writer) Set(id string, project *projects.Project) error {
	logger := r.l.With(zap.String("projectId", id))

	logger.Debug("setting project doc")

	_, err := r.db.Doc(id).Create(context.TODO(), project)
	if err != nil {
		return fmt.Errorf("error getting project: %w", err)
	}
	logger.Debug("project doc set")

	return nil
}
