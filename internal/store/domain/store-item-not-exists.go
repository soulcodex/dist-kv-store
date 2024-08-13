package storedomain

const storeItemNotExistsErrorMessage = "Store item does not exists"

type StoreItemNotExists struct {
	context *map[string]interface{}
}

func NewStoreItemNotExistsWithKey(key string) *StoreItemNotExists {
	return &StoreItemNotExists{
		context: &map[string]interface{}{
			"storage_key": key,
		},
	}
}

func (sin *StoreItemNotExists) Error() string {
	return storeItemNotExistsErrorMessage
}
