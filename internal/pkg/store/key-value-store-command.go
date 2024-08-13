package store

type Cmd string

const (
	Set    Cmd = "SET"
	Delete Cmd = "DELETE"
	Join   Cmd = "JOIN"
)

func (c Cmd) String() string {
	return string(c)
}

type SetCmd struct {
	Cmd
	Key   string
	Value string
}

func NewSetCmd(key string, value string) SetCmd {
	return SetCmd{
		Cmd:   Set,
		Key:   key,
		Value: value,
	}
}

func NewSetCmdFromMap(cmd map[string]interface{}) SetCmd {
	return SetCmd{
		Cmd:   Set,
		Key:   cmd["key"].(string),
		Value: cmd["value"].(string),
	}
}

func (sc SetCmd) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cmd":   sc.Cmd,
		"key":   sc.Key,
		"value": sc.Value,
	}
}

type DeleteCmd struct {
	Cmd
	Key string
}

func NewDeleteCmd(key string) DeleteCmd {
	return DeleteCmd{
		Cmd: Delete,
		Key: key,
	}
}

func NewDeleteCmdFromMap(cmd map[string]interface{}) DeleteCmd {
	return DeleteCmd{
		Cmd: Delete,
		Key: cmd["key"].(string),
	}
}

func (dc DeleteCmd) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cmd": dc.Cmd,
		"key": dc.Key,
	}
}

type JoinCmd struct {
	Cmd
	NodeId  string
	Address string
}

func NewJoinCmd(id, address string) JoinCmd {
	return JoinCmd{
		Cmd:     Join,
		NodeId:  id,
		Address: address,
	}
}

func NewJoinCmdFromMap(cmd map[string]interface{}) JoinCmd {
	return JoinCmd{
		Cmd:     Delete,
		NodeId:  cmd["id"].(string),
		Address: cmd["address"].(string),
	}
}

func (dc JoinCmd) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cmd":     dc.Cmd,
		"id":      dc.NodeId,
		"address": dc.Address,
	}
}
