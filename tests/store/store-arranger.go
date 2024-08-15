package storetest

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"testing"

	"codesignal/cmd/di"
	"codesignal/internal/pkg/config"
	"codesignal/internal/pkg/store"
	"codesignal/tests"
)

func storeModuleRouter() tests.HttpTestRouterFactory {
	return func(di *di.OktaDistributedKeyValueStorageContainer) *httprouter.Router {
		router := httprouter.New()
		di.StoreServices.RegisterHttpRoutes(router, di.Services)

		return router
	}
}

func setupStore(t *testing.T) (*di.OktaDistributedKeyValueStorageContainer, *httprouter.Router) {
	container, err := bootstrapSingleNodeCluster()
	if err != nil {
		t.Fatal(err)
	}

	return container, storeModuleRouter()(container)
}

func joinerAndUnlinker() (store.NodeJoiner, store.NodeUnlinker) {
	retryableClient := di.BuildNodeJoinRetryableHttpClient()
	joiner := store.NewRetryableHttpNodeJoiner(retryableClient.StandardClient())
	unlinker := store.NewRetryableHttpNodeUnlinker(retryableClient.StandardClient())
	return joiner, unlinker
}

func bootstrapClusterLeader(nodeContainer *di.OktaDistributedKeyValueStorageContainer) error {
	if err := nodeContainer.Services.KeyValueStore.Consensus().Bootstrap(nodeContainer.Config.NodeConfig); err != nil {
		return err
	}

	nodeContainer.Services.KeyValueStore.Consensus().WaitLeader()

	return nil
}

func fillStoreWithElements(store store.KeyValueStore, prefix string, times int) error {
	for i := 1; i <= times; i++ {
		key := fmt.Sprintf("%s-%s", prefix, strconv.Itoa(i))
		if err := store.Set(key, gofakeit.BitcoinAddress()); err != nil {
			return err
		}
	}

	return nil
}

func bootstrapSingleNodeCluster() (*di.OktaDistributedKeyValueStorageContainer, error) {
	serverPort, _ := tests.GetFreeTCPPort()
	replicationPort, _ := tests.GetFreeTCPPort()

	joiner, unlinker := joinerAndUnlinker()
	node := store.NewNode(1, "node", hclog.Off.String(), fmt.Sprintf("localhost:%d", replicationPort), nil, joiner, unlinker)
	peerConfig := config.NewPeerConfig("0.0.0.0", int64(serverPort), node)

	container := di.Init(peerConfig)

	if err := bootstrapClusterLeader(container); err != nil {
		return nil, err
	}

	return container, nil
}
