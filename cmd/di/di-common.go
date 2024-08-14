package di

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"os"
	"time"

	"codesignal/internal/pkg/config"
	"codesignal/internal/pkg/server"
	"codesignal/internal/pkg/store"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

type CommonServices struct {
	Log           zerolog.Logger
	HttpServer    *server.Server
	HttpRouter    *httprouter.Router
	KeyValueStore store.KeyValueStore
}

type OktaDistributedKeyValueStorageContainer struct {
	Services      *CommonServices
	Config        *config.Config
	StoreServices *StoreServices
}

func Init(configuration *config.Config) *OktaDistributedKeyValueStorageContainer {
	return bootstrapApp(configuration)
}

func bootstrapApp(configuration *config.Config) *OktaDistributedKeyValueStorageContainer {
	logger := zerolog.New(os.Stderr).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Logger()

	httpRouter := httprouter.New()
	httpServer := server.New(logger, configuration.Server, httpRouter)

	inMemoryStore := store.NewInMemoryKeyValueStore(logger, configuration.NodeConfig)

	services := &CommonServices{
		Log:           logger,
		HttpServer:    httpServer,
		HttpRouter:    httpRouter,
		KeyValueStore: inMemoryStore,
	}

	storeModule := InitStoreModule(services)
	storeModule.RegisterHttpRoutes(httpRouter, services)

	return &OktaDistributedKeyValueStorageContainer{
		Services:      services,
		Config:        configuration,
		StoreServices: storeModule,
	}
}

func BuildNodeJoinRetryableHttpClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 3 * time.Second
	client.RetryWaitMax = 6 * time.Second

	return client
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}
