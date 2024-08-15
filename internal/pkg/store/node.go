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

func (r *Replication) MarkAsUnlinked() {
	r.Joined = false
}

type Node struct {
	Index       int64
	Name        string
	LogLevel    string
	Replication *Replication
	Joiner      NodeJoiner
	Unlinker    NodeUnlinker
}

func NewNode(
	index int64,
	name string,
	logLevel string,
	replicaAddr string,
	joinAddress *string,
	joiner NodeJoiner,
	unlinker NodeUnlinker,
) Node {
	var joinAddr *string = nil
	if joinAddress != nil && *(joinAddress) != "" {
		joinAddr = joinAddress
	}

	node := Node{
		Index:    index,
		Name:     name,
		Joiner:   joiner,
		LogLevel: logLevel,
		Unlinker: unlinker,
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

func (n Node) IsJoined() bool {
	return n.Replication.Joined
}

func (n Node) IsReplica() bool {
	return !n.Replication.IsSingleNode()
}

func (n Node) MarkAsJoined() {
	n.Replication.MarkAsJoined()
}

func (n Node) MarkAsUnlinked() {
	n.Replication.MarkAsJoined()
}

func (n Node) Unlink() error {
	if n.IsJoined() {
		if err := n.Unlinker(n); err != nil {
			return err
		}

		n.MarkAsUnlinked()
	}

	return nil
}

func (n Node) Join() error {
	if !n.IsReplica() {
		return nil
	}

	if err := n.Joiner(n); err == nil {
		n.MarkAsJoined()
		return nil
	}

	panic(NewJoiningNodeErrorWithCtx(map[string]interface{}{
		"node_id":      n.NodeIdString(),
		"node_name":    n.Name,
		"address":      n.Replication.Address,
		"join_address": n.Replication.JoinAddress,
	}))
}
