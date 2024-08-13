package storeinfrastructure

import (
	"net/http"

	"github.com/rs/zerolog"

	"codesignal/internal/pkg/utils"

	application "codesignal/internal/store/application"
	domain "codesignal/internal/store/domain"

	"github.com/julienschmidt/httprouter"
)

type FetchStoreByKeyItemHttpHandler struct {
	log     zerolog.Logger
	handler *application.FetchStoreItemByKeyQueryHandler
}

func NewFetchStoreByKeyItemHttpHandler(log zerolog.Logger, handler *application.FetchStoreItemByKeyQueryHandler) *FetchStoreByKeyItemHttpHandler {
	return &FetchStoreByKeyItemHttpHandler{log: log, handler: handler}
}

func (fsi *FetchStoreByKeyItemHttpHandler) Handle(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	storeKey := params.ByName("key")

	query := &application.FetchStoreItemByKeyQuery{Key: storeKey}
	item, err := fsi.handler.Handle(query)

	switch err.(type) {
	case nil:
		utils.WriteHttpOkResponse(req.Context(), item.ToMap(), fsi.log, w)
		break
	case *domain.StoreItemNotExists:
		response := map[string]interface{}{"message": "key not found"}
		utils.WriteHttpNotFoundError(req.Context(), response, err, fsi.log, w)
		break
	default:
		utils.WriteHttpInternalServerError(req.Context(), err, fsi.log, w)
	}
}
