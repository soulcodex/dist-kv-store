package storetest

import (
	"fmt"
	"testing"
)

func BenchmarkSetStoreItemOperationWithEmptyStore(b *testing.B) {
	node, err := bootstrapSingleNodeCluster()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.ResetTimer()
		_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_v1_%d", i), fmt.Sprintf("fake_value_%d", i))
	}

	fmt.Printf("%s took %s with b.N = %d\n", "BenchmarkSetStoreItemOperationWithEmptyStore", b.Elapsed(), b.N)
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
		b.ResetTimer()
		_ = node.Services.KeyValueStore.Set(fmt.Sprintf("api_key_v2_%d", i), fmt.Sprintf("fake_value_%d", i))
	}
}
