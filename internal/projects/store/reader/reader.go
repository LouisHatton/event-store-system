package reader

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	dbFirestore "github.com/LouisHatton/insight-wave/internal/db/firestore"
	"github.com/LouisHatton/insight-wave/internal/db/query"
	"github.com/LouisHatton/insight-wave/internal/projects"
	projectStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	"go.uber.org/zap"
)

var _ projectStore.Reader = (*Reader)(nil)

type Reader struct {
	l          *zap.Logger
	collection string
	db         *firestore.CollectionRef
}

func New(logger *zap.Logger, collection string, client *firestore.Client) (*Reader, error) {
	r := Reader{
		l:          logger,
		collection: collection,
	}
	r.db = client.Collection(collection)
	return &r, nil
}

func (r *Reader) Get(id string) (*projects.Project, error) {
	logger := r.l.With(zap.String("projectId", id))

	logger.Debug("getting project doc")

	doc, err := r.db.Doc(id).Get(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error getting project: %w", err)
	}
	logger.Debug("fetched project doc")

	project := projects.Empty()
	err = doc.DataTo(project)
	if err != nil {
		return nil, fmt.Errorf("error converting response to project: %w", err)
	}

	return project, nil
}

func (r *Reader) Many(opts query.Options, wheres ...query.Where) (*[]projects.Project, error) {

	q := r.db.Query
	if len(wheres) > 0 {
		for _, w := range wheres {
			q = q.Where(w.Key, string(w.Matcher), w.Value)
		}
	}

	if opts.OrderBy != nil {
		q = q.OrderBy(opts.OrderBy.Value, dbFirestore.ToFirestoreDirection(opts.OrderBy.Direction))
	}

	if opts.Offset != nil {
		q = q.Offset(*opts.Offset)
	}

	if opts.Limit != nil {
		q = q.Limit(*opts.Limit)
	}

	itr := q.Documents(context.TODO())
	snapshots, err := itr.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching all documents: %w", err)
	}
	docs := []projects.Project{}
	for i, snap := range snapshots {
		docs = append(docs, *projects.Empty())
		err = snap.DataTo(&docs[i])
		if err != nil {
			return nil, fmt.Errorf("error converting response to project: %w", err)
		}
	}

	return &docs, nil
}
