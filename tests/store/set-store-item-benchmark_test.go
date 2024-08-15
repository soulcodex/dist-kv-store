package storetest

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"testing"
)

func BenchmarkSetStoreItemOperationWithEmptyStore(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_v1_%d", i), gofakeit.BitcoinAddress())
	}
}

func BenchmarkSetStoreItemOperationWithNotEmptyStore(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	err = fillStoreWithElements(node.Services.KeyValueStore, "api_key_v", 50)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_v2_%d", i), gofakeit.BitcoinAddress())
	}
}
