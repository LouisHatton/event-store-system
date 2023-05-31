package context

import (
	"context"

	"github.com/LouisHatton/insight-wave/internal/projects"
)

func AddProjectToContext(ctx context.Context, project projects.Project) context.Context {
	return context.WithValue(ctx, projectContextKey, project)
}

func GetProjectFromContext(ctx context.Context) (projects.Project, bool) {
	p, ok := ctx.Value(projectContextKey).(projects.Project)
	return p, ok
}
