package storeapplication

import (
	domain "codesignal/internal/store/domain"
)

type FetchStoreItemByKeyQuery struct {
	Key string
}

type FetchStoreItemByKeyQueryHandler struct {
	repository domain.StoreItemRepository
}

func NewFetchStoreItemByKeyQueryHandler(r domain.StoreItemRepository) *FetchStoreItemByKeyQueryHandler {
	return &FetchStoreItemByKeyQueryHandler{repository: r}
}

func (fsi *FetchStoreItemByKeyQueryHandler) Handle(cmd *FetchStoreItemByKeyQuery) (*StoreItemResponse, error) {
	storeKey, err := domain.NewStoreKey(cmd.Key)
	if err != nil {
		return nil, err
	}

	storeItem, err := fsi.repository.FindByKey(storeKey)
	if err != nil {
		return nil, err
	}

	return NewStoreItemResponse(storeItem)
}
