package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LouisHatton/insight-wave/internal/api/middleware"
	"github.com/LouisHatton/insight-wave/internal/api/routes"
	"github.com/LouisHatton/insight-wave/internal/config/enviroment"
	connectionsStore "github.com/LouisHatton/insight-wave/internal/connections/store"
	projectsStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Config struct {
	env enviroment.Type
}

type API struct {
	l                 *zap.Logger
	config            *Config
	projectStore      projectsStore.Manager
	connectionStore   connectionsStore.Manager
	born              time.Time
	authMiddleware    middleware.Auth
	projectMiddleware middleware.Project
}

func New(logger *zap.Logger, env enviroment.Type, authMiddleware *middleware.Auth, projectStore *projectsStore.Manager,
	connectionStore *connectionsStore.Manager) (*API, error) {

	projectMiddleware, err := middleware.NewProject(logger, &projectStore.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create project middleware: %w", err)
	}

	cfg := &Config{
		env: env,
	}
	api := API{
		l:                 logger,
		config:            cfg,
		projectStore:      *projectStore,
		connectionStore:   *connectionStore,
		born:              time.Now(),
		authMiddleware:    *authMiddleware,
		projectMiddleware: *projectMiddleware,
	}

	return &api, nil
}

func (api API) Register(r chi.Router) error {

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		render.Respond(w, r, map[string]string{
			"born":       api.born.Format(time.RFC3339),
			"enviroment": string(api.config.env),
		})
	})

	r.Route("/v1", func(r chi.Router) {

		r.Use(api.authMiddleware.Middleware)

		r.With(api.projectMiddleware.Middleware).Get(routes.ProjectIdPath, api.GetProject)
		r.Get(routes.ProjectPathPrefix, api.ListProjects)
		r.Post(routes.ProjectPathPrefix, api.CreateProject)

		r.With(api.projectMiddleware.Middleware).Get(routes.ConnectionIdPath, api.GetConnection)
		r.With(api.projectMiddleware.Middleware).Post(routes.ConnectionIdPath, api.EditConnection)
		r.With(api.projectMiddleware.Middleware).Get(routes.CreateConnectionsPath, api.ListConnections)
		r.With(api.projectMiddleware.Middleware).Post(routes.CreateConnectionsPath, api.CreateConnection)

	})

	return nil
}
