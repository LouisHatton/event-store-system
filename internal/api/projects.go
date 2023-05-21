package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LouisHatton/insight-wave/internal/api/responses"
	"github.com/LouisHatton/insight-wave/internal/db/query"
	"github.com/LouisHatton/insight-wave/internal/projects"
	"github.com/LouisHatton/insight-wave/internal/users"
	"github.com/LouisHatton/insight-wave/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	ProjectIdParam    = "projectId"
	ProjectUrlParam   = "/{projectId}"
	ProjectPathPrefix = "/project"
	CreateProjectPath = ProjectPathPrefix
	ProjectListPath   = ProjectPathPrefix + "/list"
	ProjectIdPath     = ProjectPathPrefix + ProjectUrlParam
)

func (api API) GetProject(w http.ResponseWriter, r *http.Request) {
	user := users.GetUserFromContext(r.Context())
	logger := api.l.With(zap.Any("userId", user.Id))

	id, err := getProjectIdFromUrl(r)
	if err != nil {
		api.l.Info("unable to fetch project id from url", zap.Error(err))
		render.Render(w, r, responses.NotFoundResponse("project"))
		return
	}

	logger = logger.With(zap.Any("projectId", id))

	project, err := api.projectStore.Reader.Get(id)
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

	render.JSON(w, r, &project)
}

func (api API) CreateProject(w http.ResponseWriter, r *http.Request) {
	user := users.GetUserFromContext(r.Context())
	logger := api.l.With(zap.String("userId", user.Id))

	data := projects.NewProject{}
	if err := render.Decode(r, &data); err != nil {
		logger.Error("error parsing provided project data", zap.Error(err))
		render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	id := uuid.New().String()
	logger = logger.With(zap.String("projectId", id))
	newProject := projects.Project{
		Id:   id,
		Name: data.Name,
		Metadata: projects.Metadata{
			CreatedBy: user.Id,
			CreatedAt: time.Now(),
		},
		Config: projects.Config{
			Colour: data.Colour,
		},
	}

	if err := api.projectStore.Set(id, &newProject); err != nil {
		logger.Error("failed to store new project", zap.Error(err))
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}

	logger.Info("created new project")
	render.Respond(w, r, &newProject)
}

func (api API) ListProjects(w http.ResponseWriter, r *http.Request) {
	user := users.GetUserFromContext(r.Context())
	logger := api.l.With(zap.Any("userId", user.Id))

	docs, err := api.projectStore.Many(query.Options{}, query.Where{
		Key:     "users",
		Matcher: query.Contains,
		Value:   user.Id,
	})
	if err != nil {
		logger.Fatal("failed to fetch documents", zap.Error(err))
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}

	render.JSON(w, r, docs)
}

func getProjectIdFromUrl(r *http.Request) (string, error) {
	if projectId := chi.URLParam(r, ProjectIdParam); projectId != "" {
		return projectId, nil
	} else {
		return "", fmt.Errorf("url does not contain project id: url: %s", r.URL.String())
	}
}
