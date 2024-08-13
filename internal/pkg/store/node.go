package store

import "fmt"

type Replication struct {
	Address     string
	JoinAddress *string
	Joined      bool
}

func (r *Replication) IsSingleNode() bool {
	return r.JoinAddress == nil || *r.JoinAddress == ""
}

func (r *Replication) MarkAsJoined() {
	r.Joined = true
}

type Node struct {
	Index       int64
	Name        string
	Replication *Replication
	Joiner      NodeJoiner
}

func NewNode(index int64, name string, replicaAddr string, joinAddress *string, joiner NodeJoiner) Node {
	var joinAddr *string = nil
	if joinAddress != nil && *(joinAddress) != "" {
		joinAddr = joinAddress
	}

	node := Node{
		Index:  index,
		Name:   name,
		Joiner: joiner,
		Replication: &Replication{
			Joined:      false,
			Address:     replicaAddr,
			JoinAddress: joinAddr,
		},
	}

	return node
}

func (n Node) NodeIdString() string {
	return fmt.Sprintf("%s-%d", n.Name, n.Index)
}

func (n Node) Join() error {
	if n.Replication.IsSingleNode() {
		return nil
	}

	if err := n.Joiner(n); err == nil {
		n.Replication.MarkAsJoined()
		return nil
	}

	panic(NewJoiningNodeErrorWithCtx(map[string]interface{}{
		"node_id":      n.NodeIdString(),
		"node_name":    n.Name,
		"address":      n.Replication.Address,
		"join_address": n.Replication.JoinAddress,
	}))
}
