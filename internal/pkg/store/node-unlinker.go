package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NodeUnlinker func(n Node) error

func NewRetryableHttpNodeUnlinker(client *http.Client) NodeUnlinker {
	return func(n Node) error {
		body, err := json.Marshal(map[string]interface{}{"id": n.NodeIdString()})
		if err != nil {
			return err
		}

		bodyReader := bytes.NewReader(body)
		resp, err := client.Post(fmt.Sprintf("http://%s/unlink", *n.Replication.JoinAddress), "application/json", bodyReader)
		if err != nil {
			return err
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		return nil
	}
}
