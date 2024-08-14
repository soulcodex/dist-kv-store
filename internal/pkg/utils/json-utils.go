package utils

import "encoding/json"

func MarshalFromMap(content map[string]interface{}) ([]byte, error) {
	response, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return response, nil
}
