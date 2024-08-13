package storeinfrastructure

import (
	"net/http"

	"github.com/rs/zerolog"

	"codesignal/internal/pkg/utils"

	application "codesignal/internal/store/application"
	domain "codesignal/internal/store/domain"

	"github.com/julienschmidt/httprouter"
)

type DeleteStoreByKeyItemHttpHandler struct {
	log     zerolog.Logger
	handler *application.DeleteStoreItemByKeyCommandHandler
}

func NewDeleteStoreByKeyItemHttpHandler(log zerolog.Logger, handler *application.DeleteStoreItemByKeyCommandHandler) *DeleteStoreByKeyItemHttpHandler {
	return &DeleteStoreByKeyItemHttpHandler{log: log, handler: handler}
}

func (fsi *DeleteStoreByKeyItemHttpHandler) Handle(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	storeKey := params.ByName("key")

	cmd := &application.DeleteStoreItemByKeyCommand{Key: storeKey}
	err := fsi.handler.Handle(cmd)

	switch err.(type) {
	case nil:
		response := map[string]interface{}{"message": "key deleted successfully"}
		utils.WriteHttpOkResponse(req.Context(), response, fsi.log, w)
		break
	case *domain.StoreItemNotExists:
		response := map[string]interface{}{"message": "key not found"}
		utils.WriteHttpNotFoundError(req.Context(), response, err, fsi.log, w)
		break
	default:
		utils.WriteHttpInternalServerError(req.Context(), err, fsi.log, w)
	}
}
