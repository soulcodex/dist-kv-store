package store

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/raft"
	"io"
	"maps"
)

type keyValueStoreFSM InMemoryKeyValueStore

func (fsm *keyValueStoreFSM) Apply(log *raft.Log) interface{} {
	var raw map[string]interface{}
	if err := json.Unmarshal(log.Data, &raw); err != nil {
		// We should panic given that we are in a critical path, and we should not
		// continue if we can't unmarshal the command and apply the changes.
		panic(fmt.Sprintf("command unmarshal error: %s", err.Error()))
	}

	switch raw["cmd"].(string) {
	case Set.String():
		cmd := NewSetCmdFromMap(raw)
		return fsm.applySet(cmd.Key, cmd.Value)
	case Delete.String():
		cmd := NewDeleteCmdFromMap(raw)
		return fsm.applyDelete(cmd.Key)
	default:
		panic(fmt.Sprintf("unknown command received: %s", raw["cmd"].(string)))
	}
}

func (fsm *keyValueStoreFSM) Snapshot() (raft.FSMSnapshot, error) {
	fsm.mutex.Lock()
	defer fsm.mutex.Unlock()

	backup := maps.Clone(fsm.items)
	return newKeyValueStoreFSMSnapshot(backup), nil
}

func (fsm *keyValueStoreFSM) Restore(rc io.ReadCloser) error {
	backup := make(map[string]string)
	if err := json.NewDecoder(rc).Decode(&backup); err != nil {
		return err
	}

	// Following the raft algorithm, we should not apply locking
	// mechanisms here, as the Raft library will handle the serialization
	fsm.items = backup
	return nil
}

func (fsm *keyValueStoreFSM) applySet(key string, value string) interface{} {
	fsm.mutex.Lock()
	defer fsm.mutex.Unlock()

	fsm.items[key] = value

	return nil
}

func (fsm *keyValueStoreFSM) applyDelete(key string) interface{} {
	fsm.mutex.Lock()
	defer fsm.mutex.Unlock()

	delete(fsm.items, key)

	return nil
}
