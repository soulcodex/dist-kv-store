package store

type KeyValueStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error

	Consensus() Consensus
}

type Consensus interface {
	Bootstrap(n Node) error
	Join(index, nodeAddress string) error
	Unlink(index string) error
	Stats() map[string]interface{}
}
