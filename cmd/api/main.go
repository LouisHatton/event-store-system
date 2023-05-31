package main

import (
	"context"

	"net/http"

	firebase "firebase.google.com/go/v4"
	"go.uber.org/zap"

	"github.com/LouisHatton/insight-wave/internal/api"
	apiMiddleware "github.com/LouisHatton/insight-wave/internal/api/middleware"
	"github.com/LouisHatton/insight-wave/internal/config/appconfig"
	"github.com/LouisHatton/insight-wave/internal/config/enviroment"
	connectionsStore "github.com/LouisHatton/insight-wave/internal/connections/store"
	connectionsStoreReader "github.com/LouisHatton/insight-wave/internal/connections/store/reader"
	connectionsStoreWriter "github.com/LouisHatton/insight-wave/internal/connections/store/writer"
	projectsStore "github.com/LouisHatton/insight-wave/internal/projects/store"
	projectsStoreReader "github.com/LouisHatton/insight-wave/internal/projects/store/reader"
	projectsStoreWriter "github.com/LouisHatton/insight-wave/internal/projects/store/writer"
	"github.com/caarlos0/env/v8"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	appconfig.Enviroment
	appconfig.Server
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

	authClient, err := app.Auth(ctx)
	if err != nil {
		logger.Fatal("error initializing app auth", zap.Error(err))
	}

	store, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatal("error initializing firestore", zap.Error(err))
	}

	// --- Middleware
	authMiddleware, err := apiMiddleware.NewAuth(logger, authClient)
	if err != nil {
		logger.Fatal("error initializing auth middleware", zap.Error(err))
	}

	// --- Projects Store
	const ProjectStoreCollectionName = "projects"
	projectReader, err := projectsStoreReader.New(logger, ProjectStoreCollectionName, store)
	if err != nil {
		logger.Fatal("error initializing projectsStoreReader", zap.Error(err))
	}

	projectsWriter, err := projectsStoreWriter.New(logger, ProjectStoreCollectionName, store)
	if err != nil {
		logger.Fatal("error initializing projectsStoreReader", zap.Error(err))
	}
	projectStore := projectsStore.New(projectReader, projectsWriter)

	// --- Connections Store
	const ConnectionsStoreCollectionName = "connections"
	connectionReader, err := connectionsStoreReader.New(logger, ConnectionsStoreCollectionName, store)
	if err != nil {
		logger.Fatal("error initializing connectionsStoreReader", zap.Error(err))
	}

	connectionWriter, err := connectionsStoreWriter.New(logger, ConnectionsStoreCollectionName, store)
	if err != nil {
		logger.Fatal("error initializing connectionsStoreWriter", zap.Error(err))
	}
	connectionStore := connectionsStore.New(connectionReader, connectionWriter)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	api, err := api.New(logger, cfg.Enviroment.CurrentEnv, authMiddleware, projectStore, connectionStore)
	if err != nil {
		logger.Fatal("error initializing api", zap.Error(err))
	}

	err = api.Register(r)
	if err != nil {
		logger.Fatal("error registering api routes", zap.Error(err))
	}

	logger.Info("Webserver started", zap.String("port", cfg.Port), zap.String("env", string(cfg.Enviroment.CurrentEnv)))
	http.ListenAndServe(":"+cfg.Port, r)
}
