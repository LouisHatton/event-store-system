package middleware

import (
	"fmt"
	"net/http"

	"github.com/LouisHatton/insight-wave/internal/api/responses"
	"github.com/LouisHatton/insight-wave/internal/api/routes"
	internalContext "github.com/LouisHatton/insight-wave/internal/context"
	projectsStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	"github.com/LouisHatton/insight-wave/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Project struct {
	logger       *zap.Logger
	projectStore projectsStore.Reader
}

func NewProject(l *zap.Logger, projectReader *projectsStore.Reader) (*Project, error) {
	return &Project{
		logger:       l,
		projectStore: *projectReader,
	}, nil
}

func (m *Project) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := internalContext.GetUserFromContext(ctx)
		logger := m.logger.With(zap.Any("userId", user.Id))

		id, err := getProjectIdFromUrl(r)
		if err != nil {
			logger.Warn("unable to get project id from url", zap.Error(err))
			next.ServeHTTP(w, r)
			return
		}

		logger = logger.With(zap.Any("projectId", id))

		project, err := m.projectStore.Get(id)
		if err != nil {
			logger.Warn("error getting document", zap.Error(err))
			render.Render(w, r, responses.NotFoundResponse("project"))
			return
		}

		if !utils.Contains(project.AllUsers, user.Id) {
			logger.Info("user is not in requested project")
			render.Render(w, r, responses.ErrForbidden())
			return
		}

		ctx = internalContext.AddProjectToContext(ctx, *project)

		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func getProjectIdFromUrl(r *http.Request) (string, error) {
	if projectId := chi.URLParam(r, routes.ProjectIdParam); projectId != "" {
		return projectId, nil
	} else {
		return "", fmt.Errorf("url does not contain project id: url: %s", r.URL.String())
	}
}
