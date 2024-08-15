package storeapplication

import (
	"encoding/json"

	"codesignal/internal/pkg/utils"
	domain "codesignal/internal/store/domain"
)

type CreateStoreItemCommand struct {
	Key   string
	Value interface{}
}

type CreateStoreItemCommandHandler struct {
	repository domain.StoreItemRepository
}

func NewCreateStoreItemCommandHandler(r domain.StoreItemRepository) *CreateStoreItemCommandHandler {
	return &CreateStoreItemCommandHandler{repository: r}
}

func (csh *CreateStoreItemCommandHandler) Handle(cmd *CreateStoreItemCommand) error {
	key, err := domain.NewStoreKey(cmd.Key)
	if err != nil {
		return err
	}

	var item *domain.StoreItem
	if ok := utils.ValueIsStringOrNumeric(cmd.Value); !ok {
		marshalled, marshalErr := json.Marshal(cmd.Value)
		if marshalErr != nil {
			return marshalErr
		}

		sv, svErr := domain.NewStoreValue(marshalled)
		if svErr != nil {
			return svErr
		}

		item = domain.NewStoreItem(key, sv)
	} else {
		sv, svErr := domain.NewStoreValue([]byte(cmd.Value.(string)))
		if svErr != nil {
			return svErr
		}

		item = domain.NewStoreItem(key, sv)
	}

	return csh.repository.Store(item)
}
