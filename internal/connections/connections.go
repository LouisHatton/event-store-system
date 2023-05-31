package connections

import "time"

type ConnectionStatus string

const (
	Active    ConnectionStatus = "active"
	Deploying ConnectionStatus = "deploying"
	Paused    ConnectionStatus = "paused"
	Unhealthy ConnectionStatus = "unhealthy"
)

type Connection struct {
	// Unique id which identifies the connection (can't change)
	Id string `json:"id" firestore:"id"`

	// Unique Id which identifies the connection - used for the url string (can change)
	UrlId     string                 `json:"urlId" firestore:"urlId"`
	ProjectId string                 `json:"projectId" firestore:"projectId"`
	Name      string                 `json:"name" firestore:"name"`
	Metadata  Metadata               `json:"metadata" firestore:"metadata"`
	Tags      []string               `json:"tags" firestore:"tags"`
	Status    ConnectionStatus       `json:"status" firestore:"status"`
	Schema    map[string]interface{} `json:"schema" firestore:"schema"`
}

type Metadata struct {
	CreatedBy string    `json:"createdBy" firestore:"createdBy"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

type NewConnection struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func Empty() *Connection {
	return &Connection{
		Tags:   []string{},
		Schema: map[string]interface{}{},
	}
}
