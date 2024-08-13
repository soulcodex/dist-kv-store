package storeapplication

import (
	domain "codesignal/internal/store/domain"
)

type StoreItemResponse struct {
	Key   string
	Value interface{}
}

func NewStoreItemResponse(si *domain.StoreItem) (*StoreItemResponse, error) {
	return &StoreItemResponse{
		Key:   string(si.Key()),
		Value: si.Value().String(),
	}, nil
}

func (sir *StoreItemResponse) ToMap() map[string]interface{} {
	return map[string]interface{}{
		sir.Key: sir.Value,
	}
}
