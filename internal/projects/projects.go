package projects

import "time"

type Project struct {
	Id             string   `json:"id" firestore:"id"`
	Name           string   `json:"name" firestore:"name"`
	Metadata       Metadata `json:"metadata" firestore:"metadata"`
	Config         Config   `json:"config" firestore:"config"`
	AllUsers       []string `json:"users" firestore:"users"`
	AdminUsers     []string `json:"adminUsers" firestore:"adminUsers"`
	ReadOnlyUsers  []string `json:"roUsers" firestore:"roUsers"`
	ReadWriteUsers []string `json:"rwUsers" firestore:"rwUsers"`
}

type Config struct {
	Colour string `json:"colour" firestore:"colour"`
}

type Metadata struct {
	CreatedBy string    `json:"createdBy" firestore:"createdBy"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

type NewProject struct {
	Name   string `json:"name"`
	Colour string `json:"colour"`
}

func Empty() *Project {
	return &Project{
		AllUsers:       []string{},
		AdminUsers:     []string{},
		ReadOnlyUsers:  []string{},
		ReadWriteUsers: []string{},
	}
}
