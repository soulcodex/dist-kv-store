package config

import (
	"fmt"

	"codesignal/internal/pkg/store"
)

type PeerConfigPrimitives struct {
	HttpAddress string
	HttpPort    int64
}

func NewPeerConfig(httpAddress string, httpPort int64, node store.Node) *Config {
	configuration, err := LoadFromEnv()
	if err != nil {
		panic("failed to load env vars")
	}

	// Override server address with the peer port received on init
	configuration.NodeConfig = node
	configuration.Server.Port = httpPort
	configuration.Server.Address = fmt.Sprintf("%s:%d", httpAddress, httpPort)

	return configuration
}
