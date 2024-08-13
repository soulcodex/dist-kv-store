package storeapplication

import (
	domain "codesignal/internal/store/domain"
)

type DeleteStoreItemByKeyCommand struct {
	Key string
}

type DeleteStoreItemByKeyCommandHandler struct {
	repository domain.StoreItemRepository
}

func NewDeleteStoreItemByKeyCommandHandler(r domain.StoreItemRepository) *DeleteStoreItemByKeyCommandHandler {
	return &DeleteStoreItemByKeyCommandHandler{repository: r}
}

func (dsi *DeleteStoreItemByKeyCommandHandler) Handle(cmd *DeleteStoreItemByKeyCommand) error {
	storeKey, err := domain.NewStoreKey(cmd.Key)
	if err != nil {
		return err
	}

	return dsi.repository.DeleteByKey(storeKey)
}
