package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/LouisHatton/insight-wave/internal/db/query"
)

func ToFirestoreDirection(direction query.OrderByDirection) firestore.Direction {
	if direction == query.OrderAsc {
		return firestore.Asc
	} else {
		return firestore.Desc
	}
}

func GenerateQuery(q firestore.Query, opts query.Options, wheres ...query.Where) *firestore.Query {
	if len(wheres) > 0 {
		for _, w := range wheres {
			q = q.Where(w.Key, string(w.Matcher), w.Value)
		}
	}

	if opts.OrderBy != nil {
		q = q.OrderBy(opts.OrderBy.Value, ToFirestoreDirection(opts.OrderBy.Direction))
	}

	if opts.Offset != nil {
		q = q.Offset(*opts.Offset)
	}

	if opts.Limit != nil {
		q = q.Limit(*opts.Limit)
	}

	return &q
}
