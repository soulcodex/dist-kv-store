package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonBodyAsMap(req *http.Request) (map[string]interface{}, error) {
	content := map[string]interface{}{}

	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buffer.Bytes(), &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func JsonBodyToStruct[T any](req *http.Request, target *T) error {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buffer.Bytes(), &target)
	if err != nil {
		return err
	}

	return nil
}
