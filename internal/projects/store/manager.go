package store

import (
	"github.com/LouisHatton/insight-wave/internal/db/query"
	"github.com/LouisHatton/insight-wave/internal/projects"
)

type Reader interface {
	Get(id string) (*projects.Project, error)
	One(opts query.Options, wheres ...query.Where) (*projects.Project, error)
	Many(opts query.Options, wheres ...query.Where) (*[]projects.Project, error)
}

type Writer interface {
	Set(id string, project *projects.Project) error
}

type Manager struct {
	Reader
	Writer
}

func New(r Reader, w Writer) *Manager {
	return &Manager{r, w}
}
