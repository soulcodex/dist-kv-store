package storetest

import (
	"testing"
)

func BenchmarkGetStoreItemOperationWithEmptyStore(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = node.Services.KeyValueStore.Get("api_key_v1")
	}
}

func BenchmarkGetStoreItemOperationWithOneElementAtStoreWithSuccess(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	fillErr := fillStoreWithElements(node.Services.KeyValueStore, "api_key", 1)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = node.Services.KeyValueStore.Get("api_key-1")
	}
}

func BenchmarkGetStoreItemOperationWithFiftyElementsAtStoreWithSuccess(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	fillErr := fillStoreWithElements(node.Services.KeyValueStore, "api_key_v3", 50)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = node.Services.KeyValueStore.Get("api_key_v3-1")
	}
}

func BenchmarkGetStoreItemOperationWithFiftyElementsAtStoreWithoutSuccess(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	fillErr := fillStoreWithElements(node.Services.KeyValueStore, "api_key_v3", 50)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = node.Services.KeyValueStore.Get("api_key-0")
	}
}
