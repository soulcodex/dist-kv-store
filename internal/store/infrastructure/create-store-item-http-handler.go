package storeinfrastructure

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog"
	"golang.org/x/exp/maps"

	"codesignal/internal/pkg/utils"

	application "codesignal/internal/store/application"
	domain "codesignal/internal/store/domain"
)

type CreateStoreItemRequest struct {
	Key   string
	Value interface{}
}

func (csr *CreateStoreItemRequest) IsValid() error {
	if csr.Key == "" {
		return errors.New("store item key is required")
	}

	if csr.Value == nil || csr.Value == "" {
		return errors.New("store item value is required")
	}

	return nil
}

type CreateStoreItemHttpHandler struct {
	log     zerolog.Logger
	handler *application.CreateStoreItemCommandHandler
}

func NewCreateStoreItemHttpHandler(log zerolog.Logger, handler *application.CreateStoreItemCommandHandler) *CreateStoreItemHttpHandler {
	return &CreateStoreItemHttpHandler{log: log, handler: handler}
}

func (csi *CreateStoreItemHttpHandler) Handle(w http.ResponseWriter, req *http.Request) {
	body, err := csi.requestOrError(req)
	if err != nil {
		response := map[string]interface{}{"message": err.Error()}
		utils.WriteHttpBadRequestResponse(req.Context(), response, csi.log, w)
		return
	}

	if bodyErr := body.IsValid(); bodyErr != nil {
		response := map[string]interface{}{"message": bodyErr.Error()}
		utils.WriteHttpBadRequestResponse(req.Context(), response, csi.log, w)
		return
	}

	cmd := &application.CreateStoreItemCommand{Key: body.Key, Value: body.Value}
	err = csi.handler.Handle(cmd)

	switch err.(type) {
	case nil:
		response := map[string]interface{}{"message": "key created successfully"}
		utils.WriteHttpOkResponse(req.Context(), response, csi.log, w)
		break
	case *domain.StoreItemAlreadyExists:
		response := map[string]interface{}{"message": "key already exist"}
		utils.WriteHttpConflictResponse(req.Context(), response, csi.log, w)
		break
	default:
		utils.WriteHttpInternalServerError(req.Context(), err, csi.log, w)
	}
}

func (csi *CreateStoreItemHttpHandler) requestOrError(req *http.Request) (*CreateStoreItemRequest, error) {
	content, err := utils.JsonBodyAsMap(req)
	if err != nil {
		return nil, err
	}

	contentKeys := maps.Keys(content)
	if len(contentKeys) < 1 || len(contentKeys) > 1 {
		return nil, errors.New("unexpected content keys")
	}

	key := maps.Keys(content)[0]

	var contentByKey interface{} = nil
	if keyValue, ok := content[key]; ok {
		contentByKey = keyValue
	}

	return &CreateStoreItemRequest{
		Key:   key,
		Value: contentByKey,
	}, nil
}
