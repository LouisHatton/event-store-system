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
