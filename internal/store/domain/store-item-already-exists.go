package storedomain

const storeItemAlreadyExistsErrorMessage = "Store item already exists"

type StoreItemAlreadyExists struct {
	context *map[string]interface{}
}

func NewStoreItemAlreadyExistsWithKey(key string) *StoreItemAlreadyExists {
	return &StoreItemAlreadyExists{
		context: &map[string]interface{}{
			"storage_key": key,
		},
	}
}

func (sie *StoreItemAlreadyExists) Error() string {
	return storeItemAlreadyExistsErrorMessage
}
