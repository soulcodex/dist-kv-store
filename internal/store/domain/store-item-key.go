package storedomain

const MaxStorageKeyLength int = 2048

type StoreKey string

func NewStoreKey(key string) (StoreKey, error) {
	value := StoreKey(key)

	if err := value.guard(); err != nil {
		return "", err
	}

	return value, nil
}

func (sk StoreKey) String() string {
	return string(sk)
}

func (sk StoreKey) guard() error {
	storageKeySize := len(sk)

	if storageKeySize > MaxStorageKeyLength {
		return NewStoreKeyMaxSizeExceededWithKeyAndSize(string(sk), int64(storageKeySize))
	}

	return nil
}

const maxSizeOnKeyExceededErrorMessage = "Store key size has been exceeded"

type StoreKeyMaxSizeExceeded struct {
	context *map[string]interface{}
}

func NewStoreKeyMaxSizeExceededWithKeyAndSize(key string, size int64) *StoreKeyMaxSizeExceeded {
	return &StoreKeyMaxSizeExceeded{
		context: &map[string]interface{}{
			"storage_key":      key,
			"storage_key_size": size,
		},
	}
}

func (sme *StoreKeyMaxSizeExceeded) Error() string {
	return maxSizeOnKeyExceededErrorMessage
}
