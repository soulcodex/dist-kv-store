package tests

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"net"
	"strconv"

	"codesignal/cmd/di"
	"codesignal/internal/pkg/config"
	"codesignal/internal/pkg/store"
)

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = l.Close()
	}()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func joinerAndUnlinker() (store.NodeJoiner, store.NodeUnlinker) {
	retryableClient := di.BuildNodeJoinRetryableHttpClient()
	joiner := store.NewRetryableHttpNodeJoiner(retryableClient.StandardClient())
	unlinker := store.NewRetryableHttpNodeUnlinker(retryableClient.StandardClient())
	return joiner, unlinker
}

func bootstrapCluster(nodeContainer *di.OktaDistributedKeyValueStorageContainer) error {
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
	serverPort, _ := getFreePort()
	replicationPort, _ := getFreePort()

	joiner, unlinker := joinerAndUnlinker()
	node := store.NewNode(1, "node", fmt.Sprintf("localhost:%d", replicationPort), nil, joiner, unlinker)
	peerConfig := config.NewPeerConfig("0.0.0.0", int64(serverPort), node)

	container := di.Init(peerConfig)

	if err := bootstrapCluster(container); err != nil {
		return nil, err
	}

	return container, nil
}
