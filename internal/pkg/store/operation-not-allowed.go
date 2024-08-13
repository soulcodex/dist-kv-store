package store

const storeOperationNotAllowedMessage = "Operation not allowed"

type OperationNotAllowed struct {
	context map[string]interface{}
}

func NewOperationNotAllowedWhenNotLeaderWithKey(n Node, operation Cmd, key string) *OperationNotAllowed {
	return &OperationNotAllowed{
		context: map[string]interface{}{
			"cmd":       operation,
			"key":       key,
			"node_id":   n.NodeIdString(),
			"node_name": n.Name,
		},
	}
}

func (e *OperationNotAllowed) Error() string {
	return storeOperationNotAllowedMessage
}
