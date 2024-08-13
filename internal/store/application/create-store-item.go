package storeapplication

import (
	"encoding/json"

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

	rawContent, err := json.Marshal(cmd.Value)
	if err != nil {
		return err
	}

	content, err := domain.NewStoreValue(rawContent)
	if err != nil {
		return err
	}

	item := domain.NewStoreItem(key, content)
	return csh.repository.Store(item)
}
