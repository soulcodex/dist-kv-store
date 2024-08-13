package storedomain

type StoreItemRepository interface {
	FindByKey(key StoreKey) (*StoreItem, error)
	Store(si *StoreItem) error
	DeleteByKey(key StoreKey) error
}
