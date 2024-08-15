package storetest

import (
	"testing"
)

func BenchmarkGetStoreItemOperation(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	fillErr := fillStoreWithElements(node.Services.KeyValueStore, "api_key", 1)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.Run("Get with one element at store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-1")
		}
	})

	b.Run("Get one element at store without success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-0")
		}
	})

	fillErr = fillStoreWithElements(node.Services.KeyValueStore, "api_key_v2", 9)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.Run("Get with ten elements at store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-1")
		}
	})

	b.Run("Get with ten elements at store without success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-0")
		}
	})

	fillErr = fillStoreWithElements(node.Services.KeyValueStore, "api_key_v3", 40)
	if fillErr != nil {
		b.Fatal(fillErr)
	}

	b.Run("Get with fifty elements at store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-1")
		}
	})

	b.Run("Get with fifty elements at store without success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = node.Services.KeyValueStore.Get("api_key-0")
		}
	})
}
