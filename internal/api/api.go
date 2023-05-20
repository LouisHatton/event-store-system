package api

import (
	"net/http"
	"time"

	"github.com/LouisHatton/insight-wave/internal/api/middleware"
	"github.com/LouisHatton/insight-wave/internal/config/enviroment"
	projectsStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Config struct {
	env enviroment.Type
}

type API struct {
	l              *zap.Logger
	config         *Config
	projectStore   *projectsStore.Manager
	born           time.Time
	authMiddleware *middleware.Middleware
}

func New(logger *zap.Logger, env enviroment.Type, authMiddleware *middleware.Middleware, projectStore *projectsStore.Manager) (*API, error) {
	cfg := &Config{
		env: env,
	}
	api := API{
		l:              logger,
		config:         cfg,
		projectStore:   projectStore,
		born:           time.Now(),
		authMiddleware: authMiddleware,
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

		r.Use(api.authMiddleware.AuthMiddleware)

		r.Get(ProjectIdPath, api.GetProject)
		r.Get(ProjectListPath, api.ListProjects)
		r.Post(ProjectPathPrefix, api.CreateProject)
	})

	return nil
}
