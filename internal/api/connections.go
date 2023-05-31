package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LouisHatton/insight-wave/internal/api/responses"
	"github.com/LouisHatton/insight-wave/internal/api/routes"
	"github.com/LouisHatton/insight-wave/internal/connections"
	internalContext "github.com/LouisHatton/insight-wave/internal/context"
	"github.com/LouisHatton/insight-wave/internal/db/query"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (api *API) GetConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := internalContext.GetUserFromContext(ctx)
	logger := api.l.With(zap.String("userId", user.Id))
	project, ok := internalContext.GetProjectFromContext(ctx)
	if !ok {
		logger.Error("unable to get project from context")
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}
	logger = logger.With(zap.String("projectId", project.Id))

	id, err := getConnectionIdFromUrl(r)
	if err != nil {
		logger.Error("unable to get connection id from url", zap.Error(err))
		render.Render(w, r, responses.NotFoundResponse("connection"))
		return
	}
	logger = logger.With(zap.String("connectionId", id))

	connection, err := api.connectionStore.Get(id)
	if err != nil {
		logger.Error("error getting document", zap.Error(err))
		render.Render(w, r, responses.NotFoundResponse("connection"))
		return
	}

	if connection.ProjectId != project.Id {
		logger.Warn("connection is not a member of the project")
		render.Render(w, r, responses.NotFoundResponse("connection"))
		return
	}

	render.JSON(w, r, &connection)
}

func (api *API) CreateConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := internalContext.GetUserFromContext(ctx)
	project, ok := internalContext.GetProjectFromContext(ctx)
	if !ok {
		api.l.Error("unable to get project from context")
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}
	logger := api.l.With(zap.String("userId", user.Id), zap.String("projectId", project.Id))

	data := connections.NewConnection{}
	if err := render.Decode(r, &data); err != nil {
		logger.Error("error parsing provided connection data", zap.Error(err))
		render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	id := uuid.New().String()
	urlId := uuid.New().String()
	logger = logger.With(zap.String("connectionId", id))
	newConnection := connections.Connection{
		Id:        id,
		UrlId:     urlId,
		ProjectId: project.Id,
		Name:      data.Name,
		Tags:      data.Tags,
		Metadata: connections.Metadata{
			CreatedBy: user.Id,
			CreatedAt: time.Now(),
		},
		Status: connections.Deploying,
	}

	if err := api.connectionStore.Set(id, &newConnection); err != nil {
		logger.Error("failed to store new connection", zap.Error(err))
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}

	logger.Info("new connection created")
	render.JSON(w, r, &newConnection)
}

func (api *API) ListConnections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := internalContext.GetUserFromContext(ctx)
	project, ok := internalContext.GetProjectFromContext(ctx)
	if !ok {
		api.l.Error("unable to get project from context")
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}
	logger := api.l.With(zap.String("userId", user.Id), zap.String("projectId", project.Id))

	docs, err := api.connectionStore.Many(query.Options{}, query.Where{
		Key:     "projectId",
		Matcher: query.EqualTo,
		Value:   project.Id,
	})
	if err != nil {
		logger.Fatal("failed to fetch connections", zap.Error(err))
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}

	render.JSON(w, r, docs)
}

func getConnectionIdFromUrl(r *http.Request) (string, error) {
	if id := chi.URLParam(r, routes.ConnectionIdParam); id != "" {
		return id, nil
	} else {
		return "", fmt.Errorf("url does not contain connection id: url: %s", r.URL.String())
	}
}
