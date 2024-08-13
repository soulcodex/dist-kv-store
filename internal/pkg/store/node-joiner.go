package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
)

type NodeJoiner func(n Node) error

func NewRetryableHttpNodeJoiner(client *retryablehttp.Client) NodeJoiner {
	return func(n Node) error {
		body, err := json.Marshal(map[string]interface{}{"address": n.Replication.Address, "id": n.NodeIdString()})
		if err != nil {
			return err
		}

		bodyReader := bytes.NewReader(body)
		resp, err := client.Post(fmt.Sprintf("http://%s/join", *n.Replication.JoinAddress), "application/json", bodyReader)
		if err != nil {
			return err
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		return nil
	}
}
