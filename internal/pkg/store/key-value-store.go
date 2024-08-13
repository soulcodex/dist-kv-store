package store

import (
	"encoding/json"
	"errors"
	"net"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	"github.com/rs/zerolog"
)

type InMemoryKeyValueStore struct {
	mutex  sync.Mutex
	items  map[string]string
	logger zerolog.Logger

	node           Node
	raft           *raft.Raft
	raftTimeout    time.Duration
	raftTCPTimeout time.Duration
}

func NewInMemoryKeyValueStore(logger zerolog.Logger, n Node) *InMemoryKeyValueStore {
	return &InMemoryKeyValueStore{
		mutex:          sync.Mutex{},
		items:          make(map[string]string),
		logger:         logger,
		node:           n,
		raft:           nil,
		raftTimeout:    15 * time.Second,
		raftTCPTimeout: 15 * time.Second,
	}
}

func (iks *InMemoryKeyValueStore) Consensus() Consensus {
	return iks
}

func (iks *InMemoryKeyValueStore) Bootstrap(n Node) error {
	config, configuration := NewRaftConfigFromNode(n)

	addr, err := net.ResolveTCPAddr("tcp", n.Replication.Address)
	if err != nil {
		return err
	}

	transport, err := raft.NewTCPTransport(n.Replication.Address, addr, 3, iks.raftTCPTimeout, os.Stderr)
	if err != nil {
		return err
	}

	logStore, stableStore, snapshotStore := raft.NewInmemStore(), raft.NewInmemStore(), raft.NewInmemSnapshotStore()

	server, err := raft.NewRaft(config, (*keyValueStoreFSM)(iks), logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return err
	}
	iks.raft = server

	server.BootstrapCluster(configuration)

	return nil
}

func (iks *InMemoryKeyValueStore) Join(nodeId, nodeAddress string) error {
	cmd := NewJoinCmd(nodeId, nodeAddress)

	config := iks.raft.GetConfiguration()
	if err := config.Error(); err != nil {
		return err
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeId) || server.Address == raft.ServerAddress(nodeAddress) {
			if server.Address == raft.ServerAddress(nodeAddress) && server.ID == raft.ServerID(nodeId) {
				iks.logger.Warn().
					Int64("node_id", iks.node.Index).
					Str("node_name", iks.node.Name).
					Str("ignored_node_id", cmd.NodeId).
					Str("ignored_address", cmd.Address).
					Msg("node is already cluster member")
				return nil
			}

			future := iks.raft.RemoveServer(server.ID, 0, 0)
			if err := future.Error(); err != nil {
				return err
			}
		}
	}

	addOp := iks.raft.AddVoter(raft.ServerID(cmd.NodeId), raft.ServerAddress(cmd.Address), 0, 0)
	if err := addOp.Error(); err != nil {
		return err
	}

	iks.logger.Warn().
		Int64("node_id", iks.node.Index).
		Str("node_name", iks.node.Name).
		Str("joined_id", cmd.NodeId).
		Str("joined_address", cmd.Address).
		Msg("node joined successfully")

	return nil
}

func (iks *InMemoryKeyValueStore) Stats() map[string]interface{} {
	raw, stats := iks.raft.Stats(), make(map[string]interface{})

	for k, v := range raw {
		stats[k] = v
	}

	return stats
}

func (iks *InMemoryKeyValueStore) Get(key string) (string, error) {
	iks.mutex.Lock()
	defer iks.mutex.Unlock()

	if value, ok := iks.items[key]; ok {
		return value, nil
	}

	return "", errors.New("key not found")
}

func (iks *InMemoryKeyValueStore) Set(key string, value string) error {
	if iks.raft.State() != raft.Leader {
		return NewOperationNotAllowedWhenNotLeaderWithKey(iks.node, Set, key)
	}

	cmd, err := json.Marshal(NewSetCmd(key, value).ToMap())
	if err != nil {
		return err
	}

	f := iks.raft.Apply(cmd, iks.raftTimeout)
	return f.Error()
}

func (iks *InMemoryKeyValueStore) Delete(key string) error {
	if iks.raft.State() != raft.Leader {
		return NewOperationNotAllowedWhenNotLeaderWithKey(iks.node, Delete, key)
	}

	cmd, err := json.Marshal(NewDeleteCmd(key).ToMap())
	if err != nil {
		return err
	}

	f := iks.raft.Apply(cmd, iks.raftTimeout)
	return f.Error()
}
