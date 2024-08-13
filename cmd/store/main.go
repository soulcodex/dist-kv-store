package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"codesignal/cmd/di"
	"codesignal/internal/pkg/config"
	"codesignal/internal/pkg/store"
)

const (
	defaultNodeName           = "node"
	defaultNodeId             = 1
	defaultAddress            = "0.0.0.0"
	defaultHttpPort           = 8085
	defaultReplicationAddress = "127.0.0.1"
	defaultReplicationPort    = 12000
)

var (
	nodeName           = flag.String("node-name", defaultNodeName, "The desired node name")
	nodeId             = flag.Int64("node-id", defaultNodeId, "The node identifier in the cluster")
	httpPort           = flag.Int64("http-port", defaultHttpPort, "The HTTP port to expose the server")
	replicationAddress = flag.String("replication-addr", fmt.Sprintf("%s:%d", defaultReplicationAddress, defaultReplicationPort), "The RPC server consensus address")
	joinAddr           = flag.String("join-addr", "", "Join to another node for replication quorum balancing on node bootstrap")
)

func main() {
	flag.Parse()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	nodeJoiner := store.NewRetryableHttpNodeJoiner(di.BuildNodeJoinRetryableHttpClient())
	node := store.NewNode(*nodeId, *nodeName, *replicationAddress, joinAddr, nodeJoiner)

	peerConfig := config.NewPeerConfig(defaultAddress, *httpPort, node)
	container := di.Init(peerConfig)
	context, cancel := di.Context()

	defer func() {
		cancel()
	}()

	if err := node.Join(); err != nil {
		container.Services.Log.Error().Err(err).Ctx(context).Msg("Error joining node")
		signals <- syscall.SIGTERM
	}

	if err := container.Services.KeyValueStore.Consensus().Bootstrap(container.Config.NodeConfig); err != nil {
		container.Services.Log.Error().Err(err).Ctx(context).Msg("Error bootstrapping key value store node")
		signals <- syscall.SIGTERM
	}

	container.Services.Log.Info().
		Ctx(context).
		Int64("node_id", peerConfig.NodeConfig.Index).
		Str("node_name", peerConfig.NodeConfig.Name).
		Str("replication_address", peerConfig.NodeConfig.Replication.Address).
		Bool("joined", peerConfig.NodeConfig.Replication.Joined).
		Msgf("Node <%s> starting", peerConfig.NodeConfig.NodeIdString())

	if err := container.Services.HttpServer.Run(signals); err != nil {
		container.Services.Log.Error().Err(err).Ctx(context).Msg("Error starting server")
	}
}
