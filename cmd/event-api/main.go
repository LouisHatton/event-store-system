package main

import (
	"context"
	"io"
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/LouisHatton/insight-wave/internal/api/responses"
	"github.com/LouisHatton/insight-wave/internal/config/appconfig"
	"github.com/LouisHatton/insight-wave/internal/config/enviroment"
	connections_store "github.com/LouisHatton/insight-wave/internal/connections/store"
	connections_store_reader "github.com/LouisHatton/insight-wave/internal/connections/store/reader"
	"github.com/LouisHatton/insight-wave/internal/events"
	events_store "github.com/LouisHatton/insight-wave/internal/events/store"
	events_store_writer "github.com/LouisHatton/insight-wave/internal/events/store/writer"
	"github.com/caarlos0/env/v8"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type config struct {
	appconfig.TinyBird
	appconfig.Enviroment
	appconfig.Server
}

type eventApi struct {
	logger          zap.Logger
	eventsWriter    events_store.Writer
	connectionStore connections_store.Reader
}

func main() {
	// --- ENV & Logging
	ctx := context.Background()
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		panic("failed to parse server config env: " + err.Error())
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	if cfg.Enviroment.CurrentEnv == enviroment.Production {
		logger, err = zap.NewProduction()
		if err != nil {
			panic("failed to create production logger: " + err.Error())
		}
	}
	defer logger.Sync()

	// --- GCloud
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "insight-wave-dev",
	})
	if err != nil {
		logger.Fatal("error initializing app", zap.Error(err))
	}

	store, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatal("error initializing firestore", zap.Error(err))
	}

	connectionStore, err := connections_store_reader.New(logger, "connections", store)
	if err != nil {
		logger.Fatal("error initializing projects store reader", zap.Error(err))
	}

	eventsWriter := events_store_writer.New(*logger, &cfg.TinyBird.DatasourcesCreateToken, &cfg.DatasourcesDeleteToken)

	api := eventApi{
		logger:          *logger,
		eventsWriter:    eventsWriter,
		connectionStore: connectionStore,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	api.Register(r)
	if err != nil {
		logger.Fatal("error registering api routes", zap.Error(err))
	}

	logger.Info("Webserver started", zap.String("port", cfg.Port), zap.String("env", string(cfg.Enviroment.CurrentEnv)))
	http.ListenAndServe(":"+cfg.Port, r)
}

func (api *eventApi) Register(r *chi.Mux) error {

	r.Post("/{connectionUid}", api.HandleEvent)
	return nil
}

func (api *eventApi) HandleEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlId := chi.URLParam(r, "connectionUid")

	connection, err := api.connectionStore.GetByUrl(urlId)
	if err != nil {
		api.logger.Warn("error getting connection document from url id", zap.Error(err))
		render.Render(w, r, responses.NotFoundResponse("connection"))
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		api.logger.Error("failed to read body of request", zap.Error(err))
		render.Render(w, r, responses.ErrInternalServerError())
		return
	}
	bodyString := string(bodyBytes)

	newEvent := events.Event{
		Timestamp:    time.Now(),
		ConnectionId: connection.Id,
		Version:      events.V1,
		Payload:      bodyString,
	}
	api.eventsWriter.Add(ctx, newEvent)
}
