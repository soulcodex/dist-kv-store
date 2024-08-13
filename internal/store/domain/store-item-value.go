package storedomain

const MaxStorageContentLength int = 400

type StoreValue string

func NewStoreValue(value []byte) (StoreValue, error) {
	storeValue := StoreValue(value)

	if err := storeValue.guard(); err != nil {
		return "", err
	}

	return storeValue, nil
}

func (sv StoreValue) String() string {
	return string(sv)
}

func (sv StoreValue) guard() error {
	storageContentSize := len(sv)

	if storageContentSize > MaxStorageContentLength {
		return NewStoreContentMaxSizeExceededWithContentAndSize(int64(storageContentSize))
	}

	return nil
}

const maxSizeContentExceededErrorMessage = "Store content size has been exceeded"

type StoreContentMaxSizeExceeded struct {
	context *map[string]interface{}
}

func NewStoreContentMaxSizeExceededWithContentAndSize(size int64) *StoreContentMaxSizeExceeded {
	return &StoreContentMaxSizeExceeded{
		context: &map[string]interface{}{
			"storage_content_size": size,
		},
	}
}

func (sce *StoreContentMaxSizeExceeded) Error() string {
	return maxSizeContentExceededErrorMessage
}

const unexpectedStoreItemValueContentErrorMessage = "Error retrieving store value content"

type UnexpectedStoreItemValue struct {
	context *map[string]interface{}
}

func NewUnexpectedStoreItemValue() *UnexpectedStoreItemValue {
	return &UnexpectedStoreItemValue{
		context: &map[string]interface{}{},
	}
}

func (usc *UnexpectedStoreItemValue) Error() string {
	return unexpectedStoreItemValueContentErrorMessage
}
