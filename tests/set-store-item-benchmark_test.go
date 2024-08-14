package tests

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"testing"
)

func BenchmarkSetStoreItemOperation(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	b.Run("Set elements at store when is empty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_%d", i), gofakeit.BitcoinAddress())
		}
	})

	b.Run("Set elements at store when is not empty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_v%d", i), gofakeit.BitcoinAddress())
		}
	})
}
