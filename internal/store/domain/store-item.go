package storedomain

type StoreItem struct {
	key   StoreKey
	value StoreValue
}

func NewStoreItem(key StoreKey, value StoreValue) *StoreItem {
	return &StoreItem{key: key, value: value}
}

func (si *StoreItem) Key() StoreKey {
	return si.key
}

func (si *StoreItem) Value() StoreValue {
	return si.value
}
