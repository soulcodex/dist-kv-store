package store

import (
	"encoding/json"
	"github.com/hashicorp/raft"
)

type keyValueStoreFSMSnapshot struct {
	store map[string]string
}

func newKeyValueStoreFSMSnapshot(store map[string]string) *keyValueStoreFSMSnapshot {
	return &keyValueStoreFSMSnapshot{store: store}
}

func (kvs *keyValueStoreFSMSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		backup, err := json.Marshal(kvs.store)
		if err != nil {
			return err
		}

		if _, err := sink.Write(backup); err != nil {
			return err
		}

		return sink.Close()
	}()

	if err != nil {
		_ = sink.Cancel()
	}

	return err
}

func (kvs *keyValueStoreFSMSnapshot) Release() {}
