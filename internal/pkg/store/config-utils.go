package store

import (
	"github.com/hashicorp/raft"
)

func NewRaftConfigFromNode(n Node) (*raft.Config, raft.Configuration) {
	config, configuration := raft.DefaultConfig(), &raft.Configuration{}
	config.LocalID = raft.ServerID(n.NodeIdString())
	config.LogLevel = n.LogLevel

	if n.Replication.IsSingleNode() {
		configuration.Servers = []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       raft.ServerID(n.NodeIdString()),
				Address:  raft.ServerAddress(n.Replication.Address),
			},
		}
	}

	return config, *configuration
}
