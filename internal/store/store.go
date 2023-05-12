package store

type Store interface {
	Put(doc interface{}) error
	Get(id string, doc interface{}) error
}
