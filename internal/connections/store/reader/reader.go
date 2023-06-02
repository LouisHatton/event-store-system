package reader

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/LouisHatton/insight-wave/internal/connections"
	connectionStore "github.com/LouisHatton/insight-wave/internal/connections/store"
	dbFirestore "github.com/LouisHatton/insight-wave/internal/db/firestore"
	"github.com/LouisHatton/insight-wave/internal/db/query"

	"go.uber.org/zap"
)

var _ connectionStore.Reader = (*Reader)(nil)

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

func (r *Reader) Get(id string) (*connections.Connection, error) {
	logger := r.l.With(zap.String("connectionId", id))

	logger.Debug("getting connection doc")

	doc, err := r.db.Doc(id).Get(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %w", err)
	}
	logger.Debug("fetched connection doc")

	connection := connections.Empty()
	err = doc.DataTo(&connection)
	if err != nil {
		return nil, fmt.Errorf("error converting response to connection struct: %w", err)
	}

	return &connection, nil
}

func (r *Reader) Many(opts query.Options, wheres ...query.Where) (*[]connections.Connection, error) {

	q := dbFirestore.GenerateQuery(r.db.Query, opts, wheres...)

	itr := q.Documents(context.TODO())
	snapshots, err := itr.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching all documents: %w", err)
	}
	docs := []connections.Connection{}
	for i, snap := range snapshots {
		docs = append(docs, connections.Empty())
		err = snap.DataTo(&docs[i])
		if err != nil {
			return nil, fmt.Errorf("error converting response to connection struct: %w", err)
		}
	}

	return &docs, nil
}
