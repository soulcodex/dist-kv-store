package di

import (
	"net/http"

	"codesignal/internal/pkg/store"
	application "codesignal/internal/store/application"
	domain "codesignal/internal/store/domain"
	infrastructure "codesignal/internal/store/infrastructure"

	"github.com/julienschmidt/httprouter"
)

type StoreServices struct {
	Repository domain.StoreItemRepository

	// Handlers
	CreateStoreItemHandler *application.CreateStoreItemCommandHandler
	FetchStoreItemHandler  *application.FetchStoreItemByKeyQueryHandler
	DeleteStoreItemHandler *application.DeleteStoreItemByKeyCommandHandler

	// Entrypoints
	CreateStoreHttpEntrypoint     *infrastructure.CreateStoreItemHttpHandler
	FetchStoreItemHttpEntrypoint  *infrastructure.FetchStoreByKeyItemHttpHandler
	DeleteStoreItemHttpEntrypoint *infrastructure.DeleteStoreByKeyItemHttpHandler
}

func InitStoreModule(srv *CommonServices) *StoreServices {
	repository := infrastructure.NewInMemoryStoreItemRepository(srv.KeyValueStore)

	createHandler := application.NewCreateStoreItemCommandHandler(repository)
	fetchHandler := application.NewFetchStoreItemByKeyQueryHandler(repository)
	deleteHandler := application.NewDeleteStoreItemByKeyCommandHandler(repository)

	createHttpEntrypoint := infrastructure.NewCreateStoreItemHttpHandler(srv.Log, createHandler)
	fetchHttpEntrypoint := infrastructure.NewFetchStoreByKeyItemHttpHandler(srv.Log, fetchHandler)
	deleteHttpEntrypoint := infrastructure.NewDeleteStoreByKeyItemHttpHandler(srv.Log, deleteHandler)

	return &StoreServices{
		Repository: repository,

		FetchStoreItemHandler:  fetchHandler,
		CreateStoreItemHandler: createHandler,
		DeleteStoreItemHandler: deleteHandler,

		CreateStoreHttpEntrypoint:     createHttpEntrypoint,
		FetchStoreItemHttpEntrypoint:  fetchHttpEntrypoint,
		DeleteStoreItemHttpEntrypoint: deleteHttpEntrypoint,
	}
}

func (sm *StoreServices) RegisterHttpRoutes(router *httprouter.Router, srv *CommonServices) {
	router.HandlerFunc(http.MethodPost, "/key", sm.CreateStoreHttpEntrypoint.Handle)
	router.HandlerFunc(http.MethodGet, "/key/:key", sm.FetchStoreItemHttpEntrypoint.Handle)
	router.HandlerFunc(http.MethodDelete, "/key/:key", sm.DeleteStoreItemHttpEntrypoint.Handle)

	// Raft handlers
	router.HandlerFunc(http.MethodPost, "/join", store.NodeRaftJoinHttpHandler(srv.KeyValueStore.Consensus(), srv.Log))
	router.HandlerFunc(http.MethodPost, "/unlink", store.NodeRaftUnlinkHttpHandler(srv.KeyValueStore.Consensus(), srv.Log))
	router.HandlerFunc(http.MethodGet, "/stats", store.NodeRaftStatsHttpHandler(srv.KeyValueStore.Consensus(), srv.Log))
}
