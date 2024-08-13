package storeinfrastructure

import (
	"strconv"

	"codesignal/internal/pkg/store"
	domain "codesignal/internal/store/domain"
)

type InMemoryStoreItemRepository struct {
	items store.KeyValueStore
}

func NewInMemoryStoreItemRepository(s store.KeyValueStore) *InMemoryStoreItemRepository {
	return &InMemoryStoreItemRepository{
		items: s,
	}
}

func (isr *InMemoryStoreItemRepository) FindByKey(key domain.StoreKey) (*domain.StoreItem, error) {
	item, err := isr.items.Get(key.String())
	if err != nil {
		return nil, domain.NewStoreItemNotExistsWithKey(key.String())
	}

	val, err := strconv.Unquote(item)
	if err != nil {
		return nil, domain.NewStoreItemNotExistsWithKey(key.String())
	}

	return domain.NewStoreItem(key, domain.StoreValue(val)), nil
}

func (isr *InMemoryStoreItemRepository) Store(si *domain.StoreItem) error {
	if _, err := isr.items.Get(si.Key().String()); err == nil {
		return domain.NewStoreItemAlreadyExistsWithKey(si.Key().String())
	}

	return isr.items.Set(si.Key().String(), si.Value().String())
}

func (isr *InMemoryStoreItemRepository) DeleteByKey(key domain.StoreKey) error {
	_, err := isr.FindByKey(key)
	if err != nil {
		return err
	}

	return isr.items.Delete(key.String())
}
